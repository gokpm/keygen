# Crypto Key Generator

A simple Go command-line tool for generating various types of cryptographic keys and UUIDs.

## Usage

```bash
go build -o keygen
./keygen -t <type>
```

## Supported Key Types

- `uuid` - Generate a UUID
- `eddsa` - Generate Ed25519 key pair
- `aes` - Generate AES-256 key
- `ecdsa` - Generate ECDSA key pair (P-256 curve)
- `rsa` - Generate RSA key pair (2048-bit)

## Examples

```bash
# Generate UUID
./keygen -t uuid

# Generate Ed25519 keys
./keygen -t eddsa

# Generate AES key
./keygen -t aes

# Generate ECDSA keys
./keygen -t ecdsa

# Generate RSA keys
./keygen -t rsa
```

## Output

All keys are encoded in Base64 format for easy copying and storage.

## Dependencies

- `github.com/google/uuid`

Install with: `go mod tidy`