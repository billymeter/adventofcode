#!/usr/bin/env python3

with open('input2.txt') as f:
    data = [x.split() for x in f.readlines()]


def day1():
    horizontal = 0
    depth = 0
    for e in data:
        d = e[0]
        v = int(e[1])
        if 'up' in e[0]:
            depth -= v
        if 'down' in d:
            depth += v
        if 'forward' in d:
            horizontal += v
    return horizontal * depth


def day2():
    horizontal = 0
    depth = 0
    aim = 0
    for e in data:
        d = e[0]
        v = int(e[1])
        if 'up' in e[0]:
            # depth -= v
            aim -= v
        if 'down' in d:
            # depth += v
            aim += v
        if 'forward' in d:
            horizontal += v
            depth += (aim * v)
    return horizontal * depth


print(day1())
print(day2())
