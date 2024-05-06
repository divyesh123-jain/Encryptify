package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// encrypt function encrypts the input file using AES encryption algorithm
func encrypt(inputFile string, outputFile string, key []byte) error {
	// Read the plaintext from the input file
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Generate a random initialization vector (IV)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Create a stream cipher in CFB mode for encryption
	stream := cipher.NewCFBEncrypter(block, iv)
	// Encrypt the plaintext and write the result to the ciphertext buffer
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// Write the ciphertext to the output file
	return os.WriteFile(outputFile, ciphertext, 0644)
}

// decrypt function decrypts the input file using AES encryption algorithm
func decrypt(inputFile string, outputFile string, key []byte) error {
	// Read the ciphertext from the input file
	ciphertext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Extract the initialization vector (IV) from the ciphertext
	if len(ciphertext) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Create a stream cipher in CFB mode for decryption
	stream := cipher.NewCFBDecrypter(block, iv)
	// Decrypt the ciphertext in place
	stream.XORKeyStream(ciphertext, ciphertext)

	// Write the decrypted plaintext to the output file
	return os.WriteFile(outputFile, ciphertext, 0644)
}

func main() {
	// Parse command-line flags
	mode := flag.String("mode", "encrypt", "Mode: 'encrypt' or 'decrypt'")
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")
	key := flag.String("key", "", "Encryption/Decryption key (16, 24, or 32 bytes)")
	flag.Parse()

	// Validate mode and input parameters
	if *mode != "encrypt" && *mode != "decrypt" {
		fmt.Println("Invalid mode. Please use 'encrypt' or 'decrypt'.")
		return
	}
	if *inputFile == "" || *outputFile == "" || *key == "" {
		fmt.Println("Please provide input file, output file, and key.")
		return
	}

	// Convert key string to bytes
	keyBytes := []byte(*key)

	var err error
	// Perform encryption or decryption based on the mode
	if *mode == "encrypt" {
		err = encrypt(*inputFile, *outputFile, keyBytes)
	} else {
		err = decrypt(*inputFile, *outputFile, keyBytes)
	}

	// Check for errors and print result
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%s successful.\n", *mode)
}
