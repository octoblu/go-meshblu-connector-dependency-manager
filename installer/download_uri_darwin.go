package installer

import (
	"os"
	"path"
	"strings"
)

// GetNodeURI defines the uri to download to
func GetNodeURI(tag string) string {
	return strings.Replace("https://nodejs.org/dist/:tag:/node-:tag:-darwin-x64.tar.gz", ":tag:", tag, -1)
}

// GetBinPath defines the target location
func GetBinPath() string {
	return path.Join(os.Getenv("HOME"), "Library", "Application Support", "MeshbluConnectors", "bin")
}

// ExtractBin allows you too extract the bin from the download
func ExtractBin(target, tag string) error {
	folderName := strings.Replace("node-:tag:-darwin-x64", ":tag:", tag, -1)
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
