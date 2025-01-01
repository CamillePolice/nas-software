package types

type OS string

const (
	Windows OS = "windows"
	Linux   OS = "linux"
	Darwin  OS = "darwin"
)

func (o OS) String() string {
	return string(o)
}

func (o OS) IsValid() bool {
	switch o {
	case Windows, Linux, Darwin:
		return true
	default:
		return false
	}
}

// Arch represents supported CPU architectures
type Arch string

const (
	AMD64 Arch = "amd64"
	ARM64 Arch = "arm64"
	X86   Arch = "386"
)

func (a Arch) String() string {
	return string(a)
}

func (a Arch) IsValid() bool {
	switch a {
	case AMD64, ARM64, X86:
		return true
	default:
		return false
	}
}
