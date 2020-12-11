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

// Generate generates the site
func Generate(files []string) error {
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

		if _, err = os.Stat("./public/" + strings.TrimSuffix(files[i], filepath.Ext(files[i])) + ".html"); os.IsNotExist(err) {
			os.MkdirAll("./public/"+filepath.Dir(files[i]), 0777)
		}

		output, err := os.Create("./public/" + strings.TrimSuffix(files[i], filepath.Ext(files[i])) + ".html")
		if err != nil {
			return fmt.Errorf("Error creating output file %v: %v", "public/"+strings.TrimSuffix(files[i], filepath.Ext(files[i]))+".html", err)
		}
		defer output.Close()

		output.Write(result)
	}

	return nil
}
