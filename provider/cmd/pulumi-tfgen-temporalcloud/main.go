package main

import (
	pftfgen "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"
	provider "github.com/nellisauction/pulumi-temporalcloud/provider"
)

func main() {
	pftfgen.Main("temporalcloud", provider.Provider())
}
