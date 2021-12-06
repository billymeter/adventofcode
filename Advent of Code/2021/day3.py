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
    nums = data
    for pos in range(len(nums[0])):
        ones = []
        zeros = []
        for line in range(len(nums)):
            if '1' in nums[line][pos]:
                ones.append(nums[line])
            else:
                zeros.append(nums[line])
        if len(ones) > len(zeros) or len(ones) == len(zeros):
            nums = ones
        else:
            nums = zeros

        if len(nums) == 1:
            break
    oxygen = int(''.join(nums), 2)

    nums = data
    for pos in range(len(nums[0])):
        ones = []
        zeros = []
        for line in range(len(nums)):
            if '1' in nums[line][pos]:
                ones.append(nums[line])
            else:
                zeros.append(nums[line])
        if len(ones) > len(zeros) or len(ones) == len(zeros):
            nums = zeros
        else:
            nums = ones

        if len(nums) == 1:
            break
    co2 = int(''.join(nums), 2)

    return (oxygen * co2)


print(day1())
print(day2())
