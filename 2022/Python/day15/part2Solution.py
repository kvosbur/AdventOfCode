import re
from collections import defaultdict


max_coord = 4000000


def convert_to_int(val):
    return [int(x) for x in val]


def get_manhat_dist(sensor, beacon):
    return abs(sensor[0] - beacon[0]) + abs(sensor[1] - beacon[1])


def cover_full_manhat_dist(allowed, center, target_y, dist):
    global max_coord
    if target_y < center[1] - dist or center[1] + dist < target_y:
        return

    y_dist = target_y - center[1]
    x_dist = dist - abs(y_dist)
    x_start = max(center[0] - x_dist, 0)
    x_end = min(center[0] + x_dist, max_coord)
    for x in range(x_start, x_end + 1, 1):
        if x in allowed:
            allowed.remove(x)


def print_cavern(cavern):
    keys = list(cavern.keys())
    keys.sort()

    for key in keys:
        values = list(cavern[key])
        values.sort()
        print("y", key, " x:", values)


def main(lines):
    # got to 190
    global max_coord
    combos = []
    for line in lines:
        reg = re.compile("x=([-]?[0-9]+), y=([-]?[0-9]+)")
        sensor_raw, beacon_raw = reg.findall(line)
        sensor = convert_to_int(sensor_raw)
        beacon = convert_to_int(beacon_raw)
        combos.append([sensor, beacon])

    for y in range(max_coord + 1):
        temp = set()
        for i in range(max_coord + 1):
            temp.add(i)

        for sensor, beacon in combos:
            manhat_dist = get_manhat_dist(sensor, beacon)
            cover_full_manhat_dist(temp, sensor, y, manhat_dist)

        if len(temp) > 0:
            print("y", y, "x", temp)
            exit()

        if y % 1 == 0:
            print(y)

    # do work here
    print("solution is day1 .....")
