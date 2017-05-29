package leap

import "math"

const testVersion = 3

func IsDivisible(number int, divisor int) bool {
  return math.Mod(float64(number), float64(divisor)) == 0
}

func IsLeapYear(year int) bool {
  if IsDivisible(year, 4) && !IsDivisible(year, 100) ||
    IsDivisible(year, 100) && IsDivisible(year, 400) {
    return true
  }

  return false
}
