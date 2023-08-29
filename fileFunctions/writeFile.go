package filefunctions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Hrit99/family-tree/models"
)

func WriteFilePerson(newPerson models.Person, fileName string) {
	b, err := json.Marshal(newPerson)
	if err != nil {
		fmt.Println(err)
		return
	}
	bitString := stringToBin(string(b))
	encryptStr := encryptString(bitString)
	writeInFile(encryptStr, fileName)
}

func WriteFileRelation(newRelation models.Relation, fileName string) {
	b, err := json.Marshal(newRelation)
	if err != nil {
		fmt.Println(err)
		return
	}
	bitString := stringToBin(string(b))
	encryptStr := encryptString(bitString)
	writeInFile(encryptStr, fileName)
}

func writeInFile(encryptStr string, fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("File exists\n")
		myfile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer myfile.Close()

		// Write the string to the file
		_, err = myfile.WriteString(" " + encryptStr)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("File does not exist\n")
		os.WriteFile(fileName, []byte(encryptStr), 0644)
	}
}

func encryptString(s string) string {
	lenB := len(s)/8 + 1
	bs := make([]byte, lenB)

	count, i, zeroToAdd := 0, 0, 0
	var now byte
	for _, v := range s {
		if count == 8 {
			bs[i] = now
			i++
			now, count = 0, 0
		}
		now = now<<1 + byte(v-'0')
		count++
	}
	if count != 0 {
		zeroToAdd = count
		bs[i] = now
		i++
	}

	bs = bs[:i:i]
	str := fmt.Sprintf("%v%s", zeroToAdd, string(bs))
	return str
}
func stringToBin(s string) (binString string) {
	for _, c := range s {
		str := fmt.Sprintf("%b", c)

		for i := 0; i < (7 - len(str)); i++ {
			str = fmt.Sprintf("%v%s", "0", str)
		}
		binString = fmt.Sprintf("%s%v", binString, str)
	}
	return
}
