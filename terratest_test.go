package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)
// evaluate the test cases
func TestModule(t *testing.T)  {
	t.Parallel()
	get_rest_data(subnetApi, "GET", &subnetData)
	get_rest_data(keyVaultApi, "GET", &keyVaultData)
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}
	//defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	assert.Equal(t, subnetData.Name, "mysubnet")
	assert.Equal(t, subnetData.Properties.ProvisioningState, "Succeeded")
	assert.Equal(t, keyVaultData.Name, "agwkeyvaulttest")
	assert.Equal(t, keyVaultData.Location, "eastus")
	assert.Equal(t, keyVaultData.Properties.Sku.Name, "standard")

}