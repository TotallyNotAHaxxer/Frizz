package Frizz_Helper

func Appender(x [][]string, y []string) []string {
	for i := range x {
		y = append(y, x[i][0])
	}
	return y
}
