package main

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Define command-line flag
	typ := flag.String("t", "", "Type of key to generate: eddsa, aes, ecdsa, rsa, uuid")
	flag.Parse()

	switch *typ {
	case "uuid":
		fmt.Println("UUID:", UUID())
	case "eddsa":
		priv, pub, err := EDDSA()
		if err != nil {
			fmt.Println("Error generating EDDSA keys:", err)
			return
		}
		fmt.Println("Private Key (Base64):", priv)
		fmt.Println("Public Key  (Base64):", pub)
	case "aes":
		key, err := AES()
		if err != nil {
			fmt.Println("Error generating AES key:", err)
			return
		}
		fmt.Println("AES Key (Base64):", key)
	case "ecdsa":
		priv, pub, err := ECDSA()
		if err != nil {
			fmt.Println("Error generating ECDSA keys:", err)
			return
		}
		fmt.Println("Private Key (Base64):", priv)
		fmt.Println("Public Key  (Base64):", pub)
	case "rsa":
		priv, pub, err := RSA()
		if err != nil {
			fmt.Println("Error generating RSA keys:", err)
			return
		}
		fmt.Println("Private Key (Base64):", priv)
		fmt.Println("Public Key  (Base64):", pub)
	default:
		fmt.Println("Invalid or missing -t flag. Use one of: eddsa, aes, ecdsa, rsa, uuid")
	}
}

func UUID() string {
	return uuid.NewString()
}

func EDDSA() (string, string, error) {
	public, private, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", nil
	}
	b64Public := base64.StdEncoding.EncodeToString(public)
	b64Private := base64.StdEncoding.EncodeToString(private)
	return b64Private, b64Public, nil
}

func AES() (string, error) {
	buffer := make([]byte, 32)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	b64Key := base64.StdEncoding.EncodeToString(buffer)
	return b64Key, nil
}

func ECDSA() (string, string, error) {
	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}
	privateBytes, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		return "", "", err
	}
	publicBytes, err := x509.MarshalPKIXPublicKey(&private.PublicKey)
	if err != nil {
		return "", "", err
	}
	privateBase64 := base64.StdEncoding.EncodeToString(privateBytes)
	publicBase64 := base64.StdEncoding.EncodeToString(publicBytes)
	return privateBase64, publicBase64, nil
}

func RSA() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	publicBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	privateBase64 := base64.StdEncoding.EncodeToString(privateBytes)
	publicBase64 := base64.StdEncoding.EncodeToString(publicBytes)

	return privateBase64, publicBase64, nil
}
