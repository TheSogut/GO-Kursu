package maps

import "fmt"

func Demo1() {
	// key-value

	dict := make(map[string]string)
	dict["table"] = "Masa"
	dict["book"] = "Kitap"
	dict["pencil"] = "Kalem"

	fmt.Println(dict["book"])
	fmt.Println(len(dict))
	fmt.Println(dict)
	delete(dict, "book")
	fmt.Println(len(dict))
	fmt.Println(dict)

	value, isValid := dict["book"]
	fmt.Println(value)
	fmt.Println(isValid)

	dict2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(dict2)
}
