package filesystem

// Filesystem - service helps to operate with files and directories.

import (
	"log"
	"os"
)

// StaticDir - dir with static files.
const StaticDir = "/static"

// ResolveStaticPath - resolve path to the file in `static` dir.
func ResolveStaticPath(rel string) string {
	return resolveRootDir() + StaticDir + rel
}

// resolveRootDir - resolve path to the root dir.
func resolveRootDir() string {
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal("resolveRootDir: ", err)
	}

	return pwd
}
