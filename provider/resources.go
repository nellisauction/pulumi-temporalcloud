package temporalcloud

import (
	"fmt"
	"path/filepath"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	temporalshim "github.com/temporalio/terraform-provider-temporalcloud/shim"

	"github.com/nellisauction/pulumi-temporalcloud/provider/pkg/version"
)

const (
	mainPkg = "temporalcloud"
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-temporalcloud/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 pf.ShimProvider(temporalshim.NewProvider()),
		Name:              "temporalcloud",
		Version:           version.Version,
		DisplayName:       "Temporal Cloud",
		Publisher:         "NellisAuction",
		PluginDownloadURL: "github://api.github.com/nellisauction/pulumi-temporalcloud",
		Description:       "A Pulumi package for managing Temporal Cloud resources.",
		Keywords:          []string{"pulumi", "temporalcloud", "category/infrastructure"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/nellisauction/pulumi-temporalcloud",
		Repository:        "https://github.com/nellisauction/pulumi-temporalcloud",
		GitHubOrg:         "temporalio",
		Config: map[string]*tfbridge.SchemaInfo{
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TEMPORAL_CLOUD_API_KEY"},
				},
			},
			"endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TEMPORAL_CLOUD_ENDPOINT"},
				},
			},
			"allow_insecure": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TEMPORAL_CLOUD_ALLOW_INSECURE"},
				},
			},
			"allowed_account_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TEMPORAL_CLOUD_ALLOWED_ACCOUNT_ID"},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "@nellisauction/pulumi-temporalcloud",
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{RespectSchemaVersion: true}
			i.PyProject.Enabled = true
			return i
		})(),
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/nellisauction/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences:    map[string]string{"Pulumi": "3.*"},
			Namespaces:           map[string]string{mainPkg: "TemporalCloud"},
		},
		MetadataInfo:                   tfbridge.NewProviderMetadata(metadata),
		EnableZeroDefaultSchemaVersion: true,
		EnableAccurateBridgePreview:    true,
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("temporalcloud_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
