package dt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/disiqueira/gotree"
	"github.com/timdrysdale/gradexpath"
)

func Tree(path string) (string, error) {

	treemap := make(map[string]gotree.Tree)
	treemap["."] = gotree.New(filepath.Base(path))

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if info.IsDir() && strings.Contains(info.Name(), "temp") {
			return filepath.SkipDir
		}
		if info.IsDir() {

			filelist, _ := gradexpath.GetFileList(path)
			numPdf := 0

			for _, file := range filelist {
				if gradexpath.IsPdf(file) {
					numPdf++
				}
				if gradexpath.IsTxt(file) {
					numPdf++
				}
			}

			label := fmt.Sprintf("%3d %s", numPdf, filepath.Base(path))
			if numPdf == 0 {
				label = fmt.Sprintf("  ─ %s", filepath.Base(path)) //U+00B7 ·
			}

			parent := filepath.Dir(path)

			if parent != "" {

				if _, ok := treemap[parent]; !ok {

					treemap[parent] = (treemap["."]).Add(label)

				} else {
					treemap[path] = (treemap[parent]).Add(fmt.Sprintf(label))
				}
			} else {
				treemap[path] = (treemap["."]).Add(path)
			}
		}

		return nil
	})
	if err != nil {
		return "", err

	}

	return treemap["."].Print(), nil
}
