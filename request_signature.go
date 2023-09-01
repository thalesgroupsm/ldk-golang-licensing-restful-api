package ldklicensingretfulapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/url"
	"time"
)

type IdentitySignature struct {
}

func (isp *IdentitySignature) hmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

func (isp *IdentitySignature) hmacSha256Ex(key []byte, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

func (isp *IdentitySignature) hmacSha256ToHex(key []byte, data string) string {
	return hex.EncodeToString(isp.hmacSha256Ex(key, data))
}

func (isp *IdentitySignature) GenerateSignatureHeader(clientIdentity IdentityAuth, endpoint string, requestBody *bytes.Buffer) (signature string) {
	//url := "/sentinel/ldk_runtime/v1/vendors/" + VENDOR_ID + "/sessions"
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	requestDate := time.Now().UTC().Format(time.RFC3339)

	deriveKey := isp.hmacSha256(clientIdentity.Secret, "X-LDK-Identity-WS-V1")
	content := clientIdentity.Id + requestDate + u.Path + "^"
	if requestBody != nil {
		content += requestBody.String()
	}
	signature = isp.hmacSha256ToHex(deriveKey, content)

	idheader := "V1, Identity=" + clientIdentity.Id + ", RequestDate=" + requestDate + ", Signature=" + signature
	return idheader

}
