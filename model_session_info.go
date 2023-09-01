/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingretfulapi

type SessionInfo struct {
	// Number of current logins to the Feature/application.
	Currentlogins int32 `json:"currentlogins,omitempty"`
	// Total amount of time (in seconds) that the session can be idle before it is terminated.
	IdleTimeout int32 `json:"idleTimeout,omitempty"`
}
