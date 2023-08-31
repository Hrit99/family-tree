package queries

import (
	"fmt"

	"github.com/Hrit99/family-tree/models"
)

func GetRelatives(relation models.Relation, person models.Person) {
	for k, v := range person.RelationMapForward[relation.Name] {
		for i := 0; i < len(v); i++ {
			fmt.Print(k + " ")
		}
	}
}
