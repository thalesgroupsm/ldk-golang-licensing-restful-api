/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingretfulapi

type ClientInfo struct {
	DomainName string `json:"domainName,omitempty"`
	HostName   string `json:"hostName,omitempty"`
	UserName   string `json:"userName,omitempty"`
	MachineId  string `json:"machineId,omitempty"`
	ProcessId  string `json:"processId,omitempty"`
	// client side date and time
	ClientDateTime string `json:"clientDateTime,omitempty"`
}
