package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func strining_strconv() {
	char1 := string(116) // converts 100 to asccii character
	fmt.Println(char1)

	// Integer to Ascii
	word1 := strconv.Itoa(100)
	fmt.Println("this is 100 as string", word1)
	int1, err1 := strconv.Atoi(word1)
	if err1 != nil {
		fmt.Println(word1, "converted word(100) to as int", int1)
	}
	b1, _ := strconv.ParseBool("True")
	fmt.Println(b1)
}
func string_search_examples() {
	s := " this is very lenghty line that doesnot contains all characters from a to z"
	fmt.Println(strings.Contains(s, "very"))
	fmt.Println(strings.ContainsAny(s, "test"))
	fmt.Println(strings.Index("Raman", "a"))
	fmt.Println(strings.IndexAny("Raman", "bcdm"))
	fmt.Println(strings.HasPrefix("TVRaman", "TV"))
	fmt.Println(strings.HasPrefix("Raman", "man"))
	fmt.Println(strings.Count("Raman", "m"))
}

func string_builder() {
	var sb strings.Builder

	sb.WriteString("First Line \n")
	fmt.Println(sb.String(), "capacity is ", sb.Cap(), " size is ", sb.Len())
	sb.Grow(1024)
	fmt.Println("After Growing Capacity,", sb.Cap())
	sb.Reset()
	sb.Len()
	fmt.Println("After Reset ", sb.String(), "capacity is ", sb.Cap(), " size is ", sb.Len())

}
func string_map_functions() {
	changeBy := 2
	s := "rAmaN1"
	myTransalator := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			value := int('A') + int(r) - int('A') + changeBy
			return rune(value)

		case r >= 'a' && r <= 'z':
			return r

		}
		return r + 3
	}
	fmt.Println(strings.Map(myTransalator, s))
}
func string_manipulation_examples() {
	s := " Nice to be here"
	words := strings.Split(s, " ")
	for x, word := range words {
		fmt.Println(x, word)
	}
	for x := 0; x < len(words); x++ {
		fmt.Println(x, words[x])
	}

	fmt.Println(strings.Fields(s))
	newStr := strings.Join(words, "-")
	fmt.Println(newStr)
	mySplitFunc := func(c rune) bool {
		return unicode.IsPunct(c)
	}
	newWords := strings.FieldsFunc(newStr, mySplitFunc)
	fmt.Println(newWords)
	replacer := strings.NewReplacer("a", "b", "m", "M")
	newWriter := replacer.Replace("raman")
	fmt.Println(newWriter)
}

func string_basic_examples() {
	myString := "My go Programming Learning"
	fmt.Println("Lenght of myString", len(myString))
	fmt.Println("Ones" == "111")
	fmt.Println("Raman" == "Raman")
	my := "raman"
	fmt.Println(strings.Compare("raman", my))
	fmt.Println(strings.EqualFold("Raman", my))

}
func main() {
	strining_strconv()
}
