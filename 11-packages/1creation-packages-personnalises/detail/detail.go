package detail

import "strconv"

func DisplayInfos(name string, amount float64) string {
	return "[ name : " + name + " - $" + strconv.FormatFloat(amount, 'f', 2, 64) + " ]"
}
