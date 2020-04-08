package main

import (
	"fmt"
)

type RegularExpression struct {
	character   string
	transitions []*RegularExpression
	isOptional  bool
}

func (r *RegularExpression) isMatch(character string) bool {
	if r.character == "." {
		return true
	}
	return r.character == character || r.isOptional
}

func (r *RegularExpression) isEnd() bool {
	if len(r.transitions) == 0 {
		return true
	}
	return len(r.transitions) < 2 && r.isOptional
}

func parseRegularExpression(exp string) *RegularExpression {

	root := RegularExpression{character: "root"}
	last := &root
	for _, c := range exp {
		if string(c) == "*" {
			last.transitions = append(last.transitions, last)
			last.isOptional = true
		} else {
			r := &RegularExpression{character: string(c)}
			last.transitions = append(last.transitions, r)
			last = r
		}
	}
	/*a := &root
	for i := 0; i<3; i++ {
		if len(a.transitions) > 0 {
			fmt.Println(len(a.transitions))
			fmt.Println(a.transitions[0].character)
			a = a.transitions[0]
		}
	}*/
	return root.transitions[0]
}

func isMatch(str string, r *RegularExpression) bool {

	fmt.Println("str", str)
	fmt.Println("c", r.character)
	fmt.Println()
	if str == "" {
		fmt.Println("empty str")
		return r.isEnd()
	}
	if r.isMatch(string(str[0])) && r.isEnd() {
		fmt.Println("equal")
		return true
	}
	if !r.isMatch(string(str[0])) && r.isEnd() {
		fmt.Println("diff")
		return false
	}

	res := false
	for _, t := range r.transitions {
		res = res || isMatch(str[1:], t)
	}
	return res
}

func main() {

	r := parseRegularExpression("asd")
	/*fmt.Println("asd", isMatch("asd", r))
	fmt.Println("as", isMatch("as", r))
	fmt.Println("asdd", isMatch("asdd", r))

	r = parseRegularExpression("aa.*asd*qwe")
	fmt.Println("aarrrasdddqwe", isMatch("aarrrasdddqwe", r))

	r = parseRegularExpression("*")
	fmt.Println("", isMatch("", r))

	r = parseRegularExpression(".*")
	fmt.Println("", isMatch("", r))
	fmt.Println("tyutyu", isMatch("tyutyu", r))

	r = parseRegularExpression("a*")
	fmt.Println("aa", isMatch("aa", r))*/

	r = parseRegularExpression("c*a*b*")
	fmt.Println("aab", isMatch("aab", r))

	r = parseRegularExpression("mis*is*p*.")
	fmt.Println("mississippi", isMatch("mississippi", r))
}
