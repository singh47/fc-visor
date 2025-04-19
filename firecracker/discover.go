package firecracker

import (
	"os"
	"path/filepath"
)

// DiscoverSockets scans for Firecracker sockets in known locations
func DiscoverSockets() []string {
	sockets := []string{}
	possibleDirs := []string{"/run/firecracker", "/tmp"}
	for _, dir := range possibleDirs {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && filepath.Ext(path) == ".sock" {
				sockets = append(sockets, path)
			}
			return nil
		})
	}
	return sockets
}
