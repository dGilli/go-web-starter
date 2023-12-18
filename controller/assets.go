package controller

import (
	"embed"
	"io/fs"
	"regexp"
	"strings"
)

func AssetsRewriteRules(assets embed.FS) map[string]string {
    rewriteRules := map[string]string{
        "/css/styles.css": "/tmp/css/styles.css", // Serve the compiled tailwind css
    }

    err := fs.WalkDir(assets, "assets", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if match, _ := regexp.MatchString("^[^/]+/[^/]+$", path); match {
			rule := strings.TrimPrefix(path, "assets/")
			rewriteRules["/"+rule+"/*"] = "/"+path+"/$1"
        }
        return nil
    })
	if err != nil {
		panic(err.Error())
	}

    return rewriteRules
}
