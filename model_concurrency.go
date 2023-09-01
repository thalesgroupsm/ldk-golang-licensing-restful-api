/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingretfulapi

type Concurrency struct {
	Count string `json:"count,omitempty"`
	Export string `json:"export,omitempty"`
	Seats int32 `json:"seats,omitempty"`
}
