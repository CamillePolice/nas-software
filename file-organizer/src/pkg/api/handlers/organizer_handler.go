// Package handlers provides HTTP request handlers for the file organizer API
package handlers

import (
	"bytes"
	"file-organizer/src/pkg/types"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// OrganizerHandler handles the building and distribution of platform-specific binaries
type OrganizerHandler struct {
	// binDir is the directory where built binaries are stored
	binDir string
}

// NewOrganizerHandler creates a new instance of OrganizerHandler
// It initializes the handler with the default bin directory
func NewOrganizerHandler() *OrganizerHandler {
	return &OrganizerHandler{
		binDir: "../bin",
	}
}

// buildBinary creates a binary for the specified operating system and architecture
//
// Parameters:
//   - platform: target operating system (e.g., "windows", "linux", "darwin") @See platform.go
//   - arch: target architecture (e.g., "amd64") @See platform.go
//
// Returns:
//   - string: path to the built binary
//   - error: any error that occurred during the build process
//
// The function creates a bin directory if it doesn't exist, sets up the appropriate
// build environment, and executes the go build command with the specified parameters.
func (h *OrganizerHandler) buildBinary(platform types.OS, arch types.Arch) (string, error) {
	binaryPath, err := h.getBinaryPath(platform, arch)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("go", "build", "-o", binaryPath, "./main.go")
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("GOOS=%s", platform),
		fmt.Sprintf("GOARCH=%s", arch),
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to build binary: %v\n%s", err, stderr.String())
	}

	return binaryPath, nil
}

// getBinaryPath generates the appropriate binary path based on the platform and architecture
//
// Parameters:
//   - platform: target operating system
//   - arch: target architecture
//
// Returns:
//   - string: complete path for the binary
//   - error: any error that occurred during directory creation
func (h *OrganizerHandler) getBinaryPath(platform types.OS, arch types.Arch) (string, error) {
	if err := os.MkdirAll(h.binDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create bin directory: %w", err)
	}

	binaryName := os.Getenv("BINARY_NAME")
	extension := ""
	if platform == types.Windows {
		extension = ".exe"
	}

	return filepath.Join(h.binDir, fmt.Sprintf("%s-%s-%s%s", binaryName, platform, arch, extension)), nil
}

// getBinaryResponse is a helper function that builds and prepares binary response
//
// Parameters:
//   - platform: target operating system
//   - arch: target architecture
//
// Returns:
//   - BinaryResponse: struct containing the binary data and filename
//   - error: any error that occurred during the build or read process
//
// This function handles the building of the binary and preparation of the response
// including reading the binary data and setting up the response structure
func (h *OrganizerHandler) getBinaryResponse(platform types.OS, arch types.Arch) (types.BinaryResponse, error) {
	binaryPath, err := h.buildBinary(platform, arch)
	if err != nil {
		return types.BinaryResponse{}, err
	}

	binary, err := os.ReadFile(binaryPath)
	if err != nil {
		return types.BinaryResponse{}, fmt.Errorf("failed to read binary: %w", err)
	}

	return types.BinaryResponse{
		Binary:   binary,
		Filename: filepath.Base(binaryPath),
	}, nil
}

// GetWindowsBinary handles requests for the Windows version of the file organizer
//
// This handler:
// 1. Builds a Windows-specific binary (amd64 architecture)
// 2. Returns the binary as a JSON response with the binary data and filename
// 3. Returns a JSON error response if the build fails
//
// Route: GET /binary/windows
func (h *OrganizerHandler) GetWindowsBinary(c echo.Context) error {
	resp, err := h.getBinaryResponse(types.Windows, types.AMD64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

// GetLinuxBinary handles requests for the Linux version of the file organizer
//
// This handler:
// 1. Builds a Linux-specific binary (amd64 architecture)
// 2. Returns the binary as a JSON response with the binary data and filename
// 3. Returns a JSON error response if the build fails
//
// Route: GET /binary/linux
func (h *OrganizerHandler) GetLinuxBinary(c echo.Context) error {
	resp, err := h.getBinaryResponse(types.Linux, types.AMD64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

// GetMacBinary handles requests for the macOS version of the file organizer
//
// This handler:
// 1. Builds a macOS-specific binary (amd64 architecture)
// 2. Returns the binary as a JSON response with the binary data and filename
// 3. Returns a JSON error response if the build fails
//
// Route: GET /binary/mac
func (h *OrganizerHandler) GetMacBinary(c echo.Context) error {
	resp, err := h.getBinaryResponse(types.Darwin, types.AMD64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
