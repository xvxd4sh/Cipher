package srv

import (
	"fmt"
	"math"
)

func roundup(num int) int {
	numround := math.Round(float64(num))
	if (numround - float64(num)) > .5 {
		numround += 1.0
	}
	return int(numround)
}

func findint(key int, list []int) int {
	temp := 0
	for i := 0; i < len(list); i++ {
		if key == list[i] {
			temp = i
		}
	}
	return temp
}

func makeabyteMatrix(rows int, col int) [][]byte {
	list := make([][]byte, rows)
	for row := range list {
		list[row] = make([]byte, col)
	}
	return list
}

func numerify(key string) []int {
	corectkey := ""
	for i := 0; i < len(key); i++ {
		if key[i] >= 'a' && key[i] <= 'z' {
			corectkey += string(key[i] - 'a' + 'A')

		} else {
			corectkey += string(key[i])
		}
	}
	bits := []byte(corectkey)
	numlist := make([]int, len(corectkey))
	place := 0

	for place != (len(key)) {
		holder := 0
		temp := bits[0]
		for i := 0; i < len(corectkey); i++ {
			if temp > bits[i] && bits[i] != 'a' {
				temp = bits[i]
				holder = i
			}
		}
		numlist[holder] = place
		place++
		bits[holder] = 'a'
	}

	return numlist
}

func stringify(strings []string) string {
	word := ""
	for x := 0; x < len(strings); x++ {
		temp := strings[x]
		for y := 0; y < len(temp); y++ {
			word += string(temp[y])
		}
	}
	return word
}

func ColumnaRead(args []string, key string) {
	numlist := numerify(key)
	gentext := stringify(args)
	keyint := roundup(len(gentext) / len(numlist))
	list := makeabyteMatrix(keyint, len(key))
	placement := 0

	//	fmt.Println(gentext)
	//	fmt.Println(len(gentext))

	for x := 0; x < len(key); x++ {
		for y := 0; y < keyint; y++ {
			if placement < len(gentext) {
				list[y][findint(x, numlist)] = gentext[placement]
				placement++
			}
		}
	}
	final := ""
	for i := 0; i < keyint; i++ {
		for j := 0; j < len(key); j++ {
			final += string(list[i][j])
		}
	}
	fmt.Println(final)
}

func ColumnaWrite(args []string, key string) {
	numlist := numerify(key)
	gentext := stringify(args)
	keyint := roundup(len(gentext) / len(numlist))
	list := makeabyteMatrix(keyint, len(key))
	placement := 0

	for x := 0; x < keyint; x++ {
		for y := 0; y < len(key); y++ {
			if placement < len(gentext) {
				list[x][y] = gentext[placement]
				placement++
			} else {
				list[x][y] = 'x'
			}
		}
	}
	//fmt.Printf("%c",list)
	final := ""
	for i := 0; i < len(numlist); i++ {
		for j := 0; j < keyint; j++ {
			final += string(list[j][findint(i, numlist)])
		}
	}
	fmt.Println(final)
}
