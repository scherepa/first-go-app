package practice

import (
	"fmt"
	"regexp"
	"strings"
)

func CountWordsPrint() {
	msg := `Say one two three. Say again.`
	result := make(map[string]int)
	re := regexp.MustCompile(`\p{P}`)
	words := strings.Fields(re.ReplaceAllString(strings.ToLower(msg), ""))

	for _, v := range words {
		result[v]++
	}
	fmt.Printf("The count of each word in '%s' is:\n%v\n", msg, result)
}
