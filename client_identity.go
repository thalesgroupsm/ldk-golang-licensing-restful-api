package ldklicensingapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func getCurrentTime() (ret string) {

	year := time.Now().UTC().Format("2006")
	month := time.Now().UTC().Format("01")
	day := time.Now().UTC().Format("02")
	hour := time.Now().UTC().Format("15")
	min := time.Now().UTC().Format("04")
	second := time.Now().UTC().Format("05")

	ret = year + "-" + month + "-" + day + "T" + hour + ":" + min + ":" + second + "Z"

	return ret
}

func hmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

func hmacSha256Ex(key []byte, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

func hmacSha256ToHex(key []byte, data string) string {
	return hex.EncodeToString(hmacSha256Ex(key, data))
}

func GenerateSignatureHeader(clientIdentity ClientIdentity, endpoint string, requestBody string) (signature string) {
	//url := "/sentinel/ldk_runtime/v1/vendors/" + VENDOR_ID + "/sessions"

	requestDate := getCurrentTime()

	deriveKey := hmacSha256(clientIdentity.Secret, "X-LDK-Identity-WS-V1")
	content := clientIdentity.id + requestDate + endpoint + "^" + string(requestBody)
	signature = hmacSha256ToHex(deriveKey, content)

	idheader := "V1, Identity=" + clientIdentity.id + ", RequestDate=" + requestDate + ", Signature=" + signature
	return idheader

}
