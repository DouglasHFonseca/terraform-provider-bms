# Terraform Provider BMS (Terraform Plugin Framework)

This repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework).

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Prepare Terraform for local provider install

- First, find the GOBIN path where Go installs its binaries. This path may vary depending on how your Go environment variables are set up.
  - Run `go env GOBIN` to find out. If it's not set, you will need to define it by running the command `export GOBIN=/usr/local/go/bin.`

  
  - With the appropriate settings made as mentioned above, proceed with the override. Create a file called .terraformrc in your root directory (~) and add the dev_overrides block below. Change the <PATH> to the value returned from the go env GOBIN command mentioned earlier, keeping in mind that if it's not set, the path will be what was defined in the export above.

```
provider_installation {

  dev_overrides {
      "hashicorp.com/edu/bms" = "/usr/local/go/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install .
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Fill this in for each provider
- After that, run terraform apply on the main.tf file located in the examples folder.