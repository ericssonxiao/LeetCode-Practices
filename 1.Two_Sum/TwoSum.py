

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        tempMap = {} # val:index
        
        for i, item in enumerate(nums):
            diff = target - item
            if diff in tempMap:
                return [tempMap[diff],i]
            tempMap[item]=i
        return 