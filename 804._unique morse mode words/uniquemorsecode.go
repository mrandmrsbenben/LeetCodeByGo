package main

import "fmt"

func main() {
    words := []string{"gin", "zen", "gig", "msg"}
    fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) int {
    morsevalue := []string{".-","-...","-.-.","-..",".","..-.",
        "--.","....","..",".---","-.-",".-..","--","-.",
        "---",".--.","--.-",".-.","...","-","..-","...-",
        ".--","-..-","-.--","--.."}
    morsekey := []string{"a","b","c","d","e","f","g","h","i",
        "j","k","l","m","n","o","p","q","r","s","t","u",
        "v","w","x","y","z"}
    morseT := make(map[string]string)
    for i,key := range morsekey {
        morseT[key] = morsevalue[i]
    }

    codeM := make(map[string]string)
    var code string
    for _,w := range words{
        code = ""
        for i := 0; i < len(w); i++ {
            code = code + morseT[string(w[i])]
        }
        codeM[code] = w
    }
    return len(codeM)
}
