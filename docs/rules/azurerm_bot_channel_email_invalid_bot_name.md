<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_bot_channel_email_invalid_bot_name

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

In this rule, the string must match the regular expression `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$``.

## Example

```hcl
resource "azurerm_bot_channel_email" "foo" {
  bot_name = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." does not match valid pattern ^[a-zA-Z0-9][a-zA-Z0-9_.-]*$ (azurerm_bot_channel_email_invalid_bot_name)

  on template.tf line 2:
  2:   bot_name = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_bot_channel_email_invalid_bot_name.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json