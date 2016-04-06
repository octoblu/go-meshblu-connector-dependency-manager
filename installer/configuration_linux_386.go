package installer

import (
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

// getNodeURI defines the uri to download to
func getNodeURI(tag string) string {
	return strings.Replace("https://nodejs.org/dist/:tag:/node-:tag:-linux-x64.tar.gz", ":tag:", tag, -1)
}

// GetBinPath defines the target location
func GetBinPath() string {
	return path.Join(os.Getenv("HOME"), ".octoblu", "MeshbluConnectors")
}

// ExtractBin allows you too extract the bin from the download
func ExtractBin(depType, target, tag string) error {
	return nil
}

// ExtractNode extracts the node dependencies
func ExtractNode(target, tag string) error {
	folderName := strings.Replace("node-:tag:-linux-x64", ":tag:", tag, -1)
	nodePath := path.Join(target, folderName, "bin", "node")
	nodeSymPath := path.Join(target, "node")
	os.Remove(nodeSymPath)
	err := os.Symlink(nodePath, nodeSymPath)
	if err != nil {
		return err
	}

	npmPath := path.Join(target, folderName, "bin", "npm")
	npmSymPath := path.Join(target, "npm")
	os.Remove(npmSymPath)
	err = os.Symlink(npmPath, npmSymPath)
	if err != nil {
		return err
	}
	return nil
}
