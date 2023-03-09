// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccPrivateKeyGo(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "private-key", "go"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccSelfSignedCertGoForcesNewResource(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:                  path.Join(getCwd(t), "self-signed-cert", "go"),
			ExpectRefreshChanges: true,
		})

	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-tls/sdk/v5",
		},
	})

	return baseGo
}
