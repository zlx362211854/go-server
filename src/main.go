
package main

import (
	"serverGo/src/routes"
)
type Form struct {
	Name string
}
//var DB, err = scribble.New("data", nil)

func main() {

	//db, err := scribble.New("data", nil)
	//if err != nil {
	//	fmt.Println("Error", err)
	//}
	//form := Form{
	//	Name: "dfgdfd",
	//}
	//if err := db.Write("form", "form1", form); err != nil {
	//	fmt.Println("Error", err)
	//}
	routes.Run()
}
