// Package leap contains a function to identify leap year
package leap

// IsLeapYear returns a boolean indicating if its true/false
func IsLeapYear(time int) bool {
	return time%400 == 0 || time%100 != 0 && time%4 == 0
}
