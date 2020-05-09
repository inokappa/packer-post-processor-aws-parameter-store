package main

import (
	_ "fmt"
	_ "os"

	"github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/awserr"
	_ "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"strconv"
)

func RegisterToParameterStore(amiId string, parameterName string, secureString bool, amiDataType bool) ([]string, error) {
	svc := ssm.New(session.New())
	input := &ssm.PutParameterInput{
		Name:      aws.String(parameterName),
		Value:     aws.String(amiId),
		Overwrite: aws.Bool(true),
		Type:      aws.String("String"),
	}

	if secureString && !amiDataType {
		input.SetType("SecureString")
	}

	if amiDataType {
		input.SetDataType("aws:ec2:image")
	}

	result, err := svc.PutParameter(input)
	if err != nil {
		// fmt.Println(err.Error())
		return nil, err
	}

	res := []string{
		*result.Tier,
		strconv.FormatInt(*result.Version, 10),
	}

	return res, nil
}
