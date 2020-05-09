# packer-post-processor-aws-parameter-store

## Description

* Packer post-processor plugin for AWS Systems Manager Parameter Store.

## Installation

```sh
$ mkdir -p ~/.packer.d/plugins
$ wget https://github.com/inokappa/packer-post-processor-aws-parameter-store/releases/download/v0.0.2/packer-post-processor-aws-parameter-store_darwin_amd64 -O ~/.packer.d/plugins/packer-post-processor-aws-parameter-store
```

## Usage

The following example is a template for registering an AMI ID with a given Parameter Store.

```json
{
  "variables": {
    "aws_access_key": "{{env `AWS_ACCESS_KEY_ID`}}",
    "aws_secret_key": "{{env `AWS_SECRET_ACCESS_KEY`}}",
    "ssh_keypair_name": "{{env `KEYPAIR_NAME`}}",
    "ssh_private_key_file": "{{env `PRIVATE_KEY_PATH`}}",
    "ami_name": "packer-sample-{{timestamp}}"
  },
  "builders": [
    {
      "type": "amazon-ebs",
... snip ....
    }
  ],
  "provisioners": [
    {
      "type": "shell",
      "inline": [
        "echo 'shell'"
      ]
    }
  ],
  "post-processors": [
    [
      {
        "type": "aws-parameter-store",
        "parameters": [
          {
            "name": "parameter-ami1",
            "secure_string": true
          },
          {
            "name": "parameter-ami2",
            "data_type": "ami"
          }
        ]
      }
    ]
  ]
}
```
