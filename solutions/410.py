from typing import List


class Solution:
    @staticmethod
    def check(nums, m, possible):
        i, cnt = 0, 0
        while i < len(nums) and cnt < m:
            ss, j = 0, i
            while j < len(nums):
                if nums[j] > possible:
                    return False
                if ss + nums[j] <= possible:
                    ss += nums[j]
                    j += 1
                else:
                    break
            i = j
            cnt += 1
        return cnt <= m and i >= len(nums)

    def splitArray(self, nums: List[int], m: int) -> int:
        low, high = min(nums), sum(nums)
        while low <= high:
            mid = (low + high) >> 1
            is_ok = self.check(nums, m, mid)
            if is_ok:
                high = mid - 1
            else:
                low = mid + 1
        return low
