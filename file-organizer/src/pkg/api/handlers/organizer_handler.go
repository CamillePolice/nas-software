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

	// Execute build command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to build binary: %v\n%s", err, stderr.String())
	}

	return binaryPath, nil
}

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

// GetWindowsBinary handles requests for the Windows version of the file organizer
//
// This handler:
// 1. Builds a Windows-specific binary (amd64 architecture)
// 2. Returns the binary as a downloadable attachment
// 3. Returns a JSON error response if the build fails
//
// Route: GET /binary/windows
func (h *OrganizerHandler) GetWindowsBinary(c echo.Context) error {
	binaryPath, err := h.buildBinary(types.Windows, types.AMD64)

	fmt.Println("BinaryPath", binaryPath)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Error: err.Error(),
		})
	}

	// Read the binary file
	binary, err := os.ReadFile(binaryPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Error: "Failed to read binary",
		})
	}

	// Send both binary and filename
	return c.JSON(http.StatusOK, types.BinaryResponse{
		Binary:   binary,
		Filename: filepath.Base(binaryPath),
	})
}

// GetLinuxBinary handles requests for the Linux version of the file organizer
//
// This handler:
// 1. Builds a Linux-specific binary (amd64 architecture)
// 2. Returns the binary as a downloadable attachment
// 3. Returns a JSON error response if the build fails
//
// Route: GET /binary/linux
func (h *OrganizerHandler) GetLinuxBinary(c echo.Context) error {
	binaryPath, err := h.buildBinary(types.Linux, types.AMD64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Error: err.Error(),
		})
	}

	return c.Attachment(binaryPath, filepath.Base(binaryPath))
}

// GetMacBinary handles requests for the macOS version of the file organizer
//
// This handler:
// 1. Builds a macOS-specific binary (amd64 architecture)
// 2. Returns the binary as a downloadable attachment
// 3. Returns a JSON error response if the build fails
//
// Route: GET /binary/mac
func (h *OrganizerHandler) GetMacBinary(c echo.Context) error {
	binaryPath, err := h.buildBinary(types.Darwin, types.AMD64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Error: err.Error(),
		})
	}

	return c.Attachment(binaryPath, filepath.Base(binaryPath))
}
