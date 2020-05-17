<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_stream_analytics_job_invalid_events_out_of_order_policy

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- Adjust
- Drop

## Example

```hcl
resource "azurerm_stream_analytics_job" "foo" {
  events_out_of_order_policy = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as events_out_of_order_policy (azurerm_stream_analytics_job_invalid_events_out_of_order_policy)

  on template.tf line 2:
  2:   events_out_of_order_policy = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_stream_analytics_job_invalid_events_out_of_order_policy.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/streamingjobs.json