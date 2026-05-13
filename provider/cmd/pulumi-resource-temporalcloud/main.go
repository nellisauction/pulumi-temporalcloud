package main

import (
	"context"
	_ "embed"

	pftfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	provider "github.com/nellisauction/pulumi-temporalcloud/provider"
)

//go:embed schema.json
var pulumiSchema []byte

//go:embed bridge-metadata.json
var bridgeMetadata []byte

func main() {
	meta := pftfbridge.ProviderMetadata{
		PackageSchema:  pulumiSchema,
		BridgeMetadata: bridgeMetadata,
	}
	pftfbridge.Main(context.Background(), "temporalcloud", provider.Provider(), meta)
}
