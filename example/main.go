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
		Proxy:    env.Proxy,
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
	licenseRequest.ClientInfo.HostName, _ = os.Hostname()
	licenseRequest.ClientInfo.DomainName = "test"
	licenseRequest.ClientInfo.ProcessId = strconv.Itoa(os.Getpid())
	licenseRequest.ClientInfo.ClientDateTime = time.Now().UTC().Format(time.RFC3339)

	apiResponse, _, err := licensingApiClient.LicenseApi.Login(authCtx, licenseRequest)
	if err != nil {
		log.Printf("error %#v", apiResponse)
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
