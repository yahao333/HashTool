package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

// GenerateHash generates hash for given text using specified algorithm
func GenerateHash(text, algorithm string) (string, error) {
	switch strings.ToLower(algorithm) {
	case "md5":
		hasher := md5.New()
		hasher.Write([]byte(text))
		return hex.EncodeToString(hasher.Sum(nil)), nil
	case "sha1":
		hasher := sha1.New()
		hasher.Write([]byte(text))
		return hex.EncodeToString(hasher.Sum(nil)), nil
	case "sha256":
		hasher := sha256.New()
		hasher.Write([]byte(text))
		return hex.EncodeToString(hasher.Sum(nil)), nil
	case "sha512":
		hasher := sha512.New()
		hasher.Write([]byte(text))
		return hex.EncodeToString(hasher.Sum(nil)), nil
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
	}
}

// VerifyHash verifies if the given text matches the provided hash
func VerifyHash(text, hash, algorithm string) (bool, error) {
	generatedHash, err := GenerateHash(text, algorithm)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(generatedHash, hash), nil
}

// GetSupportedAlgorithms returns list of supported hash algorithms
func GetSupportedAlgorithms() []string {
	return []string{"md5", "sha1", "sha256", "sha512"}
}