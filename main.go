//Build a Go program that utilizes Templates for webpages
package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func main() {
	ti := time.Now()

	//ti.Format("Mon Jan 01-02-2006 15:04 PM")

	myTimeStr := "footer - " + ti.Format("Mon Jan 02 2006 15:04 PM") + " - footer"

	//hf := []string{"my header from slice", "my footer from slice"}
	hf1 := map[string]string{

		"footer": myTimeStr,
		"header": "header from map!",
	}

	t, err := template.ParseFiles("tmpl1.html")
	fmt.Printf("Parsed file %v with error = %v.\n", t.Name(), err)
	t.Execute(os.Stdout, hf1)

	f, err := os.Create("output.html")
	defer f.Close()
	fmt.Printf("\nCreated the file %v with error = %v.\n", f.Name(), err)
	t.Execute(f, hf1)
	fmt.Println(hf1)

}
