package main

// In Go language, methods are defined on types.
// Here, we can define a Solution type and attached getNext and isHappy methods to it.
type Solution struct{}

func (s Solution) getNext(num int) int {
	total := 0
	for num > 0 {
		a := num % 10
		num = num / 10
		total += a * a
	}
	return total
}

func (s Solution) isHappy(n int) bool {
	slow := n
	fast := s.getNext(n)
	for fast != 1 && fast != slow {
		slow = s.getNext(slow)
		fast = s.getNext(s.getNext(fast))
	}
	return fast == 1
}