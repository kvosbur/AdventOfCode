
class Board:
    number_chosen = -1

    def __init__(self, lines, start_index):
        self.board = Board.parse_board(lines, start_index)
        self.bingo = False

    @staticmethod
    def parse_board(lines, start_index):
        base = [lines[i] for i in range(start_index, start_index + 5)]
        board = []
        for row in base:
            board.append([int(y) for y in list(filter(lambda x: x != "", row.split(' ')))])
        return board

    def mark_on_board(self, target_num):
        for row in range(5):
            for column in range(5):
                if self.board[row][column] == target_num:
                    self.board[row][column] = Board.number_chosen

    def has_bingo(self):
        row_has_bingo = list(filter(lambda y: len(y) == 0, [list(filter(lambda x: x != Board.number_chosen, row)) for row in self.board]))
        bingo = len(row_has_bingo) > 0

        for column in range(5):
            column_has_bingo = True
            for row in range(5):
                if self.board[row][column] != Board.number_chosen:
                    column_has_bingo = False
            bingo = bingo or column_has_bingo
        self.bingo = bingo
        return bingo

    def get_score(self):
        return sum([sum(y) for y in [list(filter(lambda x: x != Board.number_chosen, row)) for row in self.board]])


def parse_all_boards(lines):
    index = 2
    boards = []

    while index < len(lines):
        boards.append(Board(lines, index))
        index += 6
    return boards


def main(lines):
    # do work here
    called_nums = [int(x) for x in lines[0].split(',')]
    boards = parse_all_boards(lines)

    numbers_index = 0
    boards_with_bingo = 0
    while boards_with_bingo < len(boards):
        for board in boards:
            if not board.bingo:
                board.mark_on_board(called_nums[numbers_index])
                if board.has_bingo():
                    boards_with_bingo += 1
                    if boards_with_bingo == len(boards):
                        score = board.get_score()
                        print("score is:", (score * called_nums[numbers_index]))

        numbers_index += 1



