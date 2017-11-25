package main
import (
	"fmt"
	 "unicode"
	"strings"
)

func RemoveEven(num []int) []int {
	for i:=0; i<len(num); i++ {
	    num = append(num[:i], num[i+1:]...)
	}
	return num
}

func PowerGenerator(step int) func() int{
	sum := 1
	return func() int {
	    sum *= step
	    return  sum
	}
}

func DifferentWordsCount(a string) int {
	answer:= make(map[string]int);
	temp:= ""
	for _, word:= range a {
		wor := rune(word)
		if unicode.IsLetter(wor) {
			temp += string(word)
		} else if len(temp) > 0 {
			answer[strings.ToUpper(temp)] = 1
			temp = ""
		}
	}

	if len(temp) > 0 {
		answer[temp] = 1
	}

	return len(answer)
}
/*
func main() {
fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
}
*/
