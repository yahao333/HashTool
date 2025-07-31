package api

// This package contains API related structures and documentation

// HashRequest represents the request structure for hash generation
// Example:
// {
//   "text": "your text here",
//   "algorithm": "sha1"
// }
type HashRequest struct {
	Text      string `json:"text" example:"hello world"`
	Algorithm string `json:"algorithm" example:"sha1"`
}

// HashResponse represents the response structure for hash generation
// Example:
// {
//   "hash": "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
//   "algorithm": "sha1"
// }
type HashResponse struct {
	Hash      string `json:"hash" example:"2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"`
	Algorithm string `json:"algorithm" example:"sha1"`
	Text      string `json:"text,omitempty" example:"hello world"`
}

// VerifyRequest represents the request structure for hash verification
// Example:
// {
//   "text": "hello world",
//   "hash": "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
//   "algorithm": "sha1"
// }
type VerifyRequest struct {
	Text      string `json:"text" example:"hello world"`
	Hash      string `json:"hash" example:"2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"`
	Algorithm string `json:"algorithm" example:"sha1"`
}

// VerifyResponse represents the response structure for hash verification
// Example:
// {
//   "valid": true,
//   "algorithm": "sha1"
// }
type VerifyResponse struct {
	Valid     bool   `json:"valid" example:"true"`
	Algorithm string `json:"algorithm" example:"sha1"`
}

// ErrorResponse represents error response structure
// Example:
// {
//   "error": "unsupported algorithm: md4"
// }
type ErrorResponse struct {
	Error string `json:"error" example:"unsupported algorithm: md4"`
}