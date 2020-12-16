package generate

import (
	"fmt"
	"os"
)

func generateHomepage(pageNames, pagePaths []string, directory string) (error) {
	result := ""
	
	for i := range pageNames {
		result = result + "<a href='"+pagePaths[i]+"'>"+pageNames[i]+"</a>\n"
	}

	_, err := os.Stat(directory + "/index.html")
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("Error checking for index.html: %v", err)	
		}
	}
	if err == nil {
		return nil
	}

	file, err := os.Create(directory+"/index.html")
	if err != nil {
		return fmt.Errorf("Error opening index.html: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(result))
	if err != nil {
		return fmt.Errorf("Error writing index.html: %v", err)
	}

	return nil
}