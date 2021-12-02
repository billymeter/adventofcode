#!/usr/bin/env python3

with open('input1.txt') as f:
    data = [int(x.strip()) for x in f.readlines()]


def day1():
    prev = data[0]
    increases = 0
    for e in data[1:]:
        if e > prev:
            increases += 1
        prev = e
    return increases


def day2():
    num_groups = [data[i:i+3] for i in range(0, len(data)-3 + 1)]
    prev = sum(num_groups[0])
    increases = 0
    for n in num_groups[1:]:
        s = sum(n)
        if s > prev:
            increases += 1
        prev = s
    return increases


print(day1())
print(day2())
