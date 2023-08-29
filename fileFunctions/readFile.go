package filefunctions

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Hrit99/family-tree/models"
)

func BinToString(s string) string {
	finalStr := ""
	count := 0
	str := ""
	for _, v := range s {
		if count == 7 {
			if i, err := strconv.ParseInt(str, 2, 64); err != nil {
				fmt.Println(err)
			} else {
				finalStr = fmt.Sprintf("%s%v", finalStr, string(rune(i)))
			}
			count = 0
			str = ""
		}
		str = fmt.Sprintf("%s%v", str, string(v))
		count++
	}
	if count != 0 {
		if i, err := strconv.ParseInt(str, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			finalStr = fmt.Sprintf("%s%v", finalStr, string(rune(i)))
		}
	}
	return finalStr
}

func decryptString(s []byte) string {
	leftbit := s[0]
	s = s[1:]
	var finalStr string
	rounds := 0
	for _, b := range s {
		rounds++
		var str string
		bi, err := fmt.Printf("%b", b)
		if err != nil {
			fmt.Println("error")
		} else {
			str = fmt.Sprintf("%b", b)
			zta := (8 - bi)
			if rounds == len(s) {
				i, err := strconv.Atoi(string(leftbit))
				if err != nil {
					fmt.Println(err)
				}
				zta = i - bi
			}
			for i := 0; i < zta; i++ {
				str = fmt.Sprintf("%v%s", "0", str)
			}
		}
		finalStr = fmt.Sprintf("%s%v", finalStr, str)
	}
	return finalStr
}

func ReadFilePerson() []models.Person {

	bb, err := os.ReadFile("person.bin")
	if err != nil {
		fmt.Println("error occured")
		return make([]models.Person, 0)
	}

	persons := strings.Split(string(bb), " ")

	arrToReturn := make([]models.Person, len(persons))

	for _, person := range persons {
		decryptStr := decryptString([]byte(person))
		decodedStr := BinToString(decryptStr)
		decoded := models.Person{}
		json.Unmarshal([]byte(decodedStr), &decoded)
		println("hola")
		fmt.Println(decoded.RelationMapForward)
		if decoded.RelationMapForward == nil {
			decoded.RelationMapForward = make(map[string]map[string]int)
		}
		if decoded.RelationMapBackward == nil {
			decoded.RelationMapBackward = make(map[string]map[string]int)
		}
		arrToReturn = append(arrToReturn, decoded)
	}
	return arrToReturn
}

func ReadFileRelation() []models.Relation {

	bb, err := os.ReadFile("relation.bin")
	if err != nil {
		fmt.Println("error occured")
		return make([]models.Relation, 0)
	}

	relations := strings.Split(string(bb), " ")

	arrToReturn := make([]models.Relation, len(relations))

	for _, relation := range relations {
		decryptStr := decryptString([]byte(relation))
		decodedStr := BinToString(decryptStr)
		decoded := models.Relation{}
		json.Unmarshal([]byte(decodedStr), &decoded)
		arrToReturn = append(arrToReturn, decoded)
	}
	return arrToReturn
}
