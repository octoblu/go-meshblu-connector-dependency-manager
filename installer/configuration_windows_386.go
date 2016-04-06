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
	if depType == NSSMType {
		return getNSSMURI(tag)
	}
	return ""
}

func getNSSMURI(tag string) string {
	return fmt.Sprintf("http://nssm.cc/release/nssm-%v.zip", tag)
}

func getNodeURI(tag string) string {
	return strings.Replace("https://nodejs.org/dist/:tag:/win-x86/node.exe", ":tag:", tag, -1)
}

// GetBinPath defines the target location
func GetBinPath() string {
	return path.Join(os.Getenv("LOCALAPPDATA"), "MeshbluConnectors", "bin")
}

// ExtractBin allows you too extract the bin from the download
func ExtractBin(depType, target, tag string) error {
	if depType == NodeType {
		return ExtractNode(target, tag)
	}
	if depType == NSSMType {
		return nil
	}
	return fmt.Errorf("Unsupported platform")
}
