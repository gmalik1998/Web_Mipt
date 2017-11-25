package main
import (
	 "unicode"
	"strings"
)

func RemoveEven(num []int) []int {
	nums := make([]int, len(num)/2, (cap(num)+1) * 2)
	for i:=0; i<len(num); i++ {
	    if i % 2 == 1 {
	    nums[(i- 1)/2] = num[i]
	}
	}
	return nums
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
input := []int{0, 3, 2, 5, 7, 8}
result := RemoveEven(input)
fmt.Println(result) // Должно напечататься [3 5]
}
*/

