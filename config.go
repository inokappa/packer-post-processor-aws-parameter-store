//go:generate mapstructure-to-hcl2 -type Config
package main

import (
	awscommon "github.com/hashicorp/packer/builder/amazon/common"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig    `mapstructure:",squash"`
	awscommon.AccessConfig `mapstructure:",squash"`

	Parameters []struct {
		Name         string `mapstructure:"name"`
		SecureString bool   `mapstructure:"secure_string"`
		AmiDataType  bool   `mapstructure:"ami_data_type"`
	} `mapstructure:"parameters"`

	ctx interpolate.Context
}
