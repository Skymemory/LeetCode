from typing import List


class Solution:
    def dailyTemperatures(self, T: List[int]) -> List[int]:
        size = len(T)
        outcome = [0] * size
        stack = []
        for i in reversed(range(size)):
            while stack and T[stack[-1]] <= T[i]:
                stack.pop()
            if stack:
                outcome[i] = stack[-1] - i
            stack.append(i)
        return outcome
