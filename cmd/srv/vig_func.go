package srv

import "fmt"

func KeygenVig(args []string, key string) string {
	GenKey := ""
	Keyint := 0
	for i := 0; i < len(args); i++ {
		temp := args[i]
		for j := 0; j < len(temp); j++ {
			if Keyint == len(key) {
				Keyint = 0
			}
			GenKey += string(key[Keyint])
			Keyint++
		}
		GenKey += " "
	}
	GenKey = GenKey[:len(GenKey)-1]
	return GenKey
}

func VigRead(args []string, key string) {

	key = KeygenVig(args, key)
	decrypt := ""
	keyint := 0
	for i := 0; i < len(args); i++ {
		temp := args[i]
		for j := 0; j < len(temp); j++ {
			var fix byte
			if temp[j] >= 'A' && temp[j] <= 'Z' {
				fix = 'A'
			} else {
				fix = 'a'
			}
			decrypt += string((((temp[j] - key[keyint]) + 26) % 26) + fix)
			keyint++
		}
		decrypt += " "
	}
	fmt.Println(decrypt)
}

func VigWrite(args []string, key string) {
	key = KeygenVig(args, key)
	encrypt := ""
	keyint := 0
	for i := 0; i < len(args); i++ {
		temp := args[i]
		for j := 0; j < len(temp); j++ {
			var fix byte
			if temp[j] >= 'A' && temp[j] <= 'Z' {
				fix = 'A'
			} else {
				fix = 'a'
			}
			encrypt += string(((temp[j] + key[keyint]) % 26) + fix)
			keyint++
		}
		encrypt += " "
	}
	fmt.Println(encrypt)
}
