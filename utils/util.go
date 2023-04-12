package utils

import (
	"crypto/rand"
	"math/big"

	"github.com/vasuvanka/machinemax-assignment/definitions"
)

const allowedChars = "ABCDEF0123456789"

func generateHexString(length int) (string, error) {
	max := big.NewInt(int64(len(allowedChars)))
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b[i] = allowedChars[n.Int64()]
	}
	return string(b), nil
}

func GenerateDevEUIBatch(limit int) ([]definitions.DevEUI, error) {
	list := make([]definitions.DevEUI, limit)
	codeMap := make(map[string]bool)
	for index := 0; index < limit; index += 1 {
		uid, err := generateHexString(16)
		if err != nil {
			return nil, err
		}
		// if duplicate found subtract one from index
		if found := codeMap[uid]; found {
			index -= 1
			continue
		}
		list[index] = definitions.DevEUI(uid)
	}
	return list, nil
}
