package phplint

import (
	"fmt"
	"strings"

	phpversion "github.com/shyim/go-phplint/internal/version"
)

// Version identifies a PHP minor-language profile.
type Version uint8

const (
	PHP72 Version = iota + 1
	PHP73
	PHP74
	PHP80
	PHP81
	PHP82
	PHP83
	PHP84
	PHP85
	// PHP86 is the preview profile for the current PHP 8.6 prerelease.
	PHP86
)

var supportedVersions = [...]Version{
	PHP72,
	PHP73,
	PHP74,
	PHP80,
	PHP81,
	PHP82,
	PHP83,
	PHP84,
	PHP85,
	PHP86,
}

// ParseVersion parses one of the supported PHP minor versions.
//
// Patch versions are deliberately rejected: PHP syntax profiles are maintained
// at minor-version granularity.
func ParseVersion(value string) (Version, error) {
	switch strings.TrimSpace(value) {
	case "7.2":
		return PHP72, nil
	case "7.3":
		return PHP73, nil
	case "7.4":
		return PHP74, nil
	case "8.0":
		return PHP80, nil
	case "8.1":
		return PHP81, nil
	case "8.2":
		return PHP82, nil
	case "8.3":
		return PHP83, nil
	case "8.4":
		return PHP84, nil
	case "8.5":
		return PHP85, nil
	case "8.6":
		return PHP86, nil
	default:
		return 0, fmt.Errorf(
			"unsupported PHP version %q (supported: 7.2, 7.3, 7.4, 8.0, 8.1, 8.2, 8.3, 8.4, 8.5, 8.6 preview)",
			value,
		)
	}
}

// SupportedVersions returns all supported PHP profiles in ascending order.
func SupportedVersions() []Version {
	versions := make([]Version, len(supportedVersions))
	copy(versions, supportedVersions[:])
	return versions
}

// String returns the canonical major.minor form of the version.
func (v Version) String() string {
	switch v {
	case PHP72:
		return "7.2"
	case PHP73:
		return "7.3"
	case PHP74:
		return "7.4"
	case PHP80:
		return "8.0"
	case PHP81:
		return "8.1"
	case PHP82:
		return "8.2"
	case PHP83:
		return "8.3"
	case PHP84:
		return "8.4"
	case PHP85:
		return "8.5"
	case PHP86:
		return "8.6"
	default:
		return "unknown"
	}
}

func (v Version) valid() bool {
	return v >= PHP72 && v <= PHP86
}

func (v Version) internal() *phpversion.Version {
	if v <= PHP74 {
		return &phpversion.Version{Major: 7, Minor: uint64(v-PHP72) + 2}
	}

	return &phpversion.Version{Major: 8, Minor: uint64(v - PHP80)}
}
