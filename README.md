**HashTool** / **哈希工具**

[English](README.md) | [中文](README.cn.md)

**HashTool - A Comprehensive Hash Utility**

![HashTool Screenshot](Screenshot.png)

HashTool is a versatile hash utility that provides both command-line interface (CLI) and web service capabilities. This tool supports multiple hash algorithms and offers an easy-to-use interface for generating and verifying hash values.

**Features:**
- 🔧 **Dual Interface**: Command-line (CLI) and Web Service
- 🌐 **Web API**: RESTful web service for integration
- 📋 **Easy to Use**: Simple commands and intuitive web interface

**Use Cases:**
- File integrity verification
- Password hashing
- Data fingerprinting
- Security applications

## English Description
A comprehensive hash utility with CLI and web service capabilities...

## Installation
```bash
go install github.com/yahao333/hashtool@latest
```

## Usage
### CLI Mode
```bash
hashtool --algorithm sha1 "your text"
```

### Web Service Mode
```bash
hashtool server --port 8080
```

## API Endpoints
- `POST /api/hash` - Generate hash
- `POST /api/verify` - Verify hash

## License
MIT License
