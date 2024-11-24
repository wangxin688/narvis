// Copyright 2024 wangxin.jeffry@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package security cryptocode used for psk encryption and decryption
// secretKey should be the same between server and agent and do not shared.

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
)

// PskEncrypt encrypts the given string using the secret key.
// The encrypted string is then returned as a hex encoded string.
func PskEncrypt(psk string, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey[:32]))
	if err != nil {
		return "", err
	}
	encrypted := make([]byte, aes.BlockSize+len(psk))

	// Generate a random IV and store it in the first block size bytes of the buffer.
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	// Create a new cipher stream from the block and IV.
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt the plaintext and store it in the buffer.
	stream.XORKeyStream(encrypted[aes.BlockSize:], []byte(psk))

	// Return the encrypted buffer as a hex encoded string.
	return hex.EncodeToString(encrypted), nil
}

// PskDecrypt decrypts the given encrypted string using the secret key.
// The decrypted string is then returned.
func PskDecrypt(encrypted string, secretKey string) (string, error) {
	// Decode the encrypted string from hex to bytes.
	decoded, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("failed to decode encrypted string: %w", err)
	}
	// Create a new AES cipher block from the secret key.
	block, err := aes.NewCipher([]byte(secretKey[:32]))
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher block: %w", err)
	}
	// Check if the decoded bytes are at least one block size long.
	if len(decoded) < aes.BlockSize {
		return "", fmt.Errorf("invalid encrypted string: length is less than one block size")
	}
	// Extract the IV from the decoded bytes.
	iv := decoded[:aes.BlockSize]
	// Remove the IV from the decoded bytes.
	decoded = decoded[aes.BlockSize:]
	// Create a new cipher stream from the block and IV.
	stream := cipher.NewCFBDecrypter(block, iv)
	// Decrypt the decoded bytes using the cipher stream.
	stream.XORKeyStream(decoded, decoded)
	// Return the decrypted string.
	return string(decoded), nil
}

// VerifyAgentToken verifies the given agent token and returns the unique agent id and agent psk.
// If the token is invalid, an error is returned.
func VerifyAgentToken(token string, secretKey string) (agentId, agentPsk string, err error) {
	// Decrypt the token using the secret key.
	_token, err := PskDecrypt(token, secretKey)
	if err != nil {
		// Return the error if the decryption fails.
		return "", "", err
	}
	// Parse the decrypted token to get the proxy id and proxy psk.
	agentId, agentPsk, err = parseAgentToken(_token)
	// Return the proxy id, proxy psk and the error.
	return
}

func parseAgentToken(token string) (agentId, agentPsk string, err error) {
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

func GenerateAgentToken(agentId, agentPsk, secretKey string) (string, error) {

	token, err := PskEncrypt(agentId+","+agentPsk, secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
