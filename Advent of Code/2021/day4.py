#!/usr/bin/env python3

with open('input4.txt') as f:
    def remove_empty(x):
        if x == '':
            return True
        return False
    numbers = [int(x.strip()) for x in f.readline().split(',')]

    data = [filter(remove_empty, x.split('\n'))
            for x in f.read().split('\n\n')]
    boards = []
    for x in data:
        b = []
        for y in x:
            b.append(y)
        boards.append(x)


def day1():
    pass


def day2():
    pass


print(numbers)
print(boards)
print(day1())
print(day2())
