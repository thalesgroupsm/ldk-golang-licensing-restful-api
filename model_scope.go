/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */


package ldklicensingretfulapi

type Scope struct {
	KeyId []int32 `json:"keyId,omitempty"`
	ProductId []int32 `json:"productId,omitempty"`
}
