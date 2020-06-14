package terratest

import (
	"fmt"
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)
// evaluate the test cases
func TestModule(t *testing.T)  {
	t.Parallel()
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}
	//defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	keyVaultData := KeyVaultData{} //KeyVaultData struct
	subnetData := SubnetData{}	//SubnetData struct
	vnetData := VNet{}	//VNet struct
	storageData := Storage{}

	get_rest_data(keyVaultApi, "GET", &keyVaultData)
	get_rest_data(vnetApi, "GET", &vnetData)
	get_rest_data(subnetApi, "GET", &subnetData)
	get_rest_data(storageApi, "GET", &storageData)
	if assert.Equal(t, subnetData.Name, "mysubnet"){
		fmt.Println("Subnet Name test passed!")
	}
	if assert.Equal(t, subnetData.Properties.ProvisioningState, "Succeeded"){
		fmt.Printf("%s is provisioned\n", subnetData.Name)
	} else {
		defer terraform.Destroy(t, terraformOptions)
	}
	assert.Equal(t, subnetData.Properties.ProvisioningState, "Succeeded")
	assert.Equal(t, subnetData.Properties.AddressPrefix, "10.0.1.0/24")
	assert.Equal(t, vnetData.Properties.AddressSpace.AddressPrefixes[0], "10.0.0.0/16")
	assert.Equal(t, keyVaultData.Name, "agwkeyvaulttest")
	assert.Equal(t, keyVaultData.Location, "eastus")
	assert.Equal(t, keyVaultData.Properties.Sku.Name, "standard")
	for _, v := range storageData.Properties.NetworkAcls.IPRules{
		assert.Equal(t, v.Value, "100.100.100.100")
	}
	for _, v := range storageData.Properties.NetworkAcls.VirtualNetworkRules{
		assert.Equal(t, v.Action, "Allow")
	}
	if storageData.Properties.NetworkAcls.DefaultAction == "Allow"{
		fmt.Printf("Test Passed , DefaultAction is set to Allow")

	}		else{
		fmt.Printf("DefaultAction Test failed , DefaultAction is set to %s, Destroying the resource.", storageData.Properties.NetworkAcls.DefaultAction)
		defer terraform.Destroy(t, terraformOptions)
	}

}
