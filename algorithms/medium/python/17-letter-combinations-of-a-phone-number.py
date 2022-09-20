class Solution:
    def letterCombinations(self, digits: str) -> list[str]:
        if not digits:
            return []
        letters = ["abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]
        combinations = list(letters[int(digits[0])-2])
        for digit in digits[1:]:
            combinations = self.multiply(combinations, letters[int(digit) - 2])
        return combinations

    def multiply(self, x, y):
        return [i + j for j in y for i in x]