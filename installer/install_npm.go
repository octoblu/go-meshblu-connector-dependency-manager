package installer

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/octoblu/go-meshblu-connector-installer/osruntime"
	"github.com/spf13/afero"
)

// InstallNPM installs the specified version of Node.JS
func InstallNPM(tag string) error {
	return InstallNodeWithoutDefaults(tag, "https://github.com", afero.NewOsFs(), osruntime.New())
}

// InstallNPMWithoutDefaults installs the specified version of NPM
// if you're running Windows. Otherwise, it does nothing.
func InstallNPMWithoutDefaults(tag, baseURLStr string, Fs afero.Fs, osRuntime osruntime.OSRuntime) error {
	if osRuntime.GOOS != Windows {
		return nil
	}

	packageURL, err := npmURL(baseURLStr, tag)
	if err != nil {
		return err
	}

	http.Get(packageURL.String())
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
