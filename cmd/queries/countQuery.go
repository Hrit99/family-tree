package queries

import (
	"fmt"

	"github.com/Hrit99/family-tree/models"
)

func CountQuery(person models.Person, relation models.Relation) {
	count := 0
	for _, v := range person.RelationMapForward[relation.Name] {
		count = count + len(v)
	}
	fmt.Println(count)
}
