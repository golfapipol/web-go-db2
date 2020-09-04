package business

func GenderMapNumber(gender string) int {
	if gender == "M" {
		return 1
	}
	if gender == "F" {
		return 2
	}
	return 0
}
