package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Opérations et conversions")

	/** Utilisation des opérateurs arithmétiques **/
	var price, tax float32 = 275.00, 27.40
	var sum, difference, product, quotient float32

	sum = price + tax
	difference = price - tax
	product = price * tax
	quotient = price / tax

	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)

	var intVal = math.MaxFloat64
	var floatVal = math.MaxFloat64

	fmt.Println(intVal * 2)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))

	var posResult int = 3 % 2
	var negResult int = -3 % 2
	var absResult float64 = math.Abs(float64(negResult))

	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)

	var value float32 = 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)

	var greeting string = "Hello"
	var language string = "Go"
	var combinedString = greeting + ", " + language
	fmt.Println(combinedString)

	/** Utilisation des opérateurs de comparaison **/
	const first int = 100
	const second int = 200
	var equal bool = first == second
	var notEqual bool = first != second
	var lessThan bool = first < second
	var lessThanOrEqual bool = first <= second
	var greaterThan bool = first > second
	var greaterThanOrEqual bool = first >= second
	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)

	var max int
	if first > second {
		max = first
	} else {
		max = second
	}
	fmt.Println("MAX : ", max)

	var alpha int = 100
	var beta int = 100
	var pointer1 *int = &alpha
	var pointer2 *int = &beta
	var pointer3 *int = &alpha
	fmt.Println("Comparaison des pointeurs")
	fmt.Println(pointer1 == pointer3)
	fmt.Println(pointer1 == pointer2)
	fmt.Println("Comparaison des valeurs stockées par les pointeurs")
	fmt.Println(*pointer1 == *pointer3)
	fmt.Println(*pointer1 == *pointer2)

	var maxMph int = 50
	var passengerCapacity int = 4
	var airbags bool = true
	var familyCar bool = passengerCapacity > 2 && airbags
	var sportsCar = maxMph > 100 || passengerCapacity == 2
	var canCategorize = !familyCar && !sportsCar
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)

	/** Conversions de type explicites **/
	kayak := 275
	soccerBall := 19.50
	total1 := float64(kayak) + soccerBall
	fmt.Println(total1)
	total2 := kayak + int(soccerBall)
	fmt.Println("Total2 : ", total2)
	fmt.Println("Total2 : ", int8(total2))
	total3 := kayak + int(math.Round(soccerBall))
	fmt.Println("Total3 : ", total3)

	/** Formatage depuis une chaine de caractères **/
	var val1 string = "true"
	var val2 string = "false"
	var val3 string = "not true"
	var val4 string = "T"
	var bool1, b1Err = strconv.ParseBool(val1)
	var bool2, b2Err = strconv.ParseBool(val2)
	var bool3, b3Err = strconv.ParseBool(val3)
	fmt.Println("Bool 1 : ", bool1, b1Err)
	fmt.Println("Bool 2 : ", bool2, b2Err)
	fmt.Println("Bool 3 : ", bool3, b3Err)
	if bool4, b4Err := strconv.ParseBool(val4); b4Err == nil {
		fmt.Println("Parsed value : ", bool4)
	} else {
		fmt.Println("Cannot parse", val4)
	}

	var val5 string = "100"
	int1, int1err := strconv.ParseInt(val5, 0, 8)
	if int1err == nil {
		fmt.Println("Parsed value: ", int1)
	} else {
		fmt.Println("Cannot parse : ", val5)
	}

	var val6 string = "100"
	var int2 int64
	var int2err error
	int2, int2err = strconv.ParseInt(val6, 10, 0)
	if int2err == nil {
		var intResult int = int(int2)
		fmt.Println("Parsed value : ", intResult)
	} else {
		fmt.Println("Cannot parse : ", val6, int2err)
	}

	var val7 string = "100"
	int3, int3err := strconv.Atoi(val7)
	if int3err == nil {
		var intResult int = int3
		fmt.Println("Parsed value : ", intResult)
	} else {
		fmt.Println("Cannot parse : ", val7, int3err)
	}

	var val8 string = "48.95"
	float1, float1err := strconv.ParseFloat(val8, 64)
	if float1err == nil {
		fmt.Println("Parsed value : ", float1)
	} else {
		fmt.Println("Cannot parse : ", val8, float1err)
	}

	/** Formatage d'une valeur en chaine de caractères **/
	val9 := true
	val10 := false
	str1 := strconv.FormatBool(val9)
	str2 := strconv.FormatBool(val10)
	fmt.Println("Formatted value 1 : " + str1)
	fmt.Println("Formatted value 2 : " + str2)

	val11 := 275
	base10String1 := strconv.FormatInt(int64(val11), 10)
	base10String2 := strconv.Itoa(val11)
	base2String := strconv.FormatInt(int64(val11), 2)
	fmt.Println("Base 10 : " + base10String1)
	fmt.Println("Base 10 : " + base10String2)
	fmt.Println("Base 2 : " + base2String)

	val12 := 49.95
	Fstring := strconv.FormatFloat(val12, 'f', 2, 64)
	Estring := strconv.FormatFloat(val12, 'e', -1, 64)
	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)
}
