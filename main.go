//Build a Go program that utilizes Templates for webpages
package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {

	//hf := []string{"my header from slice", "my footer from slice"}
	hf1 := make(map[string]string)

	hf1["footer"] = "footer from map"
	hf1["header"] = "header from map"

	t, err := template.ParseFiles("tmpl1.html")
	fmt.Printf("Parsed file %v with error = %v.\n", t.Name(), err)
	t.Execute(os.Stdout, hf1)

	f, err := os.Create("output.html")
	defer f.Close()
	fmt.Printf("\nCreated the file %v with error = %v.\n", f.Name(), err)
	t.Execute(f, hf1)
	fmt.Println(hf1)

}
