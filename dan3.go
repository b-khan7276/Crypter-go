package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

const key string = "0123456789abcdef0123456789abcdef"

func main() {
	// Read in the executable file
	inputPath := "C:/Users/bot/Documents/gostuff/file.exe"
	inputFile, err := ioutil.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	// Convert the key string from hex to byte slice
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		panic(err)
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// Create a new stream cipher using AES in CFB mode
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt the executable file using the stream cipher
	encryptedFile := make([]byte, len(inputFile))
	stream.XORKeyStream(encryptedFile, inputFile)

	// Concatenate the IV and encrypted file into a single byte slice
	output := append(iv, encryptedFile...)

	// Encode the byte slice using base64
	encodedOutput := base64.StdEncoding.EncodeToString(output)

	// Write the encoded output to a new file
	outputPath := "C:/Users/bot/Documents/gostuff/output.txt"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	_, err = writer.WriteString(encodedOutput)
	if err != nil {
		panic(err)
	}

	fmt.Println("Encryption successful!")
}
