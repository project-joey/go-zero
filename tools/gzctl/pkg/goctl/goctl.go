package goctl

import (
	"path/filepath"
	"runtime"

	"github.com/project-joey/go-zero/tools/gzctl/pkg/golang"
	"github.com/project-joey/go-zero/tools/gzctl/util/console"
	"github.com/project-joey/go-zero/tools/gzctl/util/pathx"
	"github.com/project-joey/go-zero/tools/gzctl/vars"
)

func Install(cacheDir, name string, installFn func(dest string) (string, error)) (string, error) {
	goBin := golang.GoBin()
	cacheFile := filepath.Join(cacheDir, name)
	binFile := filepath.Join(goBin, name)

	goos := runtime.GOOS
	if goos == vars.OsWindows {
		cacheFile = cacheFile + ".exe"
		binFile = binFile + ".exe"
	}
	// read cache.
	err := pathx.Copy(cacheFile, binFile)
	if err == nil {
		console.Info("%q installed from cache", name)
		return binFile, nil
	}

	binFile, err = installFn(binFile)
	if err != nil {
		return "", err
	}

	// write cache.
	err = pathx.Copy(binFile, cacheFile)
	if err != nil {
		console.Warning("write cache error: %+v", err)
	}
	return binFile, nil
}
