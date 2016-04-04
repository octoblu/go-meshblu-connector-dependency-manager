package installer

import "strings"

// GetNodeURI defines the uri to download to
func GetNodeURI(tag string) string {
	return strings.Replace("https://nodejs.org/dist/:tag:/node-:tag:-linux-x64.tar.gz", ":tag:", tag, -1)
}

// GetBinPath defines the target location
func GetBinPath() string {
	return ""
}

// ExtractBin allows you too extract the bin from the download
func ExtractBin(target, tag string) error {
	return nil
}
