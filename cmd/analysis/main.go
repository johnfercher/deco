package main

import (
	"github.com/johnfercher/chaos/struct/structcore/structconsts/content"
	"github.com/johnfercher/chaos/struct/structservices"
	"log"
)

func main() {
	file := structservices.NewFile()
	classifier := structservices.NewFileClassifier()
	discover := structservices.NewDiscover(file, classifier)
	packages, err := discover.Project("docs/examples/medium")
	if err != nil {
		log.Fatal(err)
	}

	for key, p := range packages {
		if p.ContentType != content.Go {
			delete(packages, key)
		}
	}

	for _, p := range packages {
		p.Print("")
	}
}
