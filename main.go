package main

import (
	"github.com/Hrit99/family-tree/cmd"
	filefunctions "github.com/Hrit99/family-tree/fileFunctions"
)

func main() {
	cmd.Execute()
	// initPerson()
	// initWedlock()
	filefunctions.ReadFilePerson()
	// filefunctions.ReadFileRelation()
}

// func initHash() {
// 	globalvar.PersonHashMap = make(map[string][]models.Person)
// }
