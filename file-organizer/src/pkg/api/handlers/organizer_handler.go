// Package handlers provides HTTP request handlers for the file organizer API
package handlers

import (
	"archive/zip"
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

// addReadmeToZip adds the README file to the provided zip writer
//
// Parameters:
//   - w: zip writer to add the README to
//
// Returns:
//   - error: any error that occurred during the process
//
// This function reads the README file from the assets directory and adds it to the zip archive
func (h *OrganizerHandler) addReadmeToZip(w *zip.Writer) error {
	readmeContent, err := os.ReadFile("../../assets/ReadMe.md")
	if err != nil {
		return fmt.Errorf("failed to read README file: %w", err)
	}

	readmeWriter, err := w.Create("README.md")
	if err != nil {
		return fmt.Errorf("failed to create README in archive: %w", err)
	}

	if _, err := readmeWriter.Write(readmeContent); err != nil {
		return fmt.Errorf("failed to write README to archive: %w", err)
	}

	return nil
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
	fmt.Println("Starting zip creation...")

	// Create a buffer to write our archive to
	buf := new(bytes.Buffer)

	// Create a new zip archive
	w := zip.NewWriter(buf)

	// First, get and add the binary
	binaryPath, err := h.buildBinary(platform, arch)
	if err != nil {
		return types.BinaryResponse{}, err
	}
	fmt.Printf("Binary path: %s\n", binaryPath)

	binary, err := os.ReadFile(binaryPath)
	if err != nil {
		return types.BinaryResponse{}, fmt.Errorf("failed to read binary: %w", err)
	}
	fmt.Printf("Binary size: %d bytes\n", len(binary))

	// Add binary to zip with its base filename
	binaryName := filepath.Base(binaryPath)
	fmt.Printf("Adding binary to zip as: %s\n", binaryName)

	binaryWriter, err := w.Create(binaryName)
	if err != nil {
		return types.BinaryResponse{}, fmt.Errorf("failed to create binary in archive: %w", err)
	}

	written, err := binaryWriter.Write(binary)
	if err != nil {
		return types.BinaryResponse{}, fmt.Errorf("failed to write binary to archive: %w", err)
	}
	fmt.Printf("Wrote %d bytes of binary to zip\n", written)

	// Add README to zip
	fmt.Println("Adding README to zip...")
	if err := h.addReadmeToZip(w); err != nil {
		return types.BinaryResponse{}, fmt.Errorf("failed to add README to archive: %w", err)
	}

	// Important: Close the zip writer before reading the buffer
	fmt.Println("Closing zip writer...")
	err = w.Close()
	if err != nil {
		return types.BinaryResponse{}, fmt.Errorf("failed to close zip writer: %w", err)
	}

	// Get the bytes after closing the writer
	zipBytes := buf.Bytes()
	fmt.Printf("Final zip size: %d bytes\n", len(zipBytes))

	// Verify the zip is not empty
	if len(zipBytes) == 0 {
		return types.BinaryResponse{}, fmt.Errorf("generated zip archive is empty")
	}

	archiveName := fmt.Sprintf("file-organizer-%s-%s.zip", platform, arch)
	fmt.Printf("Returning zip with name: %s\n", archiveName)

	return types.BinaryResponse{
		Binary:      zipBytes,
		ArchiveName: archiveName,
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
