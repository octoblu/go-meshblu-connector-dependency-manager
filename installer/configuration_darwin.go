package installer

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// GetResourceURI defines the uri to download
func GetResourceURI(depType, tag string) string {
	if depType == NodeType {
		return getNodeURI(tag)
	}
	return ""
}

func getNodeURI(tag string) string {
	return strings.Replace("https://nodejs.org/dist/:tag:/node-:tag:-darwin-x64.tar.gz", ":tag:", tag, -1)
}

// GetBinPath defines the target location
func GetBinPath() string {
	return path.Join(os.Getenv("HOME"), "Library", "Application Support", "MeshbluConnectors", "bin")
}

// ExtractBin allows you too extract the bin from the download
func ExtractBin(depType, target, tag string) error {
	if depType == NodeType {
		return ExtractNode(target, tag)
	}
	return fmt.Errorf("Unsupported platform")
}
