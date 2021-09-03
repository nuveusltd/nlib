package nlib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	cRand "crypto/rand"
	"encoding/hex"
	"io"
	"strings"
)

func maxof2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minof2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(numbers ...int) int {
	m := 0
	for _, i := range numbers {
		m = maxof2(m, i)
	}
	return m
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
}

func Encrypt(data []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(cRand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

// XORBytes takes two byte slices and XORs them together, returning the final
// byte slice.
func XOR2Bytes(b1, b2 []byte) []byte {
	oneAndTwo := make([]byte, max(len(b1), len(b2)))

	for i := 0; i < len(b1) && i < len(b2); i++ {
		oneAndTwo[i] = b1[i] ^ b2[i]
	}
	for i := len(b2); i < len(b1); i++ {
		oneAndTwo[i] = b1[i]
	}
	for i := len(b1); i < len(b2); i++ {
		oneAndTwo[i] = b2[i]
	}
	return oneAndTwo

}

// XORBytes takes two byte slices and XORs them together, returning the final
// byte slice.
func XORBytes(buffers ...[]byte) []byte {
	maxLen := 0
	for _, buff := range buffers {
		maxLen = maxof2(maxLen, len(buff))
	}
	minLen := ^int(0)
	for _, buff := range buffers {
		minLen = minof2(minLen, len(buff))
	}
	result := make([]byte, maxLen)

	for i := 0; i < minLen; i++ {
		for bi, buff := range buffers {
			if bi == 0 {
				result[i] = buff[i]
			} else {
				result[i] = result[i] ^ buff[i]
			}
		}
	}
	for i := minLen - 1; i < maxLen; i++ {
		for bi := 0; bi < len(buffers); bi++ {
			if i < len(buffers[bi]) {
				if result[i] == byte(0) {
					result[i] = buffers[bi][i]
				} else {
					result[i] = result[i] ^ buffers[bi][i]
				}
			}
		}

	}
	return result
}
