#!/usr/bin/env python3

import re
from datetime import datetime
from collections import defaultdict

def get_lines(fname):
    with open(fname) as f:
        return f.read().split('\n')

def parse_line(line):
    pass

def parse(data):
    # [1518-03-10 23:56] Guard #443 begins shift
    # [1518-03-11 00:22] falls asleep
    # [1518-03-11 00:27] wakes up
    guards = defaultdict(dict)
    it = iter(data)
    current_guard = 0

    for line in it:
        if not line: continue
        if "shift" in line:
            m = re.search(".+ #(\d+).*", line)
            if m: current_guard = m.groups()[0]
        if "asleep" in line:
            m = re.search("[(.+)]", line)
            if m:
                print("re found sleep")
                date = datetime.strptime(m.groups()[0], "%Y-%m-%d %H:%M")
                guards[current_guard]['sleep'].append(date)
            next_line = next(it)
            m = re.search("[(.+)]", next_line)
            if m:
                print("re found wake")
                date = datetime.strptime(m.groups()[0], "%Y-%m-%d %H:%M")
                guards[current_guard]['wake'].append(date)

    return guards

if __name__ == "__main__":
    # sort = sorted(get_lines("input"))
    # with open("input_sorted", "w") as f:
    #     f.write('\n'.join(sort))
    data = get_lines("input_sorted")
    print(data[1:4])
    print(parse(data[4]))
