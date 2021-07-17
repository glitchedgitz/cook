package main

import "strings"

func splitMethods(p string) []string {
	chars := strings.Split(p, "")
	s := []string{}
	tmp := ""
	insidebrackets := false
	for _, c := range chars {

		if c == "." {
			if !insidebrackets {
				s = append(s, tmp)
				tmp = ""
				continue
			}
		}
		if c == "[" {
			insidebrackets = true
		}
		if c == "]" {
			insidebrackets = false
		}
		tmp += c
	}
	s = append(s, tmp)
	return s
}

func splitValues(p string) []string {
	chars := strings.Split(p, "")
	s := []string{}
	tmp := ""
	insideraw := false

	for _, c := range chars {

		if c == "," {
			if !insideraw {
				s = append(s, tmp)
				tmp = ""
				continue
			}
		}
		if c == "`" {
			if insideraw {
				insideraw = false
			} else {
				insideraw = true
			}
		}

		tmp += c
	}
	s = append(s, tmp)
	return s
}
