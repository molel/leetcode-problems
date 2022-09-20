class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        combinations = []
        self.parenthesisGenerator(n, 1, '(', combinations)
        return combinations

    def parenthesisGenerator(self, n, b, s, combinations):
        if len(s) < 2 * n:
            if b > 0:
                self.parenthesisGenerator(n, b - 1, s + ')', combinations)
            if 2 * n - len(s) > b:
                self.parenthesisGenerator(n, b + 1, s + '(', combinations)
        else:
            combinations.append(s)