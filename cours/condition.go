package cours

import "fmt"

func If() bool {
	a := 0

	if a == 0 {
		return true
	}

	return false
}

func IfElse() bool {
	a := 0

	if a == 0 {
		return true
	} else {
		return false
	}
}

func IfElseIfElse() bool {
	a := 0

	if a == 0 {
		return true
	} else if a != 0 {
		return false
	} else {
		return false
	}
}

func SwitchCase() {
	a := "Florent"

	switch a {
	case "Florent":
		fmt.Println("Hello", a)
	case "Toto":
		fmt.Println("Hello", a)
	default:
		fmt.Println("Hello", "world")
	}

}
