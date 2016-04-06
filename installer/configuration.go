package installer

import (
	"os"
	"path"
	"strings"
)

// ExtractNode extracts the node dependencies
func ExtractNode(target, tag string) error {
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
