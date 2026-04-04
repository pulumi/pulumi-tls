package tls

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"

	"github.com/pulumi/pulumi-tls/provider/v5/pkg/version"
)

var setTestVersion sync.Once

func providerInfoForTest() tfbridge.ProviderInfo {
	setTestVersion.Do(func() {
		version.Version = "5.0.0"
	})
	return Provider()
}

func TestProviderPreservesExistingTokens(t *testing.T) {
	t.Parallel()

	info := providerInfoForTest()

	expectedResources := map[string]string{
		"tls_cert_request":        "tls:index/certRequest:CertRequest",
		"tls_locally_signed_cert": "tls:index/locallySignedCert:LocallySignedCert",
		"tls_private_key":         "tls:index/privateKey:PrivateKey",
		"tls_self_signed_cert":    "tls:index/selfSignedCert:SelfSignedCert",
	}
	for tfToken, want := range expectedResources {
		resource, ok := info.Resources[tfToken]
		require.Truef(t, ok, "missing resource mapping for %q", tfToken)
		assert.Equal(t, want, string(resource.Tok))
	}

	expectedDataSources := map[string]string{
		"tls_certificate": "tls:index/getCertificate:getCertificate",
		"tls_public_key":  "tls:index/getPublicKey:getPublicKey",
	}
	for tfToken, want := range expectedDataSources {
		dataSource, ok := info.DataSources[tfToken]
		require.Truef(t, ok, "missing data source mapping for %q", tfToken)
		assert.Equal(t, want, string(dataSource.Tok))
	}

	require.NotNil(t, info.Resources["tls_self_signed_cert"].PreStateUpgradeHook)
}

func TestProviderComputesTokensForAllUpstreamEntities(t *testing.T) {
	t.Parallel()

	info := providerInfoForTest()

	resourceCount := 0
	info.P.ResourcesMap().Range(func(key string, _ shim.Resource) bool {
		resourceCount++

		resource, ok := info.Resources[key]
		require.Truef(t, ok, "missing resource mapping for upstream resource %q", key)
		require.NotEmptyf(t, resource.Tok, "missing resource token for upstream resource %q", key)
		return true
	})
	assert.Len(t, info.Resources, resourceCount)

	dataSourceCount := 0
	info.P.DataSourcesMap().Range(func(key string, _ shim.Resource) bool {
		dataSourceCount++

		dataSource, ok := info.DataSources[key]
		require.Truef(t, ok, "missing data source mapping for upstream data source %q", key)
		require.NotEmptyf(t, dataSource.Tok, "missing data source token for upstream data source %q", key)
		return true
	})
	assert.Len(t, info.DataSources, dataSourceCount)
}
