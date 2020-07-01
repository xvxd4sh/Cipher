package srv

import "fmt"

func DcolRead(args []string, key1 string, key2 string) {
	final := ColumnaRead(args, key2)
	final2 := ColumnaRead(final, key1)
	fmt.Println(final2)
}

func DcolWrite(args []string, key1 string, key2 string) {
	final := ColumnaWrite(args, key1)
	//fmt.Println(final)
	final2 := ColumnaWrite(final, key2)
	fmt.Println(final2)
}
