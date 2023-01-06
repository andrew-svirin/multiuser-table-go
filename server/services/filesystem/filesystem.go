package filesystem

// Filesystem - service helps to operate with files and directories.

import (
	"log"
	"os"
)

// ResourcesDir - dir with resources dir.
const ResourcesDir = "/resources"

// StaticDir - dir with static dir.
const StaticDir = "/static"

// ResolveResourcePath - resolve path to the file in `resources` dir.
func ResolveResourcePath(rel string) string {
	return resolveRootDir() + ResourcesDir + rel
}

// ResolveStaticPath - resolve path to the file in `static` dir.
func ResolveStaticPath(rel string) string {
	return resolveRootDir() + ResourcesDir + StaticDir + rel
}

// resolveRootDir - resolve path to the root dir.
func resolveRootDir() string {
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal("resolveRootDir: ", err)
	}

	return pwd
}
