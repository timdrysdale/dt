package dt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/disiqueira/gotree"
	"github.com/timdrysdale/gradexpath"
	pdf "github.com/timdrysdale/unipdf/v3/model"
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
	Name      string
}

func TestWlkTree(t *testing.T) {

	treemap := make(map[string]gotree.Tree)
	treemap["."] = gotree.New("Practice Exam")

	err := filepath.Walk("/home/tim/tg/ingester/tmp-delete-me/usr/exam/Practice Exam Drop Box/", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() && strings.Contains(info.Name(), "temp") {
			return filepath.SkipDir
		}
		if info.IsDir() {

			filelist, _ := gradexpath.GetFileList(path)
			numPdf := 0
			pageCount := 0

			for _, file := range filelist {
				if gradexpath.IsPdf(file) {
					numPdf++
					pc, err := CountPages(file)
					if err == nil {
						pageCount = pageCount + pc
					}
				}
				if gradexpath.IsTxt(file) {
					numPdf++
				}
			}

			label := fmt.Sprintf("%3d %4d %s", numPdf, pageCount, filepath.Base(path))
			if numPdf == 0 {
				label = fmt.Sprintf("  ·   · %s", filepath.Base(path)) //U+00B7
			}

			parent := filepath.Dir(path)

			if parent != "" {

				if _, ok := treemap[parent]; !ok {

					fmt.Println(parent)
					fmt.Println(path)

					treemap[parent] = (treemap["."]).Add(label)

				} else {
					//foo := (treemap[parent]).Add(path)
					treemap[path] = (treemap[parent]).Add(fmt.Sprintf(label))
				}
			} else {

				treemap[path] = (treemap["."]).Add(path)
			}

		}

		//fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		return
	}

	fmt.Println(treemap["."].Print())
}

func CountPages(inputPath string) (int, error) {

	numPages := 0

	f, err := os.Open(inputPath)
	if err != nil {
		return numPages, err
	}

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return numPages, err
	}

	defer f.Close()

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return numPages, err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return numPages, err
		}
	}

	numPages, err = pdfReader.GetNumPages()
	if err != nil {
		return numPages, err
	}

	return numPages, nil

}
