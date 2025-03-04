// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
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

func TestProviderUpgrade(t *testing.T) {
	testProviderUpgrade(t, "provider-update/ts")
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
