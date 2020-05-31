// construct the struct to evaluate the test cases.
package test

type SubnetData struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	Etag       string `json:"etag"`
	Properties struct {
		ProvisioningState                 string        `json:"provisioningState"`
		AddressPrefix                     string        `json:"addressPrefix"`
		ServiceEndpoints                  []interface{} `json:"serviceEndpoints"`
		Delegations                       []interface{} `json:"delegations"`
		PrivateEndpointNetworkPolicies    string        `json:"privateEndpointNetworkPolicies"`
		PrivateLinkServiceNetworkPolicies string        `json:"privateLinkServiceNetworkPolicies"`
	} `json:"properties"`
	Type string `json:"type"`
}

type KeyVaultData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Location string `json:"location"`
	Tags     struct {
	} `json:"tags"`
	Properties struct {
		Sku struct {
			Family string `json:"family"`
			Name   string `json:"name"`
		} `json:"sku"`
		TenantID       string `json:"tenantId"`
		AccessPolicies []struct {
			TenantID    string `json:"tenantId"`
			ObjectID    string `json:"objectId"`
			Permissions struct {
				Keys         []string `json:"keys"`
				Secrets      []string `json:"secrets"`
				Certificates []string `json:"certificates"`
			} `json:"permissions"`
		} `json:"accessPolicies"`
		EnabledForDeployment         bool   `json:"enabledForDeployment"`
		EnabledForDiskEncryption     bool   `json:"enabledForDiskEncryption"`
		EnabledForTemplateDeployment bool   `json:"enabledForTemplateDeployment"`
		VaultURI                     string `json:"vaultUri"`
	} `json:"properties"`
}