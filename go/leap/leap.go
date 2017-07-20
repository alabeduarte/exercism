package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) ||
		(year%100 == 0 && year%400 == 0)
}
