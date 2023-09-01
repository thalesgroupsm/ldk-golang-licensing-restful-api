/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingretfulapi

import (
	"time"
)

type KeyInfo struct {
	Attached           bool      `json:"attached,omitempty"`
	CloneProtected     bool      `json:"cloneProtected,omitempty"`
	Configuration      string    `json:"configuration,omitempty"`
	Detachable         bool      `json:"detachable,omitempty"`
	Disabled           bool      `json:"disabled,omitempty"`
	Driverless         bool      `json:"driverless,omitempty"`
	FingerprintChanged bool      `json:"fingerprintChanged,omitempty"`
	FormFactor         string    `json:"formFactor,omitempty"`
	HaspEnabled        bool      `json:"haspEnabled,omitempty"`
	HwPlatform         string    `json:"hwPlatform,omitempty"`
	HwVersion          string    `json:"hwVersion,omitempty"`
	KeyId              int64     `json:"keyId,omitempty"`
	KeyModel           string    `json:"keyModel,omitempty"`
	KeyType            string    `json:"keyType,omitempty"`
	ProductionDateTime time.Time `json:"productionDateTime,omitempty"`
	Recipient          bool      `json:"recipient,omitempty"`
	Rehost             bool      `json:"rehost,omitempty"`
	ResponseTime       string    `json:"responseTime,omitempty"`
	Type_              string    `json:"type,omitempty"`
	UpdateCounter      int32     `json:"updateCounter,omitempty"`
	VclockEnabled      bool      `json:"vclockEnabled,omitempty"`
	Version            string    `json:"version,omitempty"`
}

type KeysInfo struct {
	Keyinfo []KeyInfo `json:"key,omitempty"`
	Count   int32     `json:"count,omitempty"`
}

type Keys struct {
	Keyinfo KeysInfo `json:"keys,omitempty"`
}
