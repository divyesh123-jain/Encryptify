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

func encrypt(inputFile string, outputFile string, key []byte) error {
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return os.WriteFile(outputFile, ciphertext, 0644)
}

func decrypt(inputFile string, outputFile string, key []byte) error {
	ciphertext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(ciphertext) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return os.WriteFile(outputFile, ciphertext, 0644)
}

func main() {
	mode := flag.String("mode", "encrypt", "Mode: 'encrypt' or 'decrypt'")
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")
	key := flag.String("key", "", "Encryption/Decryption key (16, 24, or 32 bytes)")

	flag.Parse()

	if *mode != "encrypt" && *mode != "decrypt" {
		fmt.Println("Invalid mode. Please use 'encrypt' or 'decrypt'.")
		return
	}

	if *inputFile == "" || *outputFile == "" || *key == "" {
		fmt.Println("Please provide input file, output file, and key.")
		return
	}

	keyBytes := []byte(*key)

	var err error
	if *mode == "encrypt" {
		err = encrypt(*inputFile, *outputFile, keyBytes)
	} else {
		err = decrypt(*inputFile, *outputFile, keyBytes)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%s successful.\n", *mode)
}
