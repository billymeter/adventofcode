#!/usr/bin/env python3

# from functools import map

def get_lines(fname):
    with open(fname) as f:
        return f.read().split()

def compare(w1, w2):
    diffs = 0
    for i,c in enumerate(w1):
        if c != w2[i]:
            if diffs > 0:
                return False
            else:
                diffs += 1
    return True

def process(inp):
    ret = []
    for i in range(len(inp)):
        for j in range(i + 1, len(inp)):
            if compare(inp[i], inp[j]):
                ret.append(inp[j])
                ret.append(inp[i])
                return ret
            else:
                continue
    return ret

def part2(inp):
    vals = process(inp)
    answer = []
    for i,c in enumerate(vals[0]):
        if c != vals[1][i]:
            continue
        else:
            answer.append(c)
    return ''.join(answer)

if __name__ == "__main__":
    inp = get_lines("input")
    sets = part2(inp)
    print(part2(inp))
