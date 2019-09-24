package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	xj "github.com/basgys/goxml2json"
	git "gopkg.in/src-d/go-git.v4"
)

const gitRepository = "https://gitlab.com/libosinfo/osinfo-db"

func main() {
	os.RemoveAll("./osinfo-db/")
	_, _ = git.PlainClone("osinfo-db", false, &git.CloneOptions{
		URL:      gitRepository,
		Progress: os.Stdout,
	})

	xmlFiles, _ := filepath.Glob("osinfo-db/data/*/*/*.xml.in")
	for _, xmlFilename := range xmlFiles {
		fmt.Println(xmlFilename)

		xmlReader, _ := os.Open(xmlFilename)

		//xml := strings.NewReader(osXML)
		jsonOutput, err := xj.Convert(xmlReader)
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonOutput.String())

		filenameWithoutExt := strings.Split(xmlFilename, ".xml.in")
		jsonFilename := filenameWithoutExt[0] + ".json"
		prettyJson := &bytes.Buffer{}
		if err := json.Indent(prettyJson, jsonOutput.Bytes(), "", "  "); err != nil {
			panic(err)
		}
		fmt.Println("jsonOutput:  ", prettyJson.String())
		fmt.Println("jsonFilename:", jsonFilename)

		_ = ioutil.WriteFile(jsonFilename, prettyJson.Bytes(), 0644)

		os.Remove(xmlFilename)
		os.Rename("./osinfo-db/data", "../../db")
		os.RemoveAll("./osinfo-db/")
	}
}
