#!/usr/bin/env python3

from collections import defaultdict
from math import sqrt

def get_lines(fname):
    with open(fname) as f:
        lines = f.read().split('\n')
    claims = []
    for line in lines:
        if line:
            l = line.split()
            claim, point, area = l[0], l[2][:-1].split(','), l[3].split('x')
            claims.append((claim, point, area))
    return claims

def claim_fabric(claims):
    fabric = defaultdict(list)
    for claim in claims:
        c, x, y, w, h = claim[0], int(claim[1][0]), int(claim[1][1]), int(claim[2][0]), int(claim[2][1])
        for j in range(y + 1, y + h + 1):
            for i in range(x + 1, x + w + 1):
                fabric[(i,j)].append(c)
    return fabric

def find_overlap(fabric):
    overlap = [x for x in fabric.values() if len(x) > 1]
    return len(overlap)


def part1(inp):
    fabric = claim_fabric(claims)
    return find_overlap(fabric)

if __name__ == "__main__":
    claims = get_lines("input")
    print(part1(claims))
