class Solution:
    def isValid(self, s: str) -> bool:
        openBrackets = ['{', '[', '(']
        stack = []
        for c in s:
            if c in openBrackets:
                stack.append(c)
            elif stack:
                if c == ']' and stack[-1] != '[':
                    return False
                if c == '}' and stack[-1] != '{':
                    return False
                if c == ')' and stack[-1] != '(':
                    return False
                stack.pop()
            else:
                return False
        return len(stack) == 0
