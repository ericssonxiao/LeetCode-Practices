# Fast-Slow Pointer
# Pay more attention in Python, the // operator is used for integer division, which discards the fractional part.
# This is equivalent to / in Java when used with integers.
class Solution:
    def getNext(self, num):
        total = 0
        while num > 0:
            a = num % 10
            num = num // 10
            total += a * a
        return total

    def isHappy(self, n):
        slow = n
        fast = self.getNext(n)
        while fast != 1 and fast != slow:
            slow = self.getNext(slow)
            fast = self.getNext(self.getNext(fast))
        return fast == 1