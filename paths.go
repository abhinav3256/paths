package main

import "fmt"

var arr = [][]string{{"Lima", "Sao Paulo"}, {"London", "New York"}, {"New York", "Lima"}}

func main() {
	m := make(map[string]string)
	for i := 0; i < len(arr); i++ {

		key_ := arr[i][0]

		m[key_] = arr[i][1]

	}

	for key, value := range m {

		temp := value

		for key2, _ := range m {

			if temp == key2 {
				delete(m, key)
				break
			}

		}

	}
	for key3, _ := range m {
		fmt.Println(m[key3])
	}

}
