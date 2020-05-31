//define all the variables
package test

import (
	"flag"
	"fmt"
	"os"
)
var subnetData = SubnetData{}
var keyVaultData = KeyVaultData{}
var resourceGroup string = "agtest"
var virtualNetwork string = "agvnet"
var subnet string = "mysubnet"
var kevVaultName string = "agwkeyvaulttest"
var subscription_id = flag.String("subscriptionId", os.Getenv("ARM_SUBSCRIPTION_ID"), "Subscription Id" )
var subnetApi = fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s?api-version=2020-04-01", *subscription_id, resourceGroup, virtualNetwork, subnet)
var keyVaultApi = fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.KeyVault/vaults/%s?api-version=2019-09-01",*subscription_id, resourceGroup, kevVaultName)