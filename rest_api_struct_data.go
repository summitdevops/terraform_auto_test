// construct the struct to evaluate the test cases.
package terratest

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


type VNet struct {
	APIVersion string `json:"apiVersion"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		AddressSpace struct {
			AddressPrefixes []string `json:"addressPrefixes"`
		} `json:"addressSpace"`
		DhcpOptions struct {
			DNSServers []interface{} `json:"dnsServers"`
		} `json:"dhcpOptions"`
		EnableDdosProtection bool `json:"enableDdosProtection"`
		EnableVMProtection   bool `json:"enableVmProtection"`
		Subnets              []struct {
			Name       string `json:"name"`
			Properties struct {
				AddressPrefix                     string        `json:"addressPrefix"`
				Delegations                       []interface{} `json:"delegations"`
				PrivateEndpointNetworkPolicies    string        `json:"privateEndpointNetworkPolicies"`
				PrivateLinkServiceNetworkPolicies string        `json:"privateLinkServiceNetworkPolicies"`
				ServiceEndpoints                  []interface{} `json:"serviceEndpoints"`
			} `json:"properties"`
		} `json:"subnets"`
		VirtualNetworkPeerings []interface{} `json:"virtualNetworkPeerings"`
	} `json:"properties"`
	Type string `json:"type"`
}


type Storage struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	APIVersion string `json:"apiVersion"`
	Sku        struct {
		Name         string `json:"name"`
		Restrictions []struct {
			ReasonCode string `json:"reasonCode"`
		} `json:"restrictions"`
	} `json:"sku"`
	Kind     string `json:"kind"`
	Location string `json:"location"`
	Tags     struct {
	} `json:"tags"`
	Identity struct {
		Type string `json:"type"`
	} `json:"identity"`
	Properties struct {
		CustomDomain struct {
			Name             string `json:"name"`
			UseSubDomainName string `json:"useSubDomainName"`
		} `json:"customDomain"`
		Encryption struct {
			Services struct {
				Blob struct {
					Enabled string `json:"enabled"`
				} `json:"blob"`
				File struct {
					Enabled string `json:"enabled"`
				} `json:"file"`
			} `json:"services"`
			KeySource          string `json:"keySource"`
			Keyvaultproperties struct {
				Keyname     string `json:"keyname"`
				Keyversion  string `json:"keyversion"`
				Keyvaulturi string `json:"keyvaulturi"`
			} `json:"keyvaultproperties"`
		} `json:"encryption"`
		NetworkAcls struct {
			Bypass              string `json:"bypass"`
			VirtualNetworkRules []struct {
				ID     string `json:"id"`
				Action string `json:"action"`
				State  string `json:"state"`
			} `json:"virtualNetworkRules"`
			IPRules []struct {
				Value  string `json:"value"`
				Action string `json:"action"`
			} `json:"ipRules"`
			DefaultAction string `json:"defaultAction"`
		} `json:"networkAcls"`
		AccessTier                            string `json:"accessTier"`
		AzureFilesIdentityBasedAuthentication struct {
			DirectoryServiceOptions   string `json:"directoryServiceOptions"`
			ActiveDirectoryProperties struct {
				DomainName        string `json:"domainName"`
				NetBiosDomainName string `json:"netBiosDomainName"`
				ForestName        string `json:"forestName"`
				DomainGUID        string `json:"domainGuid"`
				DomainSid         string `json:"domainSid"`
				AzureStorageSid   string `json:"azureStorageSid"`
			} `json:"activeDirectoryProperties"`
		} `json:"azureFilesIdentityBasedAuthentication"`
		SupportsHTTPSTrafficOnly string `json:"supportsHttpsTrafficOnly"`
		IsHnsEnabled             string `json:"isHnsEnabled"`
		LargeFileSharesState     string `json:"largeFileSharesState"`
	} `json:"properties"`
	Resources []interface{} `json:"resources"`
}

/*
type SubnetData struct {
		Name       string `json:"name"`
		Properties struct {
			AddressPrefix                     string        `json:"addressPrefix"`
			Delegations                       []interface{} `json:"delegations"`
			PrivateEndpointNetworkPolicies    string        `json:"privateEndpointNetworkPolicies"`
			PrivateLinkServiceNetworkPolicies string        `json:"privateLinkServiceNetworkPolicies"`
			ServiceEndpoints                  []interface{} `json:"serviceEndpoints"`
		} `json:"properties"`
	}

 */