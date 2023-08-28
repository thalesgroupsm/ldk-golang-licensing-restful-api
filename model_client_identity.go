/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingapi

type ClientIdentity struct {
	id     string `json:"domainName,omitempty"`
	Secret string `json:"secret,omitempty"`
}
