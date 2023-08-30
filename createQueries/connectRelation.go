package createqueries

import (
	"fmt"

	filefunctions "github.com/Hrit99/family-tree/fileFunctions"
	"github.com/Hrit99/family-tree/models"
)

func ConnectRelation(personOne *models.Person, personTwo *models.Person, relation *models.Relation) {
	if personOne.Name == "" || personTwo.Name == "" || relation.Name == "" {
		fmt.Println("unable to connect")
	} else if personOne.Gender != relation.Gender {
		fmt.Println("unable to connect. Gender does not match")
	} else {
		toAddConnection(personOne, personTwo, relation.Name)

		if personOne.Gender == models.Female.String() && relation.FemaleOpposite != "" {
			toAddConnection(personTwo, personOne, relation.FemaleOpposite)
		} else if personOne.Gender == models.Male.String() && relation.MaleOpposite != "" {
			toAddConnection(personTwo, personOne, relation.MaleOpposite)
		} else if relation.Opposite != "" {
			toAddConnection(personTwo, personOne, relation.Opposite)
		}
		filefunctions.UpdateFilePerson(*personOne)
		filefunctions.UpdateFilePerson(*personTwo)
	}
}

func toAddConnection(personOne *models.Person, personTwo *models.Person, relationName string) {
	val, ok := personTwo.RelationMapForward[relationName]
	if ok {
		flag := true
		for _, v := range val[personOne.Name] {
			if v == personOne.Id {
				flag = false
			}
		}
		if flag {
			val[personOne.Name] = append(val[personOne.Name], personOne.Id)
			personTwo.RelationMapForward[relationName] = val
		}
	} else {
		arr := make(map[string][]int)
		li := make([]int, 0, 1)
		li = append(li, personOne.Id)
		arr[personOne.Name] = li
		personTwo.RelationMapForward[relationName] = arr
	}
}
