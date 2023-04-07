package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

const key string = "0123456789abcdef0123456789abcdef"

func main() {
	// Read in the encoded file
	inputPath := "C:/Users/bot/Documents/gostuff/output.txt"
	inputFile, err := ioutil.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	// Decode the encoded file from base64 to a byte slice
	decodedInputFile, err := base64.StdEncoding.DecodeString(string(inputFile))
	if err != nil {
		panic(err)
	}

	// Extract the IV and encrypted file from the decoded input file
	iv := decodedInputFile[:aes.BlockSize]
	encryptedFile := decodedInputFile[aes.BlockSize:]

	// Convert the key string to a byte slice
	keyBytes := []byte(key)

	// Create a new AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// Create a new stream cipher using AES in CFB mode
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt the encrypted file using the stream cipher
	decryptedFile := make([]byte, len(encryptedFile))
	stream.XORKeyStream(decryptedFile, encryptedFile)

	// Write the decrypted file to a new file
	outputPath := "C:/Users/bot/Documents/gostuff/decrypted-file.exe"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	_, err = writer.Write(decryptedFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decryption successful!")
}
