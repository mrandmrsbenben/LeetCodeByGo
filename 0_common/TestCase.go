package common

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func MakeTestCases(f string) {
	text, err := clipboard.ReadAll()
	output := ""
	if err != nil {
		panic(err)
	}
	str := strings.Split(text, "\n")
	cnt := 1
	var p string
	var names []string
	fmt.Println("Lines:", len(str))
	readingInput := false
	names = make([]string, 0)
	for i, s := range str {
		//Output
		if strings.HasPrefix(s, "Output:") {
			readingInput = false
			// fmt.Printf("fmt.Printf(\"Output: %%v\\n\", %s(%s))\n", f, names[len(names)-1])
			// output = output + fmt.Sprintf("fmt.Printf(\"Output: %%v\\n\", %s(%s))\n", f, names[len(names)-1])
			if strings.HasSuffix(s, "]") || len(s[len("Output: "):]) > 0 {
				//Read Current Line
				s = s[len("Output: "):]
			} else {
				//Read Next Line
				s = str[i+1]
			}
			//Output Array
			if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
				s = strings.Replace(s, "[", "{", -1)
				s = strings.Replace(s, "]", "}", -1)
			}
			s = strings.Replace(s, "\"", "", -1)
			fmt.Printf("fmt.Printf(\"Expect%d: %s\\n\")\n", cnt, s)
			output = output + fmt.Sprintf("fmt.Printf(\"Expect%d: %s\\n\")\n", cnt, s)
			output = output + outputString(names, f, cnt)
			names = make([]string, 0)
			cnt = cnt + 1
		}
		//Read Inputs
		if strings.HasPrefix(s, "Input:") || readingInput {
			readingInput = true
			if strings.TrimSpace(s) == "Input:" {
				continue
			} else {
				s = strings.TrimSpace(strings.TrimPrefix(s, "Input:"))
			}
			if strings.Index(s, "[") != -1 && strings.Index(s, "]") != -1 {
				p, names = arrayParam(s, cnt, names)
			} else if strings.Index(s, "\"") != -1 {
				p, names = stringParam(s, cnt, names)
			} else {
				p, names = intParam(s, cnt, names)
			}
			output = output + fmt.Sprintln(p)
		}
	}
	clipboard.WriteAll(output)
}

func arrayParam(s string, cnt int, names []string) (string, []string) {
	s = strings.Replace(s, "[", "{", -1)
	s = strings.Replace(s, "]", "}", -1)
	var p string
	if strings.Index(s, "=") == -1 {
		p = "input"
	} else {
		p = s[0:strings.Index(s, "=")]
	}
	p = strings.TrimSpace(p) + strconv.Itoa(cnt)
	names = append(names, p)
	rcnt := strings.Count(s[0:strings.Index(s, "}")], "{")
	if strings.Index(s, "\"") != -1 {
		p = p + ":=" + strings.Repeat("[]", rcnt) + "string"
	} else {
		p = p + ":=" + strings.Repeat("[]", rcnt) + "int"
	}
	p = p + s[strings.Index(s, "{"):]
	fmt.Println(p)
	return p, names
}

func stringParam(s string, cnt int, names []string) (string, []string) {
	var p string
	if strings.Index(s, "=") != -1 {
		p = strings.Replace(s, "=", ":=", 1)
		name := strings.TrimSpace(s[0:strings.Index(s, "=")]) + strconv.Itoa(cnt)
		p = strings.Replace(p, name[0:len(name)-1], name, 1)
		names = append(names, name)
	} else {
		name := "input" + strconv.Itoa(cnt)
		p = name + " := " + s
		names = append(names, name)
	}
	fmt.Println(p)
	return p, names
}

func intParam(s string, cnt int, names []string) (string, []string) {
	var p string
	if strings.Index(s, "=") != -1 {
		p = strings.Replace(s, "=", ":=", 1)
		name := strings.TrimSpace(s[0:strings.Index(s, "=")]) + strconv.Itoa(cnt)
		p = strings.Replace(p, name[0:len(name)-1], name, 1)
		names = append(names, name)
	} else {
		p = "input" + strconv.Itoa(cnt) + ":= " + s
		names = append(names, "input"+strconv.Itoa(cnt))
	}
	fmt.Println(p)
	return p, names
}

func outputString(names []string, f string, cnt int) string {
	str := ""
	for _, v := range names {
		str = str + v + ","
	}
	str = str[0 : len(str)-1]
	fmt.Printf("fmt.Printf(\"Output%d: %%v\\n\", %s(%s))\n", cnt, f, str)
	return fmt.Sprintf("fmt.Printf(\"Output%d: %%v\\n\", %s(%s))\n", cnt, f, str)
}
