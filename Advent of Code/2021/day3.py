#!/usr/bin/env python3

with open('input3.txt') as f:
    data = [x.strip() for x in f.readlines()]


def day1():
    gamma = []
    epsilon = []
    for pos in range(len(data[0])):
        ones = 0
        zeros = 0
        for line in range(len(data)):
            if '1' in data[line][pos]:
                ones += 1
            else:
                zeros += 1

        if ones > zeros:
            gamma.append('1')
            epsilon.append('0')
        else:
            gamma.append('0')
            epsilon.append('1')

    gamma = int(''.join(gamma), 2)
    epsilon = int(''.join(epsilon), 2)

    return (gamma * epsilon)


def day2():
    pass


print(day1())
# print(day2())
