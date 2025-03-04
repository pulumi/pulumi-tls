// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
//go:build yaml || all
// +build yaml all

package examples

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccGetCert(t *testing.T) {
	test := getBaseOptions().With(integration.ProgramTestOptions{
		Dir:                    filepath.Join(getCwd(t), "get-cert", "yaml"),
		ExtraRuntimeValidation: validateExpectedVsActual,
	})
	integration.ProgramTest(t, &test)
}

// Stacks may define tests inline by a simple convention of providing
// ${test}__expect and ${test}__actual pairs. For example:
//
//	outputs:
//	  test1__expect: 1
//	  test1__actual: ${res1.out}
//
// This function interpretes these outputs to actual tests.
func validateExpectedVsActual(t *testing.T, stack integration.RuntimeValidationStackInfo) {
	expects := map[string]interface{}{}
	actuals := map[string]interface{}{}
	for n, output := range stack.Outputs {
		if strings.HasSuffix(n, "__actual") {
			actuals[strings.TrimSuffix(n, "__actual")] = output
		}
		if strings.HasSuffix(n, "__expect") {
			expects[strings.TrimSuffix(n, "__expect")] = output
		}
	}
	for k := range expects {
		k := k
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, expects[k], actuals[k])
		})
	}
}


func TestSelfSignedKeyUpgrade(t *testing.T) {
	testProviderUpgrade(t, "self-signed-cert/yaml")
}