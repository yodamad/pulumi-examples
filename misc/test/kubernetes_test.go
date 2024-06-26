//go:build Kubernetes || all
// +build Kubernetes all

package test

import (
	"path"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pulumi/examples/misc/test/helpers"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestAccKubernetesGuestbook(t *testing.T) {
	_, err := homedir.Expand("~/.kube/config")
	if err != nil {
		t.Skipf("Missing KubeConfig to run test: %s", err)
	}

	tests := []integration.ProgramTestOptions{
		integration.ProgramTestOptions{
			Dir:        path.Join(getCwd(t), "..", "..", "kubernetes-go-guestbook", "simple"),
			NoParallel: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				endpoint := stack.Outputs["frontendIP"].(string)
				helpers.AssertHTTPResult(t, endpoint, nil, func(body string) bool {
					return assert.Contains(t, body, "Guestbook")
				})
			},
		},
		integration.ProgramTestOptions{
			Dir:        path.Join(getCwd(t), "..", "..", "kubernetes-go-guestbook", "components"),
			NoParallel: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				endpoint := stack.Outputs["frontendIP"].(string)
				helpers.AssertHTTPResult(t, endpoint, nil, func(body string) bool {
					return assert.Contains(t, body, "Guestbook")
				})
			},
		},
		integration.ProgramTestOptions{
			Dir:        path.Join(getCwd(t), "..", "..", "kubernetes-csharp-guestbook", "simple"),
			NoParallel: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				endpoint := stack.Outputs["FrontendIp"].(string)
				helpers.AssertHTTPResult(t, endpoint, nil, func(body string) bool {
					return assert.Contains(t, body, "Guestbook")
				})
			},
		},
		integration.ProgramTestOptions{
			Dir:        path.Join(getCwd(t), "..", "..", "kubernetes-csharp-guestbook", "components"),
			NoParallel: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				endpoint := stack.Outputs["FrontendIp"].(string)
				helpers.AssertHTTPResult(t, endpoint, nil, func(body string) bool {
					return assert.Contains(t, body, "Guestbook")
				})
			},
		},
		integration.ProgramTestOptions{
			Dir:        path.Join(getCwd(t), "..", "..", "kubernetes-py-guestbook", "simple"),
			NoParallel: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				endpoint := stack.Outputs["frontend_ip"].(string)
				helpers.AssertHTTPResult(t, endpoint, nil, func(body string) bool {
					return assert.Contains(t, body, "Guestbook")
				})
			},
		},
		integration.ProgramTestOptions{
			Dir:        path.Join(getCwd(t), "..", "..", "kubernetes-ts-guestbook", "simple"),
			NoParallel: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				endpoint := stack.Outputs["frontendIp"].(string)
				helpers.AssertHTTPResult(t, endpoint, nil, func(body string) bool {
					return assert.Contains(t, body, "Guestbook")
				})
			},
		},
		integration.ProgramTestOptions{
			Dir:        path.Join(getCwd(t), "..", "..", "kubernetes-ts-guestbook", "components"),
			NoParallel: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				endpoint := stack.Outputs["frontendIp"].(string)
				helpers.AssertHTTPResult(t, endpoint, nil, func(body string) bool {
					return assert.Contains(t, body, "Guestbook")
				})
			},
		},
	}

	for _, ex := range tests {
		example := ex
		t.Run(example.Dir, func(t *testing.T) {
			t.Log(example.StackName)
			integration.ProgramTest(t, &example)
		})
	}
}
