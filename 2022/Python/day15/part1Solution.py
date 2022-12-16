import re
from collections import defaultdict

def convert_to_int(val):
    return [int(x) for x in val]


def get_manhat_dist(sensor, beacon):
    return abs(sensor[0] - beacon[0]) + abs(sensor[1] - beacon[1])


def cover_full_manhat_dist(cavern, center, dist):
    y_dist = dist
    while y_dist >= -dist:
        new_y = center[1] + y_dist
        if new_y != 2000000:
            y_dist -= 1
            continue
        x_dist = dist - abs(y_dist)
        x_start = center[0] - x_dist
        x_end = center[0] + x_dist
        for x in range(x_start, x_end + 1, 1):
            new_coord = [x, new_y]
            cavern[new_coord[1]].add(new_coord[0])
        y_dist -= 1


def print_cavern(cavern):
    keys = list(cavern.keys())
    keys.sort()

    for key in keys:
        values = list(cavern[key])
        values.sort()
        print("y", key, " x:", values)


def main(lines):
    cavern = defaultdict(lambda: set())
    beacons = []
    for line in lines:
        reg = re.compile("x=([-]?[0-9]+), y=([-]?[0-9]+)")
        sensor_raw, beacon_raw = reg.findall(line)
        sensor = convert_to_int(sensor_raw)
        beacon = convert_to_int(beacon_raw)
        beacons.append(beacon)
        print(line, sensor, beacon)

        manhat_dist = get_manhat_dist(sensor, beacon)
        print("man dist:", manhat_dist)
        cover_full_manhat_dist(cavern, sensor, manhat_dist)

    # print_cavern(cavern)
    # remove beacons from cavern since they apparently don't count :(
    for beacon in beacons:
        row = cavern[beacon[1]]
        if beacon[0] in row:
            row.remove(beacon[0])

    # print_cavern(cavern)
    print(len(cavern[2000000]))
    # do work here
    print("solution is day1 .....")
