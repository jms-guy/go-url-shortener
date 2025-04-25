package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"github.com/itchyny/base58-go"
)

func hash256(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Printf("Error base58 encoding initial url: %v", err)
		return ""
	}
	return string(encoded)
}

func GenerateShortLink(initialUrl string) string {
	urlHashBytes := hash256(initialUrl)
	generatedNum := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNum)))

	return finalString[:8]
}
