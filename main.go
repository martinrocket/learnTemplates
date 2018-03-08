//Build a Go program that utilizes Templates for webpages
package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

type webContent struct {
	Footer string
	Header string
	Button template.HTML
}

func main() {
	ti := time.Now()

	myTimeStr := "footer - " + ti.Format("Mon Jan 02 2006 15:04 PM") + " - footer"
	myWebContent := webContent{
		Footer: myTimeStr,
		Header: "my new header",
		Button: template.HTML("my button"),
	}
	//hf := []string{"my header from slice", "my footer from slice"}
	// hf1 := map[string]string{

	// 	"footer": myTimeStr,
	// 	"header": "header from map!",
	// }

	t, err := template.ParseFiles("tmpl1.html")
	fmt.Printf("Parsed file %v with error = %v.\n", t.Name(), err)
	t.Execute(os.Stdout, myWebContent)

	f, err := os.Create("output.html")
	defer f.Close()
	fmt.Printf("\nCreated the file %v with error = %v.\n", f.Name(), err)
	t.Execute(f, myWebContent)
	fmt.Println(myWebContent)

}
