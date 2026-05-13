package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/temporalio/terraform-provider-temporalcloud/internal/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New("dev")()
}
