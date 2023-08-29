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

func AddRelation(relation string) {
	var newRelation models.Relation
	var newRelationOp models.Relation
	if checkHashOfNameRelation(relation) {
		// add to hash list
		fmt.Println("A relation with same name already exists. Can not add new relation.")
	} else {
		fmt.Println("If A is " + relation + " of B, then B is what of A ??")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "" {
			fmt.Println("No relation entered. Can not create a relation")
		}
		newRelation = models.Relation{
			Name:     relation,
			Opposite: text,
		}
		newRelationOp = models.Relation{
			Name:     text,
			Opposite: relation,
		}
		globalvar.RelationTypeLists = append(globalvar.RelationTypeLists, newRelation)
		globalvar.RelationTypeLists = append(globalvar.RelationTypeLists, newRelationOp)

	}
	filefunctions.WriteFileRelation(newRelation, "relation.bin")
	filefunctions.WriteFileRelation(newRelationOp, "relation.bin")
}

func checkHashOfNameRelation(name string) bool {
	for _, relation := range globalvar.RelationTypeLists {
		if relation.Name == name {
			return true
		}
	}
	return false
}
