package installer

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/octoblu/go-meshblu-connector-installer/osruntime"
	"github.com/spf13/afero"
)

// InstallNode installs the specified version of Node.JS
func InstallNode(tag, binPath string) error {
	return InstallNodeWithoutDefaults(tag, binPath, "https://nodejs.org", afero.NewOsFs(), osruntime.New())
}

// InstallNodeWithoutDefaults installs the specified version of Node.JS
func InstallNodeWithoutDefaults(tag, binPath, baseURLStr string, Fs afero.Fs, osRuntime osruntime.OSRuntime) error {
	packageURL, err := nodeURL(baseURLStr, tag, osRuntime)
	if err != nil {
		return err
	}

	http.Get(packageURL.String())
	filePath := filepath.Join(binPath, "node")
	afero.WriteFile(Fs, filePath, []byte(""), 0755)
	return nil
}

func nodeURL(baseURLStr, tag string, osRuntime osruntime.OSRuntime) (*url.URL, error) {
	nodeURL, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, err
	}

	fileName, err := nodeFileName(tag, osRuntime)
	if err != nil {
		return nil, err
	}

	filePath, err := nodeFilePath(tag, osRuntime)
	if err != nil {
		return nil, err
	}

	nodeURL.Path = fmt.Sprintf("%v/%v", filePath, fileName)
	return nodeURL, nil
}

func nodeFileName(tag string, osRuntime osruntime.OSRuntime) (string, error) {
	if osRuntime.GOOS == "windows" {
		return "node.exe", nil
	}

	nodeArch, ok := ArchMap[osRuntime.GOARCH]
	if !ok {
		return "", fmt.Errorf("Unsupported architecture: %v", osRuntime.GOARCH)
	}

	return fmt.Sprintf("node-%v-%v-%v.tar.gz", tag, osRuntime.GOOS, nodeArch), nil
}

func nodeFilePath(tag string, osRuntime osruntime.OSRuntime) (string, error) {
	nodeArch, ok := ArchMap[osRuntime.GOARCH]
	if !ok {
		return "", fmt.Errorf("Unsupported architecture: %v", osRuntime.GOARCH)
	}

	if osRuntime.GOOS == "windows" {
		return fmt.Sprintf("/dist/%v/win-%v", tag, nodeArch), nil
	}

	return fmt.Sprintf("/dist/%v", tag), nil
}

// // Install is a convenience method for constructing an installer client
// // and calling client.Do
// func Install(depType, tag string) error {
// 	// client := New()
// 	// return client.Do(depType, tag)
// 	return nil
// }
