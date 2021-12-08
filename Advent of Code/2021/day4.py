#!/usr/bin/env python3


class Board:
    def __init__(self, lines):
        self.board = [x.split('\n') for x in lines.split('\n')]
        pass

    def __str__(self):
        return str(self.board)

    def __repr__(self):
        return str(self.board)

    def check_for_win(self):
        pass

    def mark_number(self):
        pass


def build_boards(data):
    temp = data.split('\n\n')
    numbers = temp[0]
    board_data = temp[1:]
    boards = []
    for board in board_data:
        boards.append(Board(board))
    
    return numbers, boards

    

with open('input4_sample.txt') as f:
    # numbers = [int(x.strip()) for x in f.readline().split(',')]

    # data = [x.strip() for x in f.read().split('\n\n')]
    # boards = []
    # for board in data:
    #     lines = [x.split() for x in board.split('\n')]
    #     bd = []
    #     for line in lines:
    #         bd.append([[int(x), False] for x in line])
    #     boards.append(bd)
    data = f.read()


print(build_boards(data))


# def mark_number(num):
#     print(f'marking {num}')
#     for board in boards:
#         for line in board:
#             for number in line:
#                 if number[0] == num:
#                     number[1] = True
#                     return

# def check_boards():
#     # check horizontal lines
#     for board in boards:
#         for line in board:
#             win = True
#             for _, status in line:
#                 if not status:
#                     win = False
#                     continue
#             if win:
#                 return line
#     return None
#     # check vertical lines
#     # for board in boards:
        


# def day1():
#     guesses = 0
#     for number in numbers:
#         guesses += 1
#         mark_number(number)
#         if guesses >= 5:
#             w = check_boards()
#             if w:
#                 return w


# def day2():
#     pass


# # print(numbers)
# # print(boards)
# print(day1())
# print(boards)
# # print(day2())
