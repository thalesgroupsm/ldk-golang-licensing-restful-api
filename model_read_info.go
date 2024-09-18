/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingretfulapi

type ReadInfo struct {
	MemoryID int32 `json:"memoryId"`
	Offset   int32 `json:"offset"`
	Length   int32 `json:"length"`
}

type MemoryInfo struct {
	Content string `json:"content"`
}
