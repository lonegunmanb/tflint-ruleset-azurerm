// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermBatchCertificateInvalidAccountNameRule checks the pattern is valid
type AzurermBatchCertificateInvalidAccountNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermBatchCertificateInvalidAccountNameRule returns new rule with default attributes
func NewAzurermBatchCertificateInvalidAccountNameRule() *AzurermBatchCertificateInvalidAccountNameRule {
	return &AzurermBatchCertificateInvalidAccountNameRule{
		resourceType:  "azurerm_batch_certificate",
		attributeName: "account_name",
		pattern:       regexp.MustCompile(`^[-\w\._]+$`),
	}
}

// Name returns the rule name
func (r *AzurermBatchCertificateInvalidAccountNameRule) Name() string {
	return "azurerm_batch_certificate_invalid_account_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermBatchCertificateInvalidAccountNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermBatchCertificateInvalidAccountNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermBatchCertificateInvalidAccountNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermBatchCertificateInvalidAccountNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[-\w\._]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
