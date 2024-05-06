package bf

import (
	"flag"
	"fmt"
)

func main() {
	var _file = flag.String("file", "", "A file with Brainfuck instructions to interpret")
	var _txt = flag.String("text", "", "A string of Brainfuck instructions to interpret")
	flag.Parse()
	file := *_file
	txt := *_txt

	if file == "" && txt == "" {
		fmt.Println("No inputs provided!")
	} else if txt == "" {
		out, err := ParseFile(file)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
	} else if file == "" {
		out, err := ParseString(txt)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
	} else {
		fmt.Println("This program can only interpret one file or string of text at a time")
		fmt.Println("Using both flags is ambiguous, please provide a file OR text to interpret")
	}
}
