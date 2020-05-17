<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_sql_elasticpool_invalid_edition

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- Basic
- Standard
- Premium
- GeneralPurpose
- BusinessCritical

## Example

```hcl
resource "azurerm_sql_elasticpool" "foo" {
  edition = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as edition (azurerm_sql_elasticpool_invalid_edition)

  on template.tf line 2:
  2:   edition = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_sql_elasticpool_invalid_edition.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/elasticPools.json