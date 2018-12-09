#!/usr/bin/env python3

def get_input(fname):
    with open(fname) as f:
        return f.read()


def part1(inp):
    polymer = list(inp)
    new = polymer
    done = False

    while not done:
        polymer = new
        new = []
        it = iter(polymer)
        for c in it:
            try:
                c2 = next(it)
                if abs(ord(c) - ord(c2)) == 0x20:
                    continue
                else:
                    new.append(c)
                    new.append(c2)
            except:
                new.append(c)

        done = True if new == polymer else False
    return ''.join(new)

if __name__ == "__main__":
    inp = get_input("input")
    print(part1(inp))
