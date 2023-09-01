/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingretfulapi

type LicenseRequest struct {
	FeatureId               int32       `json:"featureId"`
	ClientInfo              *ClientInfo `json:"clientInfo,omitempty"`
	Scope                   *Scope      `json:"scope,omitempty"`
	DieAtExpiration         bool        `json:"dieAtExpiration,omitempty"`
	ExecutionCountToConsume int32       `json:"executionCountToConsume,omitempty"`
	NetworkSeatsToConsume   int32       `json:"networkSeatsToConsume,omitempty"`
}
