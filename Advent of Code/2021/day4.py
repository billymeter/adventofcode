#!/usr/bin/env python3

from functools import reduce
import operator


with open('input4.txt') as f:
    data = f.read()


class Board:
    def __init__(self, lines):
        self.board = [[*map(int, x.split('\n')[0].split())]
                      for x in lines.split('\n')]
        if self.board[-1] == []:
            self.board = self.board[:-1]
        self.numbers = []
        self.i = 0
        pass

    def __str__(self):
        return str(self.board)

    def __repr__(self):
        return str(self.board)

    def __iter__(self):
        return self

    def __next__(self):
        if self.i < len(self.board):
            self.i += 1
            return self.board[self.i - 1]
        else:
            raise StopIteration

    def check_for_win(self):
        # check for horizontal lines first
        for line in self.board:
            wins = 0
            for num in line:
                if num in self.numbers:
                    wins += 1
                if wins == 5:
                    return line

        # transpose the board to easily check
        # the vertical lines
        try:
            self.board = [[row[i] for row in self.board] for i in range(5)]
        except:
            print('problem boi:')
            print(self.board)

        for line in self.board:
            wins = 0
            for num in line:
                if num in self.numbers:
                    wins += 1
                if wins == 5:
                    return line

        # transpose the board again to
        # restore the original board
        self.board = [[row[i] for row in self.board] for i in range(5)]

        return None

    def mark_number(self, number):
        self.numbers.append(number)


def build_boards(data):
    temp = data.split('\n\n')
    numbers = [*map(int, temp[0].split(','))]
    board_data = temp[1:]
    boards = []
    for board in board_data:
        boards.append(Board(board))

    return numbers, boards


def day1():
    numbers, boards = build_boards(data)
    for num in numbers:
        for board in boards:
            board.mark_number(num)
            win = board.check_for_win()
            if win:
                # get all the lines that are not the winning line
                d = [x for x in board if x != win]
                # flatten the list and remove any marked numbers
                flattened = [
                    x for sublist in d for x in sublist if x not in board.numbers]
                sum = reduce(operator.add, flattened)
                return sum * num


def day2():
    # def find_winning_board(boards, nums):
    #     for num in numbers:
    #         for board in boards:
    #             board.mark_number(num)
    #             win = board.check_for_win()
    #             if win:
    #                 return board, num

    numbers, boards = build_boards(data)
    while True:
        if len(numbers) > 1:
            num = numbers.pop(0)
            to_remove = None
            for board in boards:
                board.mark_number(num)
                if board.check_for_win():
                    to_remove = board
                    break

            if to_remove:
                print(f'removing {to_remove}')
                boards.remove(to_remove)

            if len(boards) == 1:
                # this one should win:
                # [[[69, 51, 39, 95, 98], [90, 61, 91, 6, 21], [25, 57, 81, 10, 49], [67, 55, 43, 96, 17], [78, 11, 3, 64, 77]]]
                print(boards)
                flatten = [x for sl in boards[0] for x in sl]
                sum = reduce(operator.add, flatten)
                return sum * num
        else:
            print(f'exhausted numbers. here are my boards: {boards}')
            print(f'{num}')
            return


print(day1())
print(day2())
