package cours

import "fmt"

func boucleWhile() {
	i := 0
	for {
		if i > 10 {
			break
		}

		// algo

		i++
	}
}

func boucleDoWhile() {
	i := 0
	for {

		// algo

		i++
		if i > 10 {
			break
		}
	}
}

func boucleFor() {

	for i := 0; i < 10; i++ {
		// algo
	}

}

func boucleForEach() {

	a := []string{"Toto", "Tata", "Titi"}

	for index := range a {
		fmt.Println(index)
	}

	for index, value := range a {
		fmt.Println(index, value)
	}

	for _, value := range a {
		fmt.Println(value)
	}

}
