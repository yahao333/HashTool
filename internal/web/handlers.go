package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/yahao333/hashtool/internal/hash"
)

// HashRequest represents the request structure for hash generation
type HashRequest struct {
	Text      string `json:"text"`
	Algorithm string `json:"algorithm"`
}

// HashResponse represents the response structure for hash generation
type HashResponse struct {
	Hash      string `json:"hash"`
	Algorithm string `json:"algorithm"`
	Text      string `json:"text,omitempty"`
}

// VerifyRequest represents the request structure for hash verification
type VerifyRequest struct {
	Text      string `json:"text"`
	Hash      string `json:"hash"`
	Algorithm string `json:"algorithm"`
}

// VerifyResponse represents the response structure for hash verification
type VerifyResponse struct {
	Valid     bool   `json:"valid"`
	Algorithm string `json:"algorithm"`
}

// ErrorResponse represents error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// GenerateHashHandler handles hash generation requests
func GenerateHashHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req HashRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid JSON format"})
		return
	}

	if req.Text == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Text is required"})
		return
	}

	if req.Algorithm == "" {
		req.Algorithm = "sha1" // default algorithm
	}

	hashResult, err := hash.GenerateHash(req.Text, strings.ToLower(req.Algorithm))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	response := HashResponse{
		Hash:      hashResult,
		Algorithm: strings.ToLower(req.Algorithm),
	}

	json.NewEncoder(w).Encode(response)
}

// VerifyHashHandler handles hash verification requests
func VerifyHashHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid JSON format"})
		return
	}

	if req.Text == "" || req.Hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Text and hash are required"})
		return
	}

	if req.Algorithm == "" {
		req.Algorithm = "sha1" // default algorithm
	}

	isValid, err := hash.VerifyHash(req.Text, req.Hash, strings.ToLower(req.Algorithm))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	response := VerifyResponse{
		Valid:     isValid,
		Algorithm: strings.ToLower(req.Algorithm),
	}

	json.NewEncoder(w).Encode(response)
}

// SetupRoutes configures and returns the HTTP router
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/hash", GenerateHashHandler).Methods("POST")
	api.HandleFunc("/verify", VerifyHashHandler).Methods("POST")

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods("GET")

	return r
}