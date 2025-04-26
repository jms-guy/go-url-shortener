package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// Algorithm for creating a random, unique short url

//Create a sha256 hash based off input url string
func hash256(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	
	return algorithm.Sum(nil)
}

//Encode created hash using base58 encoding
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Printf("Error base58 encoding initial url: %v", err)
		return ""
	}
	return string(encoded)
}

//Generate a new short url based off of an initial url input
func GenerateShortLink(initialUrl string) string {
	urlHashBytes := hash256(initialUrl)
	generatedNum := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNum)))

	return finalString[:8]
}

func SetNewUrl(short string) string {
	domain := os.Getenv("SERVER_DOMAIN")
	port := os.Getenv("PORT")
	url := fmt.Sprintf("%s:%v/%s", domain, port, short)
	return url
}
