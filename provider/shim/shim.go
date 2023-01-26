package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	tlsProvider "github.com/hashicorp/terraform-provider-tls/internal/provider"
)

func NewProvider() provider.Provider {
	p := tlsProvider.New()
	return p
}
