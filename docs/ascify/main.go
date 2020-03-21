package Ascify

import (
	"bufio"
	"os"
	"strings"
)

type Error500 struct {
	err500 bool
}

func AsciiArt(text, filename string) (string, bool) {
	filename += ".txt"
	err := &Error500{err500: false}
	cont := ""
	res := ""
	args := strings.Split(text, "\r\n")
	for _, word := range args {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				cont += GetLine(2+int(letter-' ')*9+i, filename, err)
			}
			cont += "\n"
			res += cont
			cont = ""
		}
	}
	return res, err.err500
}

func GetLine(num int, filename string, err *Error500) string {
	str := ""
	f, e := os.Open("./docs/ascify/" + filename)
	if e != nil {
		err.err500 = true
		return ""
	}
	defer f.Close()

	f.Seek(0, 0)
	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		str, _ = content.ReadString('\n')
	}
	str = strings.TrimSuffix(str, "\n")

	return str
}
