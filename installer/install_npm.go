package installer

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/octoblu/go-meshblu-connector-installer/osruntime"
	"github.com/octoblu/unzipit"
)

// InstallNPM installs the specified version of Node.JS
func InstallNPM(tag, binPath string) error {
	return InstallNPMWithoutDefaults(tag, binPath, "https://github.com", osruntime.New())
}

// InstallNPMWithoutDefaults installs the specified version of NPM
// if you're running Windows. Otherwise, it does nothing.
func InstallNPMWithoutDefaults(tag, binPath, baseURLStr string, osRuntime osruntime.OSRuntime) error {
	if osRuntime.GOOS != Windows {
		return nil
	}

	packageURL, err := npmURL(baseURLStr, tag)
	if err != nil {
		return err
	}

	response, err := http.Get(packageURL.String())
	if err != nil {
		return err
	}

	return installNPMOnTheFS(tag, binPath, response.Body)
}

func installNPMOnTheFS(tag, binPath string, compressedReader io.Reader) error {
	archivePath, err := unzipit.UnpackStream(compressedReader, binPath)
	if err != nil {
		return err
	}
	version := strings.Replace(tag, "v", "", -1)
	archiveBinPath := filepath.Join(archivePath, fmt.Sprintf("npm-%v/bin", version))

	npmSrcFilePath := filepath.Join(archiveBinPath, "npm")
	npmDestFilePath := filepath.Join(binPath, "npm")
	err = os.Rename(npmSrcFilePath, npmDestFilePath)
	if err != nil {
		return err
	}

	npmCmdSrcFilePath := filepath.Join(archiveBinPath, "npm.cmd")
	npmCmdDestFilePath := filepath.Join(binPath, "npm.cmd")
	err = os.Rename(npmCmdSrcFilePath, npmCmdDestFilePath)
	if err != nil {
		return err
	}

	return nil
}

func npmURL(baseURLStr, tag string) (*url.URL, error) {
	npmURL, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, err
	}

	npmURL.Path = fmt.Sprintf("/npm/npm/archive/%v.zip", tag)
	return npmURL, nil
}
