## API Endpoints

All URIs are relative to *https://localhost:8088/sentinel/ldk_runtime/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*LicenseApi* | [**GetFeatureInfo**] | **Get** /vendors/{vendorId}/features | getFeatureInfo
*LicenseApi* | [**GetKeyInfo**] | **Get** /vendors/{vendorId}/keys | getKeyInfo
*LicenseApi* | [**GetProductInfo**] | **Get** /vendors/{vendorId}/products | getProductInfo
*LicenseApi* | [**Login**] | **Post** /vendors/{vendorId}/sessions | login
*LicenseApi* | [**Logout**] | **Delete** /vendors/{vendorId}/sessions/{sessionId} | logout
*LicenseApi* | [**Refresh**] | **Post** /vendors/{vendorId}/sessions/{sessionId}/refresh | refresh


## Authorization
Whether this header is required depends on the 'Allow Access from Remote Clients' value in the license manager server. In Sentinel Admin Control Center, this value can be found under Configuration > Access from Remote Clients.

When applying a web service signature, the expected header is similar to the following:
### Use identity
X-LDK-Identity-WS: V1, Identity=KZMSEU3, RequestDate=2015-08-30T12:36:00Z, Signature=98cd2651598ac9460e8a336912d8bf683c4690d6043ca8a51680143cde080f3c

where

V1 is a fixed string defining the version
Identity defines the identity code
RequestDate is formatted as YYYY-MM-DDTHH:MM:SSZ (20 characters)
The signature is computed as follows:

IdentitySecret = 16 bytes secret from the identity
DerivedKey = HMAC-SHA256(IdentitySecret, "X-LDK-Identity-WS-V1") (32 bytes)
Signature = HMAC-SHA256 (DerivedKey, Identity + RequestDate + Url + "^" + Body) (32 bytes)
where

Identity and RequestDate are the exact bytes that are passed in the X-LDK-Identity-WS header
Url example: "/sentinel/ldk_runtime/v1/vendors/37515/keys"
"^" ensures that Url and Body are clearly separated. Both Url and Body are invalidated if the cutoff is moved.
### Use JWT access token
X-LDK-User-Id: user id for authorization. The header should be set when using Credentials access token.

Authorization: JWT access token

where

JWT access token:
1. Credentials or Public access token from authorization server.The Login should use this access token for authorization.  
2. Licensed access token, generated when call login or refresh API. Once this access token generated, it's recommended to use it for all following calls.

## Sample
```
package main

import (
	"context"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/denisbrodbeck/machineid"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	api "github.com/thalesgroupsm/ldk-golang-licensing-restful-api"
)

type EnvCfg struct {
	VendorId        string `env:"SNTL_VENDOR_ID"         description:"Vendor Id"        long:"vendor-id"`
	ClientIdentity  string `env:"SNTL_CLIENT_IDENTITY"   description:"Client Identity"  long:"client-identity"`
	EndpointScheme  string `env:"SNTL_ENDPOINT_SCHEME"   description:"Endpoint Scheme"  long:"endpoint-scheme"`
	ServerAddr      string `env:"SNTL_SERVER_ADDR"   description:"Server Address"  long:"servver-address"`
	ServerPort      string `env:"SNTL_SERVER_PORT"   description:"Server Port"  long:"server-port"`
	Proxy           string `env:"SNTL_PROXY"   description:"Proxy"  long:"proxy"`
	AccessToken     string `env:"SNTL_ACCESS_TOKEN"   description:"Access Token"  long:"access-token"`
	AccessTokenType int    `env:"SNTL_ACCESS_TOKEN_TYPE"   description:"Access Token Type"  long:"access-token-type"`
	UserId          string `env:"SNTL_USER_ID"   description:"User ID"  long:"user-id"`
}

var env EnvCfg

func main() {

	// parse & validate environment variables
	godotenv.Load()
	flags.Parse(&env)

	var authCtx context.Context
	if env.AccessToken == "" {
		// parse the client identity
		clientIdResult := strings.Split(env.ClientIdentity, ":")
		if clientIdResult == nil || len(clientIdResult) != 2 {
			log.Fatal("Client Identity is not valid")
			return
		}
		authCtx = context.WithValue(context.Background(), api.ContextIdentity, api.IdentityAuth{
			Id:     clientIdResult[0],
			Secret: clientIdResult[1],
		})
	} else {
		authCtx = context.WithValue(context.Background(), api.ContextAccessToken, api.AccessTokenAuth{
			UserId:          env.UserId,
			AccessToken:     env.AccessToken,
			AccessTokenType: env.AccessTokenType,
		})
	}

	cfg := &api.Configuration{
		Host:     env.ServerAddr,
		VendorId: env.VendorId,
		Scheme:   env.EndpointScheme,
		BasePath: env.EndpointScheme + "://" + env.ServerAddr + ":" + env.ServerPort + "/sentinel/ldk_runtime/v1",
	}

	licensingApiClient := api.NewAPIClient(cfg)
	licenseRequest := api.LicenseRequest{}
	licenseRequest.FeatureId = 0

	// get client info
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	licenseRequest.ClientInfo = &api.ClientInfo{}
	licenseRequest.ClientInfo.MachineId, _ = machineid.ID()
	licenseRequest.ClientInfo.UserName = user.Username
	licenseRequest.ClientInfo.DomainName, _ = os.Hostname()
	licenseRequest.ClientInfo.ProcessId = strconv.Itoa(os.Getpid())
	licenseRequest.ClientInfo.ClientDateTime = time.Now().UTC().Format(time.RFC3339)

	apiResponse, _, err := licensingApiClient.LicenseApi.Login(authCtx, licenseRequest)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("licensingApi.LicenseApi.Login %#v", apiResponse)

	if env.AccessToken != "" {
		//use LAT access token after login success
		authCtx = context.WithValue(context.Background(), api.ContextAccessToken, api.AccessTokenAuth{
			UserId:          "",
			AccessToken:     apiResponse.LmAccessToken,
			AccessTokenType: 1,
		})
	}

	localVarOptionals := &api.QueryInfoOpts{
		PageStartIndex: optional.NewInt32(0),
		PageSize:       optional.NewInt32(1),
	}
	keys, _, err := licensingApiClient.LicenseApi.GetKeyInfo(authCtx, localVarOptionals)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("licensingApi.LicenseApi.GetKeyInfo %#v", keys)

	products, _, err := licensingApiClient.LicenseApi.GetProductInfo(authCtx, localVarOptionals)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("licensingApi.LicenseApi.GetProductInfo %#v", products)

	features, _, err := licensingApiClient.LicenseApi.GetFeatureInfo(authCtx, localVarOptionals)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("licensingApi.LicenseApi.GetFeatureInfo %#v", features)

	licenseResponse, _, err := licensingApiClient.LicenseApi.Refresh(authCtx, apiResponse.SessionId)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("licensingApi.LicenseApi.Refresh  %#v", licenseResponse)

	_, err = licensingApiClient.LicenseApi.Logout(authCtx, apiResponse.SessionId)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("licensingApi.LicenseApi.Logout", apiResponse.SessionId)
}
```
