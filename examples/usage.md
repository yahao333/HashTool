# HashTool Usage Examples

## CLI Usage

### Basic Hash Generation
```bash
# Generate SHA1 hash (default)
hashtool hash "hello world"
# Output: 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed

# Generate MD5 hash
hashtool hash -a md5 "hello world"
# Output: 5d41402abc4b2a76b9719d911017c592

# Generate SHA256 hash
hashtool hash -a sha256 "hello world"
# Output: b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9

# Generate SHA512 hash
hashtool hash -a sha512 "hello world"
# Output: 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
```

### Hash a Password (like the provided example)
```bash
hashtool hash -a sha1 "13691902544"
# Output: 356a192b7913b04c54574d18c28d46e6395428ab
```

## Web Service Usage

### Start the Server
```bash
# Start server on default port 8080
hashtool server

# Start server on custom port
hashtool server -p 3000
```

### API Examples

#### Generate Hash
```bash
# Using curl to generate SHA1 hash
curl -X POST http://localhost:8080/api/hash \
  -H "Content-Type: application/json" \
  -d '{
    "text": "hello world",
    "algorithm": "sha1"
  }'

# Response:
# {
#   "hash": "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
#   "algorithm": "sha1"
# }
```

```bash
# Generate MD5 hash
curl -X POST http://localhost:8080/api/hash \
  -H "Content-Type: application/json" \
  -d '{
    "text": "13691902544",
    "algorithm": "md5"
  }'

# Response:
# {
#   "hash": "c4ca4238a0b923820dcc509a6f75849b",
#   "algorithm": "md5"
# }
```

#### Verify Hash
```bash
# Verify a hash
curl -X POST http://localhost:8080/api/verify \
  -H "Content-Type: application/json" \
  -d '{
    "text": "hello world",
    "hash": "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
    "algorithm": "sha1"
  }'

# Response:
# {
#   "valid": true,
#   "algorithm": "sha1"
# }
```

```bash
# Verify with wrong hash
curl -X POST http://localhost:8080/api/verify \
  -H "Content-Type: application/json" \
  -d '{
    "text": "hello world",
    "hash": "wronghash",
    "algorithm": "sha1"
  }'

# Response:
# {
#   "valid": false,
#   "algorithm": "sha1"
# }
```

#### Health Check
```bash
# Check server health
curl http://localhost:8080/health

# Response:
# {
#   "status": "ok"
# }
```

## Supported Algorithms

- `md5` - MD5 hash algorithm
- `sha1` - SHA-1 hash algorithm (default)
- `sha256` - SHA-256 hash algorithm
- `sha512` - SHA-512 hash algorithm

## Error Handling

### CLI Errors
```bash
# Unsupported algorithm
hashtool hash -a md4 "test"
# Output: Error: unsupported algorithm: md4
```

### API Errors
```bash
# Missing required field
curl -X POST http://localhost:8080/api/hash \
  -H "Content-Type: application/json" \
  -d '{
    "algorithm": "sha1"
  }'

# Response:
# {
#   "error": "Text is required"
# }
```

```bash
# Unsupported algorithm
curl -X POST http://localhost:8080/api/hash \
  -H "Content-Type: application/json" \
  -d '{
    "text": "test",
    "algorithm": "md4"
  }'

# Response:
# {
#   "error": "unsupported algorithm: md4"
# }
```