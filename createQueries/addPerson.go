package createqueries

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	filefunctions "github.com/Hrit99/family-tree/fileFunctions"
	globalvar "github.com/Hrit99/family-tree/globalVar"
	"github.com/Hrit99/family-tree/models"
)

func AddPerson(name string) {
	var newPerson models.Person
	if checkHashOfName(name) {
		// add to hash list
		fmt.Println("A person with same name already exists. Do you want to add person with same name??")
		fmt.Print(`Type yes to add -> `)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("yes", text) == 0 {
			listOfPerson := globalvar.PersonHashMap[name]
			newPerson = models.Person{
				Id:   len(listOfPerson),
				Name: name,
			}
			listOfPerson = append(listOfPerson, newPerson)
			globalvar.PersonHashMap[name] = listOfPerson
		} else {
			fmt.Println("dont add")
		}

	} else {
		newPerson = models.Person{
			Id:   0,
			Name: name,
		}
		listOfPerson := make([]models.Person, 0, 1)
		listOfPerson = append(listOfPerson, newPerson)
		globalvar.PersonHashMap[name] = listOfPerson
	}
	filefunctions.WriteFilePerson(newPerson, "person.bin")
}

func checkHashOfName(name string) bool {
	if globalvar.PersonHashMap[name] != nil {
		return true
	}
	return false
}
