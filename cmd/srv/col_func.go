package srv

import (
	"math"
)

func roundup(num float64) float64 {
	numround := math.Round(num)
	//fmt.Println(numround - num)
	if (numround - num) < 0.0 {
		numround += 1.0
	}
	//fmt.Println(numround)
	return numround
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

func ColumnaRead(args []string, key string) []string {
	numlist := numerify(key)
	gentext := stringify(args)
	keyintfloat := roundup(float64(len(gentext)) / float64(len(key)))
	keyint := int(keyintfloat)
	list := makeabyteMatrix(keyint, len(key))
	placement := 0

	leftover := ((keyint * len(key)) - len(gentext))
	//fmt.Println(leftover)
	if leftover != 0 {
		counter := 0
		for 0 < leftover {
			list[keyint-1][(len(key)-1)-counter] = '*'
			counter++
			leftover--
		}
	}
	//	fmt.Println(gentext)
	//	fmt.Println(len(gentext))

	for x := 0; x < len(key); x++ {
		for y := 0; y < keyint; y++ {
			if placement < len(gentext) {
				if list[y][findint(x, numlist)] != '*' {
					list[y][findint(x, numlist)] = gentext[placement]
					placement++
				}
			}
		}
	}
	final := ""
	finalList := make([]string, 1)

	for i := 0; i < keyint; i++ {
		for j := 0; j < len(key); j++ {
			if list[i][j] != '*' {
				final += string(list[i][j])
			}
		}
	}

	//fmt.Println(final)
	finalList[0] = final
	return finalList
}

func ColumnaWrite(args []string, key string) []string {
	numlist := numerify(key)
	gentext := stringify(args)
	keyintfloat := roundup(float64(len(gentext)) / float64(len(key)))
	keyint := int(keyintfloat)
	list := makeabyteMatrix(keyint, len(key))
	placement := 0

	for x := 0; x < keyint; x++ {
		for y := 0; y < len(key); y++ {
			if placement < len(gentext) {
				list[x][y] = gentext[placement]
				placement++
			} else {
				list[x][y] = '*'
			}
		}
	}
	//fmt.Printf("%c",list)
	final := ""
	finalList := make([]string, 1)
	for i := 0; i < len(numlist); i++ {
		for j := 0; j < keyint; j++ {
			if list[j][findint(i, numlist)] != '*' {
				final += string(list[j][findint(i, numlist)])
			}
		}
	}
	//fmt.Println(final)
	finalList[0] = final
	return finalList
}
