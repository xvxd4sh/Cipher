package srv

import (
	"fmt"
)

var Capabet = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var Lowebet = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

/*

Over-engineering ** proceeded with mathematical approach

func R_rotator(arg []byte) []byte {
	//right shift of array
	alp := arg
	holder := alp[0]
	for i := 1; i < len(alp); i++ {
		temp := alp[i]
		alp[i] = holder
		holder = temp
	}
	alp[0] = holder
	return alp
}

func L_rotator(args []byte) []byte {
	//left shift of array
	alp := args
	holder := alp[0]
	for i := 1; i < len(alp); i++ {
		alp[i-1] = alp[i]
	}
	alp[len(alp)-1] = holder
	return alp
}

func keybet(alphabet []byte, key int) []byte {
	rotated := alphabet
	if key > 0 {
		for i := 0; i < key; i++ {
			rotated = L_rotator(rotated)
		}
	} else {
		for i := key; i < 0; i++ {
			rotated = R_rotator(rotated)
		}
	}
	return rotated
}
*/

func CaesarRead(args []string, key int) {
	//fmt.Println("reading")
	//UpperalphabetList := keybet(Capabet, key)
	//LoweralphaberList := keybet(Lowebet, key)
	final := ""
	sample := args

	for i := 0; i < len(sample); i++ {
		word := sample[i]
		for j := 0; j < len(word); j++ {
			if key > 0 {
				if word[j] > 'A' && word[j] < 'Z' {
					final += string(Capabet[(int(word[j]-'A')-key)%26])
				} else {
					final += string(Lowebet[(int(word[j]-'a')-key)%26])
				}
			} else { /*else if key < 0 {
					TO:DO - Work on the negative key caesar cipher.
					if word[j] > 'A' && word[j] < 'Z' {
						final += string(Capabet[(int(word[j]-'A')-key)%26])
					} else {
						final += string(Lowebet[(int(word[j]-'a')-key)%26])
					}
				}*/
				final += string(word[j])
			}
		}
		final += " "
	}

	fmt.Println(final)

}

func CaesarWrite(args []string, key int) {
	//fmt.Println("writing")
	final := ""
	for i := 0; i < len(args); i++ {
		word := args[i]
		for j := 0; j < len(word); j++ {
			if key > 0 {
				if word[j] > 'A' && word[j] < 'Z' {
					final += string(Capabet[((int(word[j])-'A')+key)%26])
				} else {
					final += string(Lowebet[((int(word[j])-'a')+key)%26])
				}
			} else {
				final += string(word[j])
			}
		}
		final += " "
	}
	fmt.Println(final)
}
