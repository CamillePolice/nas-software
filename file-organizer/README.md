# File Organizer API Documentation

## Overview
The File Organizer API provides endpoints to build and download platform-specific binaries of the file organizer tool. The API supports Windows, Linux, and macOS platforms.

## Base URL
```
http://localhost:8080
```

## Endpoints

### Health Check
Get the API health status.

```
GET /health
```

**Response**
```json
{
    "status": "healthy"
}
```

### Binary Downloads

#### Get Windows Binary
Build and download the Windows version of the file organizer.

```
GET /binary/windows
```

**Response**
- Content-Type: application/octet-stream
- Content-Disposition: attachment; filename=file-organizer-windows-amd64.exe

**Error Response**
```json
{
    "error": "error message"
}
```

#### Get Linux Binary
Build and download the Linux version of the file organizer.

```
GET /binary/linux
```

**Response**
- Content-Type: application/octet-stream
- Content-Disposition: attachment; filename=file-organizer-linux-amd64

**Error Response**
```json
{
    "error": "error message"
}
```

#### Get macOS Binary
Build and download the macOS version of the file organizer.

```
GET /binary/mac
```

**Response**
- Content-Type: application/octet-stream
- Content-Disposition: attachment; filename=file-organizer-darwin-amd64

**Error Response**
```json
{
    "error": "error message"
}
```

## Response Codes
- **200 OK**: Request successful
- **500 Internal Server Error**: Build failed or internal error occurred

## Example Usage

### Using cURL

Download Windows binary:
```bash
curl -OJ http://localhost:8080/binary/windows
```

Download Linux binary:
```bash
curl -OJ http://localhost:8080/binary/linux
```

Download macOS binary:
```bash
curl -OJ http://localhost:8080/binary/mac
```

### Using wget

```bash
wget http://localhost:8080/binary/windows
wget http://localhost:8080/binary/linux
wget http://localhost:8080/binary/mac
```

## Implementation Details

### Binary Build Process
- Binaries are built on-demand when requested
- Built using `go build` with appropriate GOOS and GOARCH settings
- Stored in a `bin` directory with platform-specific naming
- Windows binaries include `.exe` extension

### File Structure
```
bin/
├── file-organizer-windows-amd64.exe
├── file-organizer-linux-amd64
└── file-organizer-darwin-amd64
```

## Error Handling
All endpoints return JSON error responses in the following format:
```json
{
    "error": "Detailed error message"
}
```

Common errors:
- Failed to create bin directory
- Build command failed
- Invalid platform specified

## Notes
- All binaries are built for amd64 architecture
- Binaries are built fresh for each request
- Build process may take a few seconds depending on the system