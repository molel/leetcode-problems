class Solution:
    def bestHand(self, ranks: List[int], suits: List[str]) -> str:
        best_hand = 'High Card'
        ranks_dict = {rank: 0 for rank in set(ranks)}
        suits_dict = {suit: 0 for suit in set(suits)}
        for rank in ranks:
            ranks_dict[rank] += 1
        for suit in suits:
            suits_dict[suit] += 1
        max_rank = max(ranks_dict.values())
        if max_rank == 2:
            best_hand = 'Pair'
        elif max_rank >= 3:
            best_hand = 'Three of a Kind'
        if max(suits_dict.values()) == 5:
            best_hand = 'Flush'
        return best_hand
        