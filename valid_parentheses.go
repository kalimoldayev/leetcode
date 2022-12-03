package main

import "fmt"

func main() {
	// symbols := "{()}"
	// symbols := "()[]{}"
	// symbols := "(]"
	// symbols := "([)]"
	symbols := "([}}])"
	fmt.Println(isValid(symbols))
}

func isValid(s string) bool {

	if len(s)%2 == 0 { // если остаток равен к 0
		stack := []string{}           // создаю стек для хранения закрывающих скобок
		hashMap := map[string]string{ // мапка для сравнения во время итерации слайса
			"(": ")",
			"{": "}",
			"[": "]",
		}

		for i := range s {
			if _, ok := hashMap[string(s[i])]; ok { // если это открывающая скобка
				stack = append(stack, hashMap[string(s[i])]) // то в стек добавляю закрывающую скобку
			} else {
				if len(stack) > 0 && stack[len(stack)-1] == string(s[i]) { // если последний элемент в стеке и элемент итерации совпадают
					stack = stack[:len(stack)-1] // то удаляю последний элемент
				} else {
					return false // для этого ([)], в противном случае
				}
			}
		}

		if len(stack) == 0 { // если стек пустой
			return true
		}
	}

	return false
}


func isValid2(s string) bool { // неправильный вариант
	var (
		a, b, c int
	)

	for i := range s {
		switch {
		case string(s[i]) == "(":
			a++
		case string(s[i]) == "[":
			b++
		case string(s[i]) == "{":
			c++
		case string(s[i]) == ")":
			a--
		case string(s[i]) == "]":
			b--
		case string(s[i]) == "}":
			c--
		}
	}

	if a == 0 && b == 0 && c == 0 {
		return true
	}
	return false
}