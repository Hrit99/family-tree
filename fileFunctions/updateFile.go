package filefunctions

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Hrit99/family-tree/models"
)

func UpdateFilePerson(newPerson models.Person) {
	str := getEncyptedPersonStr(newPerson)
	for _, person := range ReadFilePerson() {
		if (person.Name == newPerson.Name) && (person.Id == newPerson.Id) {
			oldstr := getEncyptedPersonStr(person)
			bb, _ := os.ReadFile("person.bin")
			os.WriteFile("person.bin", []byte(strings.Replace(string(bb), oldstr, str, -1)), 0644)
		}
	}
}

// func readFileUpdate(s string) []string {
// 	bb, err := os.ReadFile(s)
// 	if err != nil {
// 		fmt.Println("error occured")
// 		return make([]string, 0)
// 	}

// 	return strings.Split(string(bb), " ")

// }

func getEncyptedPersonStr(newPerson models.Person) string {
	b, err := json.Marshal(newPerson)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bitString := stringToBin(string(b))
	return encryptString(bitString)
}
