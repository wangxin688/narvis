package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/wangxin688/narvis/client/config"
)


func PskEncrypt(psk string) (string, error) {
	block, err := aes.NewCipher([]byte(config.Settings.SECRET_KEY[:32]))
	if err != nil {
		return "", err
	}
	encrypted := make([]byte, aes.BlockSize+len(psk))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], []byte(psk))
	return hex.EncodeToString(encrypted), nil
}

func PskDecrypt(encrypted string) (string, error) {
	decoded, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(config.Settings.SECRET_KEY[:32]))
	if err != nil {
		return "", err
	}
	if len(decoded) < aes.BlockSize {
		return "", err
	}
	iv := decoded[:aes.BlockSize]
	decoded = decoded[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decoded, decoded)
	return string(decoded), nil
}

func VerifyProxyToken(token string) (proxyId, proxyPsk string, err error) {

	_token, err := PskDecrypt(token)
	if err != nil {
		return "", "", err
	}
	proxyId, proxyPsk, err = parseProxyToken(_token)
	return
}

func parseProxyToken(token string) (proxyId, proxyPsk string, err error) {
	if token == "" {
		return "", "", nil
	}
	parts := strings.Split(token, ",")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid proxy token")
	}
	_, err = uuid.Parse(parts[0])
	if err != nil {
		return "", "", err
	}
	return parts[0], parts[1], nil
}

func ProxyToken(proxyId, proxyPsk string) (string, error) {

	token, err := PskEncrypt(proxyId + "," + proxyPsk)
	if err != nil {
		return "", err
	}
	return token, nil
}
