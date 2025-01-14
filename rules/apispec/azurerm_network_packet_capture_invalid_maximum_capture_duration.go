// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule checks the pattern is valid
type AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule returns new rule with default attributes
func NewAzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule() *AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule {
	return &AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule{
		resourceType:  "azurerm_network_packet_capture",
		attributeName: "maximum_capture_duration",
		max:           18000,
		min:           0,
	}
}

// Name returns the rule name
func (r *AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule) Name() string {
	return "azurerm_network_packet_capture_invalid_maximum_capture_duration"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermNetworkPacketCaptureInvalidMaximumCaptureDurationRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssue(
					r,
					"maximum_capture_duration must be 18000 or less",
					attribute.Expr.Range(),
				)
			}
			if val < r.min {
				runner.EmitIssue(
					r,
					"maximum_capture_duration must be 0 or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
