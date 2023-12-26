package main

// getNext and isHappy, these two function are just standalone functions.
func getNext(num int) int {
	total := 0
	for num > 0 {
		a := num % 10
		num = num / 10
		total += a * a
	}
	return total
}

func isHappy(n int) bool {
	slow := n
	fast := getNext(n)
	for fast != 1 && fast != slow {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}