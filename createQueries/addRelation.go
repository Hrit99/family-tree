package createqueries

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	filefunctions "github.com/Hrit99/family-tree/fileFunctions"
	getqueries "github.com/Hrit99/family-tree/getQueries"
	globalvar "github.com/Hrit99/family-tree/globalVar"
	"github.com/Hrit99/family-tree/models"
)

func AddRelation(relation string, gender string) {
	var newRelation models.Relation
	var newRelationOpm models.Relation
	var newRelationOpf models.Relation
	var newRelationOp models.Relation
	// var newRelationOp models.Relation

	if gender == "male" {
		gender = models.Male.String()
	} else if gender == "female" {
		gender = models.Female.String()
	} else {
		gender = models.Unknown.String()
	}
	fmt.Println("If A is " + relation + " of B, then B is what of A ??")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("For male (Press enter if NA) : ")
	textm, _ := reader.ReadString('\n')
	textm = strings.Replace(textm, "\n", "", -1)
	fmt.Println("For female (Press enter if NA) : ")
	textf, _ := reader.ReadString('\n')
	textf = strings.Replace(textf, "\n", "", -1)
	fmt.Println("Common for both gender (Press enter if NA) : ")
	textc, _ := reader.ReadString('\n')
	textc = strings.Replace(textc, "\n", "", -1)
	newRelation = models.Relation{
		Name:           relation,
		MaleOpposite:   textm,
		FemaleOpposite: textf,
		Opposite:       textc,
		Gender:         gender,
	}
	if newRelation.Gender == models.Male.String() {
		if textm != "" {
			newRelationOpm = models.Relation{
				Name:           textm,
				MaleOpposite:   relation,
				Gender:         models.Male.String(),
				FemaleOpposite: getqueries.GetRelation(textm).FemaleOpposite,
				Opposite:       getqueries.GetRelation(textm).Opposite,
			}
		}
		if textf != "" {
			newRelationOpf = models.Relation{
				Name:           textf,
				MaleOpposite:   relation,
				Gender:         models.Female.String(),
				FemaleOpposite: getqueries.GetRelation(textf).FemaleOpposite,
				Opposite:       getqueries.GetRelation(textf).Opposite,
			}
		}
		if textc != "" {
			newRelationOp = models.Relation{
				Name:           textc,
				MaleOpposite:   relation,
				Gender:         models.Unknown.String(),
				FemaleOpposite: getqueries.GetRelation(textc).FemaleOpposite,
				Opposite:       getqueries.GetRelation(textc).Opposite,
			}
		}
	} else if newRelation.Gender == models.Female.String() {
		if textm != "" {
			newRelationOpm = models.Relation{
				Name:           textm,
				FemaleOpposite: relation,
				Gender:         models.Male.String(),
				MaleOpposite:   getqueries.GetRelation(textm).MaleOpposite,
				Opposite:       getqueries.GetRelation(textm).Opposite,
			}
		}
		if textf != "" {
			newRelationOpf = models.Relation{
				Name:           textf,
				FemaleOpposite: relation,
				Gender:         models.Female.String(),
				MaleOpposite:   getqueries.GetRelation(textf).MaleOpposite,
				Opposite:       getqueries.GetRelation(textf).Opposite,
			}
		}
		if textc != "" {
			newRelationOp = models.Relation{
				Name:           textc,
				FemaleOpposite: relation,
				Gender:         models.Unknown.String(),
				MaleOpposite:   getqueries.GetRelation(textc).MaleOpposite,
				Opposite:       getqueries.GetRelation(textc).Opposite,
			}
		}
	}
	globalvar.RelationTypeLists[newRelation.Name] = newRelation
	globalvar.RelationTypeLists[newRelationOpm.Name] = newRelationOpm
	globalvar.RelationTypeLists[newRelationOpf.Name] = newRelationOpf
	globalvar.RelationTypeLists[newRelationOp.Name] = newRelationOp

	checkHashOfNameRelation(newRelation)
	checkHashOfNameRelation(newRelationOpm)
	checkHashOfNameRelation(newRelationOpf)
	checkHashOfNameRelation(newRelationOp)
}

func checkHashOfNameRelation(rel models.Relation) {
	for _, relation := range globalvar.RelationTypeLists {
		if relation.Name == rel.Name {
			filefunctions.UpdateFileRelation(rel)
		}
	}
	filefunctions.WriteFileRelation(rel, "relation.bin")
}
