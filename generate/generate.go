package generate

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

const directory = "markview-public"

// Generate generates the site
func Generate(files []string) error {
	pageNames := []string{}
	pagePaths := []string{}

	for i := range files {
		file, err := ioutil.ReadFile("./" + files[i])
		if err != nil {
			return fmt.Errorf("Error opening file %v: %v", files[i], err)
		}

		parsed := new(bytes.Buffer)

		err = goldmark.Convert(file, parsed)
		if err != nil {
			return fmt.Errorf("Error parsing file %v: %v", files[i], err)
		}

		result, err := ioutil.ReadAll(parsed)
		if err != nil {
			return fmt.Errorf("Error reading result: %v", err)
		}

		_, err = os.Stat(directory + "/" + strings.TrimSuffix(files[i], filepath.Ext(files[i])) + ".html")
		if os.IsNotExist(err) {
			os.MkdirAll(directory+"/"+filepath.Dir(files[i]), 0777)
		}

		output, err := os.Create(directory + "/" + strings.TrimSuffix(files[i], filepath.Ext(files[i])) + ".html")
		if err != nil {
			return fmt.Errorf("Error creating output file %v: %v", directory+"/"+strings.TrimSuffix(files[i], filepath.Ext(files[i]))+".html", err)
		}
		defer output.Close()

		output.Write(result)

		pageNames = append(pageNames, strings.TrimSuffix(files[i], filepath.Ext(files[i])))
		pagePaths = append(pagePaths, strings.TrimSuffix(files[i], filepath.Ext(files[i])) + ".html")
	}

	generateHomepage(pageNames, pagePaths, directory)

	return nil
}
