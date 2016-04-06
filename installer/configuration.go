package installer

import (
	"io"
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

// CopyFile copies a file
func CopyFile(source, target string) error {
	os.Remove(target)
	fileRead, err := os.Open(source)
	if err != nil {
		return err
	}
	defer fileRead.Close()

	fileWrite, err := os.Create(target)
	if err != nil {
		return err
	}
	defer fileWrite.Close()

	_, err = io.Copy(fileWrite, fileRead)
	if err != nil {
		return err
	}
	return nil
}
