// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	testutils "github.com/pulumi/providertest/replay"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	provider "github.com/pulumi/pulumi-tls/provider/v5"
	"github.com/pulumi/pulumi-tls/provider/v5/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
)

// Tests that upgrading from to Plugin Framework is NOT a replacement plan for SelfSignedCert, when running Diff against
// a state populated by an older version of the provider.
//
// See https://github.com/pulumi/pulumi-tls/issues/173
func TestRegress173(t *testing.T) {
	ctx := context.Background()
	testCase := `
{
  "method": "/pulumirpc.ResourceProvider/Diff",
  "request": {
    "id": "253125812167639846332020603447562130563",
    "urn": "urn:pulumi:dev::pulumi-tls-173::tls:index/selfSignedCert:SelfSignedCert::test",
    "olds": {
      "certPem": "-----BEGIN CERTIFICATE-----\nMIICmzCCAYOgAwIBAgIRAL5uQhUkUvQ7zW/eMfaG1IMwDQYJKoZIhvcNAQELBQAw\nADAeFw0yMzAzMDMyMDEwNDhaFw0yNTAzMDMwODEwNDhaMAAwggEiMA0GCSqGSIb3\nDQEBAQUAA4IBDwAwggEKAoIBAQC9nMBYMHVxYVXxNdrB451HqjtdqlhA8hzPj+KO\n6pcU5tNDMn7LspiWpEJYWxJNCu3E5+Z08D4NuDxAyLE6SISMh5+ncJN6lYsLVBUg\nmipu0ov+/kbaGduSTYk47Ul3FMSgBqeKjopnZ58eWaDe0uX0ZJTvPm3Q61VoL0HC\nN5pf0n3844LehNYAwYdMat2tLb9HlZEfZ07qZv3R5s3F32ebRB8ngVP4VGSPAy7V\nqslv+ZPTZLh48dgh7JA6DgSVacOdLQlMUj11SBmS+THX7T+jXor8qhpCN0AVwxGu\nKQeg0PCwEPTfBiw+nzPXWrUh7DsVa7CjJzKQQ22M3hyaW1cjAgMBAAGjEDAOMAwG\nA1UdEwEB/wQCMAAwDQYJKoZIhvcNAQELBQADggEBAICmUHA8AKR+Y/wqhgnYEuc6\nuiweZRgrB0h/GIqNisrcWZfi+Tu28cDtAQLC2GuMluZZVocGBdBdNSc6VrXZN/gq\n1w7VgM5oNqHLoNmw1R6AetnUDBtokTKuCbYY8kwBtbzYmvJoERCOiWmBXHLb/BQm\n0o8C7i3A2hlOKI/ZohjVNrZro48nw/oyYyvQnc3ky98qn7MNhx5zx5is/UIxxwST\n2BqQhaXFggtX6dLTmp6aSWkg5qCoxCfKnq7LXQarUOT1JrH9PFDWlHlWL2w2FxAz\nG14eWrTMO0VtEqXbtEtVzO6W1T7rms+u5RmEB4g7SbxQuSnmarLUH05802N0KCI=\n-----END CERTIFICATE-----\n",
      "earlyRenewalHours": 0,
      "id": "253125812167639846332020603447562130563",
      "keyAlgorithm": "RSA",
      "privateKeyPem": {
        "4dabf18193072939515e22adb298388d": "1b47061264138c4ac30d75fd1eb44270",
        "value": "097d55ae27a2e81f19b7dc586014b394e10cdc12"
      },
      "readyForRenewal": false,
      "validityEndTime": "2025-03-03T03:10:48.531545-05:00",
      "validityPeriodHours": 17532,
      "validityStartTime": "2023-03-03T15:10:48.531545-05:00"
    },
    "news": {
      "allowedUses": [],
      "privateKeyPem": {
        "4dabf18193072939515e22adb298388d": "1b47061264138c4ac30d75fd1eb44270",
        "value": "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAvZzAWDB1cWFV8TXaweOdR6o7XapYQPIcz4/ijuqXFObTQzJ+\ny7KYlqRCWFsSTQrtxOfmdPA+Dbg8QMixOkiEjIefp3CTepWLC1QVIJoqbtKL/v5G\n2hnbkk2JOO1JdxTEoAanio6KZ2efHlmg3tLl9GSU7z5t0OtVaC9BwjeaX9J9/OOC\n3oTWAMGHTGrdrS2/R5WRH2dO6mb90ebNxd9nm0QfJ4FT+FRkjwMu1arJb/mT02S4\nePHYIeyQOg4ElWnDnS0JTFI9dUgZkvkx1+0/o16K/KoaQjdAFcMRrikHoNDwsBD0\n3wYsPp8z11q1Iew7FWuwoycykENtjN4cmltXIwIDAQABAoIBAB/0nAIA4LokAlzt\ni5mjp60dRnYJsGf5pdthT3hwltfB8xbfGrlvtwGkWz4S0ynCzsGhp5hLqNmdFCKC\n8EN3V0drz/9jHKfLLizRleHxuXcF6uwlpIE7XGLyyE6cxwXXrS+fD3ttfGvi7dEq\nn88N0g13KhaL+ev8zV8Kn9WAk3bUc2npfUhx5WU+svC7u+XH2THBS/yl40vS+vXM\n2bP2iKMn2QbL9LlyplSxvXoWg1inQiTPm4KA+3EAQwsxUhaPuLV7GhlSCfWRXEOn\nDPMWyMYbI7YIN3bCiRW/sd115z4oopGB2+L0//mjj/I+LLq/6Ad2hteGBFmDdJIZ\nXRZHd1kCgYEA0mhK1DKIeenKOEaNBxeHgGZ1hxaRAZCmHYhB2pHvcMuA324jxNXX\nnLMvJHLhfiIKJVQchiC9jpHOkdJnjskvYM1u4tuMWjt9SfOTrZmH1FX+CtKbMB05\n64iTXynZYRt2cD2hwzP+F4A9Cltlf9TP+x6qdsDz8SIiFxIWZwh853UCgYEA5rLq\nVWRcpPsgj6IGD3lyoQYbD9p79LmWcT9VpY+rh2z9iYHyjs0FlPO5GpdaofxnkUMv\n/dIhvw1ashe5b7Rn1ONyyQwD85k1UUGideS0LkzwRwbaSg4i2+F3X1Pl2aZihxKV\nc50O0868DjtD96ZUn6grNiyI+G5AKZgl5mQIiTcCgYEAnRFjwtJSVQf0gFwSTRgA\nfBaAZ66t4sgzaVpdJqfIYaBY/PHAW+DyirSsXX4w3LLWdhU7EdmBB9vKo8q4qbt/\n1bilrU4NkRJVrg3Z0T8KSbVD5ppfZOR8Z1pWATVBZB9XI+SuTAUVCkAd0Qx3UZzx\nAVpcEDhsIjbD0gsblCCe9T0CgYEA45rNGuZjiNnCbDL8K7Q0Za1ycZerB+06Agzs\nYOV346qiEVJFjqGzyhsTGqsM3hf5zhUtegwhAy8XtfE2IIEql5y3GKdkFqenNL2+\nnPXA1pVN0aVvI/UCa1Dsxv4tHSjMuFqbXG8tu8aRPrE2A1T01HfD+jTCBegwIVlc\nuwsiaM0CgYBMMDIAIZc4rBw7CpZeALdH8twOM+2vjoKwHHv+QCvhG1sZTodCuPIs\nftGM+1xwtspUmmc/3US3Y+jrJnDX5tId78MeybnI8MpgHR7DwDUqPPppAP1q7Hwz\nIIk9GktwKa7vc2TlGo43oepTxYMtgLFb275IMW0Qwy0QFu/qi/px3Q==\n-----END RSA PRIVATE KEY-----\n"
      },
      "validityPeriodHours": 17532
    }
  },
  "response": {
    "changes": "DIFF_SOME",
    "diffs": [
      "privateKeyPem"
    ]
  },
  "metadata": {
    "kind": "resource",
    "mode": "client",
    "name": "tls"
  }
}
`
	version.Version = "0.0.1"
	info := provider.Provider()

	readFile := func(f string) []byte {
		c, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
			return nil
		}
		return c
	}

	meta := tfbridge.ProviderMetadata{
		PackageSchema:  readFile("../provider/cmd/pulumi-resource-tls/schema.json"),
		BridgeMetadata: readFile("../provider/cmd/pulumi-resource-tls/bridge-metadata.json"),
	}

	server0, err := tfbridge.NewProvider(ctx, info, meta)
	require.NoError(t, err)

	providerServer := plugin.NewProviderServer(server0)

	testutils.Replay(t, providerServer, testCase)
}
