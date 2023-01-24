// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build dotnet || all
// +build dotnet all

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccPrivateKeyCSharp(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			DebugLogLevel: 9,
			Verbose:       true,
			Stdout:        os.Stdout,
			Stderr:        os.Stderr,
			Dir:           path.Join(getCwd(t), "private-key", "dotnet"),
		})

	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.Tls",
		},
	})

	return baseCsharp
}
