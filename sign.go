package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
)

type Credential struct {
	AccessKey string
	SecretKey []byte
}

func NewCredential(accessKey, secretKey string) (*Credential, error) {
	bytes, err := base64.URLEncoding.DecodeString(secretKey)
	if err != nil {
		return nil, err
	}

	return &Credential{accessKey, bytes}, nil
}

func hashBytes(data []byte, secretKey []byte) string {
	h := hmac.New(sha1.New, secretKey)
	h.Write(data)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func HashString(s string, secretKey []byte) string {
	return hashBytes([]byte(s), secretKey)
}

func (credential *Credential) signString(s string) string {
	return HashString(s, credential.SecretKey)
}

type CommonParameter struct {
	AccessKey      string
	SignatureNonce string
	Timestamp      string
}

func prepareStringForSign(commonParameter CommonParameter, specificParameter map[string]string) string {
	commonParameterString := fmt.Sprintf("%s&%s&%s", commonParameter.AccessKey, commonParameter.SignatureNonce, commonParameter.Timestamp)

	// Sort specificParameter by Key
	var keys = make([]string, 0)
	for key, _ := range specificParameter {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var values = make([]string, 0)
	for _, key := range keys {
		values = append(values, specificParameter[key])
	}

	specificParameterString := strings.Join(values, "&")

	return fmt.Sprintf("%s&%s", commonParameterString, specificParameterString)
}

func (credential *Credential) Sign(commonParameter CommonParameter, specificParameter map[string]string) string {
	s := prepareStringForSign(commonParameter, specificParameter)
	return credential.signString(s)
}
