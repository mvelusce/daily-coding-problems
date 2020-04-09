package main

import (
	"fmt"
)

type RegularExpression struct {
	transitions []*Transition
	final       bool
}

type Transition struct {
	character   string
	destination *RegularExpression
}

func (r *RegularExpression) isEnd() bool {
	return r.final
}

func (r *Transition) isMatch(character string) bool {
	if r.character == "." {
		return true
	}
	return r.character == character
}

func parseRegularExpression(exp string) *RegularExpression {

	root := &RegularExpression{}
	last := root
	for i, c := range exp {
		if string(c) == "*" {
			continue
		} else {
			if i < len(exp)-1 && string(exp[i+1]) == "*" {
				t := &Transition{character: string(c), destination: last}
				last.transitions = append(last.transitions, t)
			} else {
				s := &RegularExpression{}
				t := &Transition{character: string(c), destination: s}
				last.transitions = append(last.transitions, t)
				last = s
			}
		}
	}
	last.final = true
	return root
}

func isMatch(str string, r *RegularExpression) bool {

	//fmt.Println("str", str)
	//fmt.Println("f", r.final)
	if str == "" && r.isEnd() {
		//fmt.Println("end")
		return true
	}
	if str == "" {
		//fmt.Println("empty")
		return false
	}
	res := false
	for _, t := range r.transitions {
		//fmt.Println("tb", t.character)
		if t.isMatch(string(str[0])) {
			//fmt.Println("t", t.character)
			res = res || isMatch(str[1:], t.destination)
		}
	}
	return res
}

func main() {

	r := parseRegularExpression("asd")
	fmt.Println("asd", isMatch("asd", r))
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
	fmt.Println("aa", isMatch("aa", r))

	r = parseRegularExpression("c*a*b*")
	fmt.Println("aab", isMatch("aab", r))

	r = parseRegularExpression("mis*is*p*.")
	fmt.Println("mississippi", isMatch("mississippi", r))

	r = parseRegularExpression("ab*a*c*a")
	fmt.Println("aaba", isMatch("aaba", r))
}
