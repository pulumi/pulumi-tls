// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest"
	"github.com/pulumi/providertest/optproviderupgrade"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/require"
)

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
		SkipRefresh:          true,
		Quick:                true,
		RunUpdateTest:        false,
	}
}

const (
	providerName    = "tls"
	baselineVersion = "5.0.0"
)

func testProviderUpgrade(t *testing.T, dir string, opts ...opttest.Option) {
	cwd, err := os.Getwd()
	require.NoError(t, err)
	options := []opttest.Option{
		opttest.DownloadProviderVersion(providerName, baselineVersion),
		opttest.LocalProviderPath(providerName, filepath.Join(cwd, "..", "bin")),
	}
	options = append(options, opts...)
	pt := pulumitest.NewPulumiTest(t, dir, options...)
	result := providertest.PreviewProviderUpgrade(
		t,
		pt,
		providerName,
		baselineVersion,
		optproviderupgrade.DisableAttach())
	assertpreview.HasNoReplacements(t, result)
}