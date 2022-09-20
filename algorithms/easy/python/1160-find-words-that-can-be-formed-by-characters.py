from collections import Counter


class Solution:
    def countCharacters(self, words: list[str], chars: str) -> int:
        chars = Counter(chars)
        goodWordsLength = 0
        for word in words:
            word = Counter(word)
            wordFrequency = 0
            for char, frequency in word.items():
                if char not in chars or frequency > chars[char]:
                    break
                else:
                    wordFrequency += frequency
            else:
                goodWordsLength += wordFrequency
        return goodWordsLength