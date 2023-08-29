package createqueries

import (
	"fmt"

	filefunctions "github.com/Hrit99/family-tree/fileFunctions"
	"github.com/Hrit99/family-tree/models"
)

func ConnectRelation(personOne *models.Person, personTwo *models.Person, relation *models.Relation) {
	if personOne.Name == "" || personTwo.Name == "" || relation.Name == "" {
		fmt.Println("unable to connect")
	} else {
		val, ok := personOne.RelationMapForward[relation.Name]
		if ok {

			_, okk := val[personTwo.Name]
			if okk {
				//do nothing
			} else {
				val[personTwo.Name] = personTwo.Id
			}

			personOne.RelationMapForward[relation.Name] = val
		} else {
			arr := make(map[string]int)
			arr[personTwo.Name] = personTwo.Id
			personOne.RelationMapForward[relation.Name] = arr
		}

		val, ok = personTwo.RelationMapForward[relation.Opposite]
		if ok {
			_, okk := val[personOne.Name]
			if okk {
				//do nothing
			} else {
				val[personOne.Name] = personOne.Id
			}
			personTwo.RelationMapForward[relation.Opposite] = val
		} else {
			arr := make(map[string]int)
			arr[personOne.Name] = personOne.Id
			personTwo.RelationMapForward[relation.Opposite] = arr
		}
	}
	filefunctions.UpdateFilePerson(*personOne)
	filefunctions.UpdateFilePerson(*personTwo)
}
