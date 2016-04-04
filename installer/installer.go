package installer

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/octoblu/go-meshblu-connector-installer/extractor"
)

// Installer interfaces with the a remote download server
type Installer interface {
	Do(depType, tag string) error
}

// Client interfaces with the a remote download server
type Client struct {
}

// New constructs a new Installer instance
func New() *Client {
	return &Client{}
}

// Do download and install
func (client *Client) Do(depType, tag string) error {
	uri := ""
	if depType == "node" {
		uri = GetNodeURI(tag)
	} else {
		return nil
	}

	target := GetBinPath()
	err := os.MkdirAll(target, 0755)
	if err != nil {
		return err
	}

	fmt.Println("downloading...", uri)
	downloadFile, err := download(uri, target)
	if err != nil {
		return err
	}

	if isTarGZ(uri) {
		fmt.Println("extracting...", downloadFile, target)
		extractorClient := extractor.New()
		err = extractorClient.Do(downloadFile, target)
		if err != nil {
			return err
		}
	}

	err = ExtractBin(target, tag)
	if err != nil {
		return err
	}

	return nil
}

func getFileName(uriRaw string) (string, error) {
	uri, err := url.Parse(uriRaw)
	if err != nil {
		return "", err
	}
	segments := strings.Split(uri.Path, "/")
	return segments[len(segments)-1], nil
}

func isTarGZ(uriRaw string) bool {
	fileName, err := getFileName(uriRaw)
	if err != nil {
		return false
	}
	return strings.Index(fileName, ".tar.gz") > -1
}

func download(uri, target string) (string, error) {
	fileName, err := getFileName(uri)
	if err != nil {
		return "", err
	}
	downloadFile := path.Join(target, fileName)
	outputStream, err := os.Create(downloadFile)

	if err != nil {
		fmt.Println("Error on os.Create", err.Error())
		return "", err
	}

	defer outputStream.Close()

	response, err := http.Get(uri)

	if err != nil {
		fmt.Println("Error on http.Get", err.Error())
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("Download invalid status code: %v", response.StatusCode)
	}

	_, err = io.Copy(outputStream, response.Body)

	if err != nil {
		fmt.Println("Error on io.Copy", err.Error())
		return "", err
	}
	return downloadFile, nil
}
