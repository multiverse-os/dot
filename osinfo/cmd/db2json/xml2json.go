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
	fmt.Println("[ Download `libosinfo` Database Convert to JSON ]")
	fmt.Println("=================================================")
	fmt.Println("A simple tool to make accessing the libosinfo data")
	fmt.Println("easier by converting to json, and providing a script")
	fmt.Println("script to download the newest version of the db.\n\n")

	os.RemoveAll("./osinfo-db/")
	_, _ = git.PlainClone("osinfo-db", false, &git.CloneOptions{
		URL:      gitRepository,
		Progress: os.Stdout,
	})

	xmlFiles, _ := filepath.Glob("osinfo-db/data/*/*/*.xml.in")
	for _, xmlFilename := range xmlFiles {
		xmlReader, err := os.Open(xmlFilename)
		if err != nil {
			fmt.Println("[error] failed to open xml:", err)
		}
		jsonOutput := &bytes.Buffer{}
		jsonOutput, err = xj.Convert(xmlReader)
		if err != nil {
			fmt.Println("[error] failed to convert xml to json:", err)
		}

		filenameWithoutExt := strings.Split(xmlFilename, ".xml.in")
		jsonFilename := filenameWithoutExt[0] + ".json"

		prettyJson := &bytes.Buffer{}
		json.Indent(prettyJson, jsonOutput.Bytes(), "", "    ")

		err = ioutil.WriteFile(jsonFilename, prettyJson.Bytes(), 0644)
		if err != nil {
			fmt.Println("[error] failed to write json to file:", err)
		}

		os.Remove(xmlFilename)
	}
	os.RemoveAll("../../db")
	os.Rename("./osinfo-db/data", "../../db")
	os.RemoveAll("./osinfo-db/")
}
