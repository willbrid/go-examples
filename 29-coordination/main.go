package main

func doSum(count int, val *int) {
	for i := 0; i < count; i++ {
		*val++
	}
}

func main() {
	counter := 0
	doSum(5000, &counter)
	Printfln("Total : %v", counter)
}
