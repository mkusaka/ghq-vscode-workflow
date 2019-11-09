package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	startDir     string
	minimumScore float64
	wf           *aw.Workflow
)

func init() {
	startDir = filepath.Join(os.Getenv("HOME"), "src")
	wf = aw.New()
}

// repo path: site/user(org)name/repo
func readRepositoryDir(dirpath string, depth int) []string {
	paths := []string{}
	files, _ := ioutil.ReadDir(dirpath)

	for _, file := range files {
		if !file.IsDir() || strings.HasPrefix(file.Name(), ".") {
			continue
		}

		filepath := filepath.Join(dirpath, file.Name())
		if depth <= 2 {
			for _, path := range readRepositoryDir(filepath, depth+1) {
				paths = append(paths, path)
			}
		} else {
			paths = append(paths, filepath)
		}
	}

	return paths
}

func run() {
	var query string

	if args := wf.Args(); len(args) > 0 {
		query = args[0]
	}

	for _, path := range readRepositoryDir(startDir, 1) {
		log.Printf("%s", path)
		it := wf.NewFileItem(path)
		it.NewModifier("cmd").
			Subtitle("Browse in Alfred")
	}

	if query != "" {
		res := wf.Filter(query)

		log.Printf("%d results match \"%s\"", len(res), query)

		for i, r := range res {
			log.Printf("%02d. score=%0.1f sortkey=%s", i+1, r.Score, wf.Feedback.Keywords(i))
		}
	}

	wf.WarnEmpty("No matching folders found", "Try a different query?")

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
