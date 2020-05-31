package test

import (
	"encoding/json"
	"fmt"
	"flag"
	"os"
	"net/http"
	"net/url"
	)
const grant_type string = "client_credentials"
const resource string = "https://management.azure.com/"
var clientId = flag.String("clientId", os.Getenv("ARM_CLIENT_ID"), "Client Id" )
var clientSecret = flag.String("clientSecret", os.Getenv("ARM_CLIENT_SECRET"), "Client Secret" )
var tenant_id  = flag.String("tenantId", os.Getenv("ARM_TENANT_ID"), "Tenant Id")
var api_endpoint = flag.String("api_endpoint", fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", *tenant_id), "api endpoint" )

//generate the token to access the Azure rest api
func get_access_token() string{
	flag.Parse()
	formData := url.Values{
		"grant_type" : {grant_type},
		"client_id" : {*clientId},
		"client_secret" : {*clientSecret},
		"resource"	: {resource},
	}
	resp, err := http.PostForm(*api_endpoint, formData)
	if err != nil {
		fmt.Printf("Error for http.PostForm() %s\n", err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["access_token"].(string)

}
