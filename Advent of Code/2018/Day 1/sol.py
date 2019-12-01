#!/usr/bin/env python3

import operator
from functools import reduce
from itertools import accumulate
from collections import defaultdict

def get_lines(fname):
    with open(fname) as f:
        return [int(i) for i in f.read().split()]

def part1(inp):
    return reduce(operator.add, inp)

def part2(inp):
    visited = defaultdict(list)
    freqs = list(accumulate(inp))
    visited[0] = 1
    
    # shift = freqs[-1]
    # for f in freqs:
    #     visited[f % shift].append(f)
    #
    # min_freq, min_diff = None, 0
    #
    # for v in visited.values():
    #     min_diff = v[1] - v[0] if ((v[1] - v[0]) < min_diff) else min_diff
    #
    # return min_diff


if __name__ == "__main__":
    inp = get_lines("input")
    print(part1(inp))
    print(part2(inp))
