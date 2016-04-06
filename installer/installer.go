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
	uri := GetResourceURI(depType, tag)
	if uri == "" {
		return fmt.Errorf("Unsupported platform")
	}

	target := GetBinPath()
	err := os.MkdirAll(target, 0777)
	if err != nil {
		return err
	}

	fmt.Println("downloading...", uri)
	downloadFile, err := download(uri, target)
	if err != nil {
		return err
	}

	fmt.Println("extracting...", downloadFile, target)
	extractorClient := extractor.New()
	err = extractorClient.Do(downloadFile, target)
	if err != nil {
		return err
	}

	err = ExtractBin(depType, target, tag)
	if err != nil {
		return err
	}

	return nil
}

func getFileName(source string) (string, error) {
	uri, err := url.Parse(source)
	if err != nil {
		return "", err
	}
	segments := strings.Split(uri.Path, "/")
	return segments[len(segments)-1], nil
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
