package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var txt string
	var nw_txt string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pt := scanner.Text()

	scanner.Scan()
	txt = scanner.Text()

	fmt.Println("Select 2 options")
	fmt.Println(" - 1 encrypt with ROT 13")
	fmt.Println(" - 2 decrypt with ROT 13")
	fmt.Print("\n")
	nw_txt = ""
	if pt == "1" { //แปลง
		for i := 0; i < len(txt); i++ {
			if 65 <= txt[i] && txt[i] <= 90 {
				if txt[i]+13 > 90 {
					nw_txt += string(txt[i] - 13)
				} else {
					nw_txt += string(txt[i] + 13)
				}
			} else if 97 <= txt[i] && txt[i] <= 122 {
				if txt[i]+13 > 122 {
					nw_txt += string(txt[i] - 13)
				} else {
					nw_txt += string(txt[i] + 13)
				}
			} else {
				nw_txt += string(txt[i])
			}
		}
		fmt.Printf(`Choose option: Enter text: Ciphertext is "%v"`, string(nw_txt))
	} else if pt == "2" { //ถอด
		for i := 0; i < len(txt); i++ {
			if 65 <= txt[i] && txt[i] <= 90 {
				if txt[i]-13 < 65 {
					nw_txt += string(txt[i] + 13)
				} else {
					nw_txt += string(txt[i] - 13)
				}
			} else if 97 <= txt[i] && txt[i] <= 122 {
				if txt[i]-13 < 97 {
					nw_txt += string(txt[i] + 13)
				} else {
					nw_txt += string(txt[i] - 13)
				}
			} else {
				nw_txt += string(txt[i])
			}
		}
		fmt.Printf(`Choose option: Enter text: Plaintext is "%v"`, string(nw_txt))
	}

}
