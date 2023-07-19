package main

func main() {
	// if, else if, else (if only is ok.)
	age := 22
	if age >= 20 {
		println("adult")
	} else if age == 0 {
		println("baby")
	} else {
		println("child")
	}

	// simple way to create val and if together
	if age := 0; age >= 20 {
		println("adult", age)
	} else if age == 0 {
		println("baby", age)
	} else {
		println("child", age)
	}

}
