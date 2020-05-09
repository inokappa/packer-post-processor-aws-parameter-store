package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2/hcldec"
	awscommon "github.com/hashicorp/packer/builder/amazon/common"
	_ "github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	_ "github.com/hashicorp/packer/template/interpolate"
	_ "os"
	_ "regexp"
	"strings"
)

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	p.config.ctx.Funcs = awscommon.TemplateFuncs
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
	}, raws...)

	if err != nil {
		return err
	}

	for _, tpl := range p.config.Parameters {
		if tpl.Name == "" {
			return errors.New("empty `name` is not allowed. Please make sure that it is set correctly")
		}
	}

	return nil
}

func (p *PostProcessor) PostProcess(ctx context.Context, ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, bool, error) {
	session, err := p.config.Session()
	if err != nil {
		return nil, false, false, err
	}
	config := session.Config

	amiId := p.GetImageId(artifact)
	for _, tpl := range p.config.Parameters {
		// if tpl.SecureString == nil {
		// 	tpl.SecureString = false
		// }

		// if tpl.AmiDataType == nil {
		// 	tpl.AmiDataType = false
		// }

		message := fmt.Sprintf("Register the AMI ID to Parameter Store.(Parameter Name: %s, AMI ID: %s)", tpl.Name, amiId)
		ui.Say(message)
		_, err = RegisterToParameterStore(amiId, tpl.Name, tpl.SecureString, tpl.AmiDataType)
		if err != nil {
			return nil, true, false, err
		}
	}

	artifact = &awscommon.Artifact{
		Amis: map[string]string{
			*config.Region: amiId,
		},
	}

	return artifact, true, false, nil
	// return artifact, true, nil
}

func (p *PostProcessor) GetImageId(artifact packer.Artifact) string {
	// example: ap-northeast-1:ami-xxxxxxxxxxxxxxxx
	splitedString := strings.Split(artifact.Id(), ":")
	amiId := splitedString[len(splitedString)-1]

	return amiId
}
