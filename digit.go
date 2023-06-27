package digit

import "fmt"

var digitPronumciation = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func Pronunciation(d int) string {
	if d < 0 || d > 9 {
		return fmt.Sprintf("%d is not digit", d)
	}

	return digitPronumciation[d]
}
