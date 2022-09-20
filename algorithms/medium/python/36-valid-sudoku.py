class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        rows = {row: [] for row in range(9)}
        columns = {column: [] for column in range(9)}
        squares = {square: [] for square in range(9)}
        for i in range(9):
            for j in range(9):
                if board[i][j] != '.':
                    if (board[i][j] in rows[i] or
                            board[i][j] in columns[j] or
                            board[i][j] in squares[i // 3 * 3 + j // 3 % 3]):
                        return False
                    else:
                        rows[i].append(board[i][j])
                        columns[j].append(board[i][j])
                        squares[i // 3 * 3 + j // 3 % 3].append(board[i][j])
        return True