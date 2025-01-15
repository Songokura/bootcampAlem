package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func PrintString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func Valid(s string) bool {
	return s == "cur" || s == "maltipoo" || s == "simple" || s == "tazy"
}

func Split(s string, sep string) []string {
	var t string
	var res []string

	if s == "" {
		return res
	}

	for _, c := range s {
		t += string(c)
		if c == '\n' {
			res = append(res, t[:len(t)-1])
			t = ""
		}
	}
	if t != "" || s[len(s)-1] == '\n' {
		res = append(res, t)
	}

	return res
}

func LineCount(s string) int {
	t := Split(s, "\n")
	if t[len(t)-1] == "" {
		return len(t) - 1
	}
	return len(t)
}

func Max(s []string) int {
	mx := 0

	for _, v := range s {
		if mx < len(v) {
			mx = len(v)
		}
	}

	return mx
}

func PrintText(s string) {
	lines := Split(s, "\n")
	length := Max(lines)

	for i := range lines {
		for len(lines[i]) < length {
			lines[i] += " "
		}
	}

	ap.PutRune(' ')
	PrintString("______")
	for i := 6; i < length+2; i++ {
		ap.PutRune('_')
	}
	ap.PutRune('\n')

	if len(lines) == 1 {
		PrintString("< " + lines[0] + " >\n")
	} else if len(lines) == 0 {
		PrintString("<  >\n")
	} else {
		PrintString("/ " + lines[0] + " \\\n")
		for i := 1; i+1 < len(lines); i++ {
			PrintString("| " + lines[i] + " |\n")
		}
		PrintString("\\ " + lines[len(lines)-1] + " /\n")
	}

	ap.PutRune(' ')
	PrintString("------")
	for i := 6; i < length+2; i++ {
		ap.PutRune('-')
	}
	ap.PutRune('\n')
}

func PrintSimple() {
	PrintString(`  \
^..^      /
/_/\_____/
   /\   /\
  /  \ /  \`)
}

func PrintMaltipoo() {
	PrintString(`  \
   \ __    __
   o-''))_____\\
   "--__/ * * * )
   c_c__/-c____/`)
}

func PrintTazy() {
	PrintString(`  \                __
   \___________   /  \
               \ / ..|\
                (_\  |_)
                /  \@'
               /     \
           _  /  ` + "`" + `   |
          \\ /  \  | _\
           \   /_ || \\_
            \____)|_) \_)`)
}

func PrintCur() {
	PrintString(`   \
    \ D\___/\
     \ (0_o)
        (V)
----oOo--U--oOo------
_______|_______|_____`)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) == 2 || len(args) > 3 {
		PrintString("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
	} else if len(args) == 1 {
		if args[0] == "-b" {
			PrintString("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
		} else {
			PrintText(args[0])
			PrintSimple()
			ap.PutRune('\n')
		}
	} else if len(args) == 3 {
		if args[0] != "-b" || (args[0] == "-b" && !Valid(args[1])) {
			PrintString("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
		} else {
			PrintText(args[2])
			if args[1] == "cur" {
				PrintCur()
			} else if args[1] == "maltipoo" {
				PrintMaltipoo()
			} else if args[1] == "simple" {
				PrintSimple()
			} else {
				PrintTazy()
			}
			ap.PutRune('\n')
		}
	}
}
