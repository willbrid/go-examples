package main

import "fmt"

func main() {
	fmt.Println("Hello, Collections")

	fmt.Println("Travailler avec des tableaux")
	var names [3]string
	fmt.Println(names)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)

	var names1 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names1)

	names2 := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names2)

	names3 := [5]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names3)

	var coords [3][3]int
	coords[0][1] = 7
	coords[1][2] = 10
	fmt.Println(coords)

	var names4 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var names5 [3]string = names4

	names[0] = "Canoe"
	fmt.Println(names)
	fmt.Println(names5)

	var names6 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var pointerNames6 *[3]string = &names6
	names6[0] = "Canoe"
	fmt.Println("Names : ", names6)
	fmt.Println("Pointer : ", *pointerNames6)

	var names7 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var pointerNames7Element1 *string = &names7[1]
	fmt.Println("Pointer to Element 1 Before : ", *pointerNames7Element1)
	names7[1] = "Canoe"
	fmt.Println("Names : ", names7)
	fmt.Println("Pointer to Element 1 After : ", *pointerNames7Element1)

	var names8 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var names9 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var isSame bool = names8 == names9
	fmt.Println("Comparaison : ", isSame)

	var names10 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	for index, value := range names10 {
		fmt.Println("Index : ", index, " - Value : ", value)
	}

	var names11 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	for _, value := range names11 {
		fmt.Println("Value : ", value)
	}

	fmt.Println("Travailler avec Slices : Tableau dont on ne connait pas sa longueur en avance ou dont sa longueur est variable.")
	var names12 []string = make([]string, 3)
	names12[0] = "Kayak"
	names12[1] = "Lifejacket"
	names12[2] = "Paddle"
	fmt.Println(names12)

	names13 := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names13)

	names14 := []string{"Kayak", "Lifejacket", "Paddle"}
	names14 = append(names14, "Hat", "Gloves")
	fmt.Println(names14)

	var names15 []string = []string{"Kayak", "Lifejacket", "Paddle"}
	var appendedNames15 []string = append(names15, "Hat", "Gloves")
	names15[0] = "Canoe"
	fmt.Println(names15)
	fmt.Println(appendedNames15)

	var names16 []string = make([]string, 3, 7)
	names16[0] = "Kayak"
	names16[1] = "Lifejacket"
	names16[2] = "Paddle"
	fmt.Println("Tableau : ", names16)
	fmt.Println("Longueur : ", len(names16))
	fmt.Println("Capacité : ", cap(names16))

	var names17 []string = make([]string, 3, 7)
	names17[0] = "Kayak"
	names17[1] = "Lifejacket"
	names17[2] = "Paddle"
	moreName := []string{"Hat Gloves"}
	appendedNames17 := append(names17, moreName...)
	fmt.Println("AppendedNames : ", appendedNames17)

	var products [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))
	someNames = append(someNames, "Gloves")
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))
	someNames = append(someNames, "Boots")
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))

	var products1 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames1 := products1[1:3:3]
	allNames1 := products1[:]
	someNames1 = append(someNames1, "Gloves")
	fmt.Println("SomeNames : ", someNames1)
	fmt.Println("SomeNames len : ", len(someNames1), " cap : ", cap(someNames1))
	fmt.Println("AllNames : ", allNames1)
	fmt.Println("AllNames len : ", len(allNames1), " cap : ", cap(allNames1))
}
