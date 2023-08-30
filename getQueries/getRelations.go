package getqueries

import (
	globalvar "github.com/Hrit99/family-tree/globalVar"
	"github.com/Hrit99/family-tree/models"
)

// func GetRelations() {
// 	for _, v := range globalvar.RelationTypeLists {
// 		// fmt.Println(v.Name)
// 	}
// }

func GetRelation(s string) *models.Relation {
	for _, v := range globalvar.RelationTypeLists {
		if v.Name == s {
			return &v
		}
	}
	return &models.Relation{}
}
