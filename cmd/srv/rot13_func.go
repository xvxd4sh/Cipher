package srv

import (
	"fmt"
)

func Rot13Read(args []string) {

	temp2 := ""

	for i := 0; i < len(args); i++ {
		temp := args[i]

		for i := 0; i < len(temp); i++ {
			if (temp[i] >= 'A' && temp[i] <= 'M') || (temp[i] >= 'a' && temp[i] <= 'm') {
				temp2 = temp2 + string(temp[i]+13)
			} else {
				temp2 = temp2 + string(temp[i]-13)
			}
		}
		temp2 += " "
	}
	fmt.Println(temp2)

}

func Rot13Write(args []string) {
	temp2 := ""
	for i := 0; i < len(args); i++ {
		temp := args[i]

		for i := 0; i < len(temp); i++ {
			if (temp[i] >= 'A' && temp[i] <= 'M') || (temp[i] >= 'a' && temp[i] <= 'm') {
				temp2 = temp2 + string(temp[i]+13)
			} else {
				temp2 = temp2 + string(temp[i]-13)
			}
		}
		temp2 += " "
	}
	fmt.Println(temp2)
}
