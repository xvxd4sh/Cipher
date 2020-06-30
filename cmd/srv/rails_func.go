package srv

import (
	"fmt"
)

func printMatrix(randMatrix [][]byte) {
	//looping over 2D slice and extracting 1D slice to val
	for _, val := range randMatrix {
		fmt.Printf("%c", val)
		fmt.Print("\n") // printing each slice
	}
}

func railplacing(args []string, key int) [][]byte {
	upper := false
	lower := true
	counter := 0
	Word := ""

	for x := 0; x < len(args); x++ {
		temp := args[x]
		for y := 0; y < len(temp); y++ {
			Word += string(temp[y])
		}
	}
	//fmt.Println(Word)
	//fmt.Println(len(Word))

	railed := make([][]byte, key)
	for row := range railed {
		railed[row] = make([]byte, len(Word))
	}

	for i := 0; i < len(Word); i++ {

		railed[counter][i] = Word[i]

		if upper {
			counter--
		} else if lower {
			counter++
		}

		if counter <= 0 {
			lower = true
			upper = false
		} else if counter >= key-1 {
			lower = false
			upper = true
		}

	}

	return railed

}

func railplacingstraight(args []string, railed [][]byte) [][]byte {
	Word := ""

	for x := 0; x < len(args); x++ {
		temp := args[x]
		for y := 0; y < len(temp); y++ {
			Word += string(temp[y])
		}
	}

	counter := 0
	for row := 0; row < len(railed); row++ {
		line := railed[row]
		for col := 0; col < len(line); col++ {
			if line[col] != 0 {
				railed[row][col] = Word[counter]
				counter++
			}
		}
	}
	return railed
}

func RailsRead(args []string, key int) {
	railed := railplacingstraight(args, railplacing(args, key))

	printMatrix(railed)

	upper := false
	lower := true
	counter := 0
	finalWord := ""

	for i := 0; i < len(railed[0]); i++ {

		finalWord += string(railed[counter][i])

		if upper {
			counter--
		} else if lower {
			counter++
		}

		if counter <= 0 {
			lower = true
			upper = false
		} else if counter >= key-1 {
			lower = false
			upper = true
		}

	}

	fmt.Println(finalWord)
}

func RailsWrite(args []string, key int) {
	railed := railplacing(args, key)

	printMatrix(railed)

	finalWord := ""

	for i := 0; i < key; i++ {
		temp := railed[i]
		for j := 0; j < len(temp); j++ {
			if temp[j] != 0 {
				finalWord += string(railed[i][j])
			}
		}
	}
	fmt.Println(finalWord)
}
