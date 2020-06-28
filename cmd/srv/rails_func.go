package srv

import "fmt"

func railplacing(args []string, key int) {
	upper := false
	lower := true
	counter := 0
	Word := ""
	var railed [][]byte
	for x := 0; x < len(args); x++ {
		temp := args[x]
		for y := 0; y < len(temp); y++ {
			Word += string(temp[y])
		}
	}
	fmt.Println(Word)
	fmt.Println(len(Word))
	for i := 0; i < len(Word); i++ {
		railed[counter][i] = Word[i]
		if upper {
			counter--
		} else if lower {
			counter++
		}

		if counter == 0 {
			lower = true
			upper = false
		} else if counter == key-1 {
			lower = false
			upper = true
		}

	}

}

func RailsRead(args []string, key int) {
	railplacing(args, key)
}

func RailsWrite(args []string, key int) {

}
