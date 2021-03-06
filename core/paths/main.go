package paths

import (
	"github.com/hashload/boss/env"
	"github.com/hashload/boss/models"
	"github.com/masterminds/glide/msg"
	"os"
	"path/filepath"
)

func EnsureCacheDir(dep models.Dependency) {
	cacheDir := filepath.Join(env.GetCacheDir(), dep.GetHashName())

	fi, err := os.Stat(cacheDir)
	if err != nil {
		msg.Debug("Creating %s", cacheDir)
		if err := os.MkdirAll(cacheDir, os.ModeDir|0755); err != nil {
			msg.Die("Could not create %s: %s", cacheDir, err)
		}
	} else if !fi.IsDir() {
		msg.Die(".cache is not a directory")
	}
}

func EnsureModulesDir() {
	cacheDir := env.GetModulesDir()
	fi, err := os.Stat(cacheDir)
	if os.IsNotExist(err) {
		msg.Debug("Creating %s", cacheDir)
		if err := os.MkdirAll(cacheDir, os.ModeDir|0755); err != nil {
			msg.Die("Could not create %s: %s", cacheDir, err)
		}
	} else if !fi.IsDir() {
		msg.Die("modules is not a directory")
	} else {
		os.RemoveAll(cacheDir)
		EnsureModulesDir()
	}
}
