package manifest

import (
	"fmt"
	"path/filepath"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
)

// Path returns the manifest file path
func Path(browser, libpath string, globalInstall bool) (string, error) {
	location, err := GetLocation(browser, libpath, globalInstall)
	if err != nil {
		return "", err
	}

	expanded, err := homedir.Expand(location)
	if err != nil {
		return "", err
	}

	return filepath.Join(expanded, Name+".json"), nil
}

// GetLocation returns only the manifest path
func GetLocation(browser, libpath string, globalInstall bool) (string, error) {
	if globalInstall {
		return getGlobalLocation(browser, libpath)
	}

	pm, found := manifestPath[runtime.GOOS]
	if !found {
		return "", fmt.Errorf("platform %s is currently not supported", runtime.GOOS)
	}
	path, found := pm[browser]
	if !found {
		return "", fmt.Errorf("browser %s on %s is currently not supported", browser, runtime.GOOS)
	}
	return path, nil
}

func getGlobalLocation(browser, libpath string) (string, error) {
	pm, found := globalManifestPath[runtime.GOOS]
	if !found {
		return "", fmt.Errorf("platform %s is currently not supported", runtime.GOOS)
	}
	path, found := pm[browser]
	if !found {
		return "", fmt.Errorf("browser %s on %s is currently not supported", browser, runtime.GOOS)
	}
	if runtime.GOOS != "windows" && browser == "firefox" {
		path = libpath + "/" + path
	}
	return path, nil
}
