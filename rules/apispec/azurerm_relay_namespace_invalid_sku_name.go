// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermRelayNamespaceInvalidSkuNameRule checks the pattern is valid
type AzurermRelayNamespaceInvalidSkuNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermRelayNamespaceInvalidSkuNameRule returns new rule with default attributes
func NewAzurermRelayNamespaceInvalidSkuNameRule() *AzurermRelayNamespaceInvalidSkuNameRule {
	return &AzurermRelayNamespaceInvalidSkuNameRule{
		resourceType:  "azurerm_relay_namespace",
		attributeName: "sku_name",
		enum: []string{
			"Standard",
		},
	}
}

// Name returns the rule name
func (r *AzurermRelayNamespaceInvalidSkuNameRule) Name() string {
	return "azurerm_relay_namespace_invalid_sku_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermRelayNamespaceInvalidSkuNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermRelayNamespaceInvalidSkuNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermRelayNamespaceInvalidSkuNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermRelayNamespaceInvalidSkuNameRule) Check(runner tflint.Runner) error {
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
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sku_name`, truncateLongMessage(val)),
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
