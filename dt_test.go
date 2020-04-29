package dt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/disiqueira/gotree"
)

func TestTree(t *testing.T) {

	artist := gotree.New("Pantera")

	treemap := make(map[string]gotree.Tree)

	treemap["."] = artist

	foo := (treemap["."]).Add("Bar")

	treemap["bar"] = foo

	album := artist.Add("Far Beyond Driven")
	album2 := artist.Add("Foo")
	album3 := artist.Add("Bar")

	album.Add("5 minutes Alone")
	album2.Add("2 minutes Alone")
	album3.Add("1 minutes Alone")
	album3.Add("0 minutes Alone")
	fmt.Println(artist.Print())
	fmt.Printf("%T\n", album3)

}

type Node struct {
	Tree      gotree.Tree
	FileCount int
	PageCount int
}

func TestWlkTree(t *testing.T) {

	treemap := make(map[string]gotree.Tree)
	treemap["."] = gotree.New("Exam123")

	err := filepath.Walk("/home/tim/tg/ingester/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {

			parent := strings.TrimSuffix(path, filepath.Base(path))
			if parent != "" {
				//fmt.Printf("%s->%s\n", parent, path)
				if _, ok := treemap[parent]; !ok {
					treemap[parent] = (treemap["."]).Add(filepath.Base(path))
				} else {
					//foo := (treemap[parent]).Add(path)
					treemap[path] = (treemap[parent]).Add(filepath.Base(path))
				}
			} else {
				treemap[path] = (treemap["."]).Add(path)
			}

		}

		//if info.IsDir() { // && info.Name() != DirToWalk {
		//	return filepath.SkipDir
		//}
		//fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		return
	}

	fmt.Println(treemap["."].Print())
}
