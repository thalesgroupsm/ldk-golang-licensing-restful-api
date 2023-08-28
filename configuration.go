/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingapi

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")

	// ContextAPIKey takes an identity as authentication for the request
	ContextIdentity = contextKey("identity")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath       string            `json:"basePath,omitempty"`
	Host           string            `json:"host,omitempty"`
	Scheme         string            `json:"scheme,omitempty"`
	DefaultHeader  map[string]string `json:"defaultHeader,omitempty"`
	UserAgent      string            `json:"userAgent,omitempty"`
	HTTPClient     *http.Client
	VendorId       string         `json:"vendorId,omitempty"`
	ClientIdentity ClientIdentity `json:"clientIdentity,omitempty"`
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://localhost:1947/sentinel/ldk_runtime/v1",
		DefaultHeader: make(map[string]string),
	}
	cfg.VendorId = "37515"
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
