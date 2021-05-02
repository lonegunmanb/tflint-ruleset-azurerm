mapping "azurerm_cosmosdb_table" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-01-15/cosmos-db.json"

  name                = tableNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}
