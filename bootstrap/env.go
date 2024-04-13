package bootstrap

import "os"

const _nameAppDefault = "reto-financiera-compartamos"

func getApplicationName() string {
	appName := os.Getenv("K_SERVICE")
	if appName == "" {
		return _nameAppDefault
	}

	return appName
}
