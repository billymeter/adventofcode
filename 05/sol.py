#!/usr/bin/env python3

import re
import string

lower_upper = "aA|bB|cC|dD|eE|fF|gG|hH|iI|jJ|kK|lL|mM|nN|oO|pP|qQ|rR|sS|tT|uU|vV|wW|xX|yY|zZ"
upper_lower = "Aa|Bb|Cc|Dd|Ee|Ff|Gg|Hh|Ii|Jj|Kk|Ll|Mm|Nn|Oo|Pp|Qq|Rr|Ss|Tt|Uu|Vv|Ww|Xx|Yy|Zz"

def get_input(fname):
    with open(fname) as f:
        return f.read()

def react(inp):
    done = False
    start = inp
    while not done:
        test = re.sub(upper_lower, '', start)
        test = re.sub(lower_upper, '', test)
        if start == test:
            done = True
        else:
            start = test
    return test


def part1(inp):
    return len(react(inp))

def part2(inp):
    shortest = part1(inp)
    test = inp
    for c in string.ascii_lowercase:
        test = re.sub(c, '', test)
        test = re.sub(c.upper(), '', test)
        test = react(test)
        if len(test) < shortest: shortest = len(test)
        test = inp
    return shortest


if __name__ == "__main__":
    inp = get_input("input").strip()
    print(part1(inp))
    print(part2(inp))
