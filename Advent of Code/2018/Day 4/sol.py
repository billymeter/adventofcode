#!/usr/bin/env python3

import re
import operator
from datetime import datetime
from functools import reduce
from collections import defaultdict

def get_lines(fname):
    with open(fname) as f:
        return f.read().split('\n')

def parse_line(line):
    pass

def parse(data):
    guards = defaultdict(dict)
    current_guard = 0
    it = iter(data)

    for line in it:
        if not line: continue
        if "shift" in line:
            m = re.search(".+ #(\d+).*", line)
            if m:
                current_guard = m.groups()[0]
                if 'sleep' not in guards[current_guard].keys():
                    guards[current_guard] = defaultdict(list)
        if "asleep" in line:
            m = re.search("\[(.*)\]", line)
            if m:
                sleep = datetime.strptime(m.groups()[0], "%Y-%m-%d %H:%M")
                guards[current_guard]['sleep'].append(sleep)
            next_line = next(it)
            m = re.search("\[(.+)\]", next_line)
            if m:
                wake = datetime.strptime(m.groups()[0], "%Y-%m-%d %H:%M")
                guards[current_guard]['wake'].append(wake)
                guards[current_guard]['sleep_seconds'].append((wake - sleep).total_seconds())
    return guards

def most_time_asleep(guards):
    sleep = defaultdict(dict)

    for k in guards:
        sleep[k]['sleep'] = reduce(operator.add, guards[k]['sleep_seconds'], 0)
        sleep[k]['mins'] = []
        for min in zip(guards[k]['sleep'], guards[k]['wake']):
            sleep[k]['mins'] += [x for x in range(min[0].minute, min[1].minute)]
    return sorted(sleep.items(), key=lambda x: x[1]['sleep'], reverse=True)

def part1(guards):
    gid = int(guards[0][0])
    return gid * max(guards[0][1]['mins'], key=guards[0][1]['mins'].count)

def part2(guards):
    max_min, max_count, gid = 0, 0, 0
    for g in guards:
        if not g[1]['mins']: continue
        d = defaultdict(int)
        for m in g[1]['mins']: d[m] += 1
        test = max(d.items(), key=lambda x: x[1])
        if test[1] > max_count:
            max_count = test[1]
            max_min = test[0]
            gid = int(g[0])
    return gid * max_min

if __name__ == "__main__":
    # sort = sorted(get_lines("input"))
    # with open("input_sorted", "w") as f:
    #     f.write('\n'.join(sort))
    data = get_lines("input_sorted")
    guards = parse(data)
    sleep_data = most_time_asleep(guards)
    print(part1(sleep_data))
    print(part2(sleep_data))
