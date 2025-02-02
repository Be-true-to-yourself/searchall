//go:build linux

package browser

import (
	"searchall3.5/tuozhan/liulanqi/item"
)

var (
	chromiumList = map[string]struct {
		name        string
		storage     string
		profilePath string
		items       []item.Item
	}{
		"chrome": {
			name:        chromeName,
			storage:     chromeStorageName,
			profilePath: chromeProfilePath,
			items:       item.DefaultChromium,
		},
		"edge": {
			name:        edgeName,
			storage:     edgeStorageName,
			profilePath: edgeProfilePath,
			items:       item.DefaultChromium,
		},
		"chromium": {
			name:        chromiumName,
			storage:     chromiumStorageName,
			profilePath: chromiumProfilePath,
			items:       item.DefaultChromium,
		},
		"chrome-beta": {
			name:        chromeBetaName,
			storage:     chromeBetaStorageName,
			profilePath: chromeBetaProfilePath,
			items:       item.DefaultChromium,
		},
	}
	firefoxList = map[string]struct {
		name        string
		storage     string
		profilePath string
		items       []item.Item
	}{
		"firefox": {
			name:        firefoxName,
			profilePath: firefoxProfilePath,
			items:       item.DefaultFirefox,
		},
	}
)

var (
	firefoxProfilePath    = homeDir + "/.mozilla/firefox/"
	chromeProfilePath     = homeDir + "/.config/google-chrome/Default/"
	chromiumProfilePath   = homeDir + "/.config/chromium/Default/"
	edgeProfilePath       = homeDir + "/.config/microsoft-edge/Default/"
	chromeBetaProfilePath = homeDir + "/.config/google-chrome-beta/Default/"
)

const (
	chromeStorageName     = "Chrome Safe Storage"
	chromiumStorageName   = "Chromium Safe Storage"
	edgeStorageName       = "Chromium Safe Storage"
	chromeBetaStorageName = "Chrome Safe Storage"
)
