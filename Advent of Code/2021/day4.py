#!/usr/bin/env python3

from functools import reduce
import operator


class Board:
    def __init__(self, lines):
        self.board = [[*map(int, x.split('\n')[0].split())]
                      for x in lines.split('\n')]
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
        self.board = [[row[i] for row in self.board] for i in range(5)]

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


with open('input4_sample.txt') as f:
    data = f.read()


def day1():
    numbers, boards = build_boards(data)
    for num in numbers:
        for board in boards:
            board.mark_number(num)
            win = board.check_for_win()
            if win:
                # get all the lines that are not the winning line
                d = [x for x in board if x != win]
                # flatten the list and get its sum
                sum = reduce(operator.add, [
                             x for sublist in d for x in sublist])
                return sum * num


print(day1())
