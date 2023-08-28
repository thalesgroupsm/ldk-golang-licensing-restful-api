/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */


package ldklicensingapi

type ErrorInfo struct {
	// The error code
	ErrorCode int32 `json:"errorCode,omitempty"`
	// Descirption of error code
	Message string `json:"message,omitempty"`
}
