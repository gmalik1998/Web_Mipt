package main
import (
	 "unicode"
	"strings"
)

func RemoveEven(num []int) []int {
	var nums []int;
	copy(nums, num)
	for i:=0; i<len(nums); i++ {
	    nums = append(nums[:i], nums[i+1:]...)
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
fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
}
*/
