// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build dotnet || all
// +build dotnet all

package examples

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccPrivateKeyCSharp(t *testing.T) {

	d := path.Join(getCwd(t), "private-key", "dotnet")

	debugLog := fmt.Sprintf("%s/debug.json", d)
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Env: []string{"PULUMI_DEBUG_GRPC=" + debugLog},
			Dir: d,
		})

	integration.ProgramTest(t, &test)

	jq, err := exec.LookPath("jq")
	if err != nil {
		t.Fatal(err)
	}

	opts := &integration.ProgramTestOptions{
		Verbose: true,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
	}
	if err := integration.RunCommand(t, "jq", []string{jq, ".", debugLog}, d, opts); err != nil {
		t.Fatal(err)
	}

	t.Errorf("FAIL to show JQ debug logs")
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
