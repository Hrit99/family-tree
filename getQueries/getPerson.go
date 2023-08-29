package getqueries

import (
	"bufio"
	"fmt"
	"os"

	globalvar "github.com/Hrit99/family-tree/globalVar"
	"github.com/Hrit99/family-tree/models"
)

func GetPersonWithName(name string) *models.Person {

	if len(globalvar.PersonHashMap[name]) > 1 {
		buf := make([]byte, 1024)
		fmt.Println("More than one user with same name exists. Enter the id of the desired user : ")
		reader := bufio.NewReader(os.Stdin)
		id, _ := reader.Read(buf)
		return &globalvar.PersonHashMap[name][id]

	} else if len(globalvar.PersonHashMap[name]) == 1 {
		return &globalvar.PersonHashMap[name][0]
	}
	return &models.Person{}
}
