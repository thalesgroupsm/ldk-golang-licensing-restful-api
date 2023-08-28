/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingapi

import (
	"time"
)

type LicenseInfo struct {
	TotalCounter int32 `json:"totalCounter,omitempty"`
	CurrentCounter int32 `json:"currentCounter,omitempty"`
	ExpirationDateTime time.Time `json:"expirationDateTime,omitempty"`
	StartDateTime time.Time `json:"startDateTime,omitempty"`
	TotalDuration int32 `json:"totalDuration,omitempty"`
	Type_ string `json:"type,omitempty"`
}
