package tutorial1

//CheckIfPrime
func CheckIfPrime(num int) bool {
	result := true
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			result = false
			break
		}
	}
	return result
}
