package main

import "fmt"

func main() {
    words := []string{"flower","flow","flht"}
    // words := []string{"car","cir"}

    fmt.Println(longestCommonPrefix(words))
}

func longestCommonPrefix(words []string) string {
    hashMap := make(map[string]int, len(words[0]))
    result := ""

    for i := range words[0] { // итерация по первому слову
        lenWords := len(words) - 1
        res := 0 // счетчик

        for j:=0; j<=lenWords; j++ { // итерация по всем словам
            if len(words[j]) > i { // если есть буква
                if string(words[j][i]) == string(words[0][i]) { // если есть буква
                    res++ // обновляю счетчик
                }
                hashMap[string(words[j][i])] = res // обновляю счетчик по ключу, ключ равен к букве
            }
        }

        if val, ok := hashMap[string(words[0][i])]; ok {
            if val == len(words) {
                result = result + string(words[0][i]) // если равен
            } else {
                break // если след буква нету в остальных словах, выхожу
            }
        }

    }

    return result
}