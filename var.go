//define all the variables
package terratest

import (
	"flag"
	"fmt"
	"os"
)
//token.go variables
const grant_type string = "client_credentials"
const resource string = "https://management.azure.com/"
var client_id = flag.String("clientId", os.Getenv("ARM_CLIENT_ID"), "Client Id" )
var client_secret = flag.String("clientSecret", os.Getenv("ARM_CLIENT_SECRET"), "Client Secret" )
var tenant_id  = flag.String("tenantId", os.Getenv("ARM_TENANT_ID"), "Tenant Id")
var api_endpoint = flag.String("api_endpoint", fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", *tenant_id), "api endpoint" )

//get_rest_data.go variables
var subscription_id = flag.String("subscriptionId", os.Getenv("ARM_SUBSCRIPTION_ID"), "Subscription Id" )
var resourceGroup string = "agtest"


//VNet variables
var virtualNetwork string = "agvnet"
var vnetApi string = fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s?api-version=2020-04-01", *subscription_id, resourceGroup, virtualNetwork)


//Subnet variables
//var subnetData []SubnetData
var subnet string = "mysubnet"
var subnetApi string = fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s?api-version=2020-04-01", *subscription_id, resourceGroup, virtualNetwork, subnet)

// KeyVault variables
var kevVaultName string = "agwkeyvaulttest"
var keyVaultApi string = fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.KeyVault/vaults/%s?api-version=2019-09-01",*subscription_id, resourceGroup, kevVaultName)

//Storage account variables
//var storageAccount string = "teststorageaccountblob"
var storageApi string = "https://management.azure.com/subscriptions/52ba0546-5dab-4191-b1e6-d9177b9e5eda/resourceGroups/agtest/providers/Microsoft.Storage/storageAccounts/teststorageaccountblob?api-version=2018-02-01"