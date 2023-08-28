/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingapi

type ProductInfo struct {
	CloneProtected bool `json:"cloneProtected,omitempty"`
	// Whether the license for the Product can be detached from the network pool.
	Detachable bool `json:"detachable,omitempty"`
	FingerprintChanged bool `json:"fingerprintChanged,omitempty"`
	ProductId int32 `json:"productId,omitempty"`
	KeyId int64 `json:"keyId,omitempty"`
	Name string `json:"name,omitempty"`
}
