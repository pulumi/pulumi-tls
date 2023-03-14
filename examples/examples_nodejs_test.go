// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"

	"github.com/pulumi/pulumi-tls/examples/v5/internal/testutil"
)

func TestAccPrivateKeyTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "private-key", "ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccSelfSignedCert(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "self-signed-cert", "ts"),
		})

	integration.ProgramTest(t, &test)
}

// This test is a bit special as it is not an example, but a way to detect changes in the current version of provider or
// provider SDK that would generate unexpected plans for users that have stacks provisioned on the baseline version
// which is typically the latest released version.
//
// Currently the baseline version is specified by editing ./provider-update/ts/package.json
// .dependencies["@pulumi/tls"].
//
// Note that this is currently pointing at v5.0.0, as pointing it at "^4.0.0" breaks the test. We do not expect the
// update test to pass across major versions in general, although in the specific case of this provider there is some
// compensating PreStateUpgradeHook code that makes the plans more tolerable (Update instead of Replace).
func TestProviderUpdate(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "provider-update", "ts"),
		})

	testutil.ProviderUpdateTest(t, "pulumi-resource-tls", test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/tls",
		},
	})

	return baseJS
}
