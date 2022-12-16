import re

max_coord = 4000000


def convert_to_int(val):
    return [int(x) for x in val]


def get_manhat_dist(sensor, beacon):
    return abs(sensor[0] - beacon[0]) + abs(sensor[1] - beacon[1])


def cover_full_manhat_dist(center, target_y, dist):
    global max_coord
    if target_y < center[1] - dist or center[1] + dist < target_y:
        return

    y_dist = target_y - center[1]
    x_dist = dist - abs(y_dist)
    x_start = max(center[0] - x_dist, 0)
    x_end = min(center[0] + x_dist, max_coord)
    return [x_start, x_end]


def print_cavern(cavern):
    keys = list(cavern.keys())
    keys.sort()

    for key in keys:
        values = list(cavern[key])
        values.sort()
        print("y", key, " x:", values)


def main(lines):
    global max_coord
    combos = []
    for line in lines:
        reg = re.compile("x=([-]?[0-9]+), y=([-]?[0-9]+)")
        sensor_raw, beacon_raw = reg.findall(line)
        sensor = convert_to_int(sensor_raw)
        beacon = convert_to_int(beacon_raw)
        combos.append([sensor, beacon])

    found_x = 0
    found_y = 0
    for y in range(max_coord + 1):
        pairs = []

        for sensor, beacon in combos:
            manhat_dist = get_manhat_dist(sensor, beacon)
            temp = cover_full_manhat_dist(sensor, y, manhat_dist)
            if temp is not None:
                pairs.append(temp)

        pairs.sort(key=lambda x: x[0])
        while len(pairs) != 1:
            first = pairs.pop(0)
            second = pairs.pop(0)
            if second[0] - 1 <= first[1] <= second[1]:
                new = [first[0], second[1]]
                pairs.insert(0, new)
            elif second[1] < first[1]:
                pairs.insert(0, first)
            else:
                pairs.insert(0, first)
                pairs.insert(1, second)
                break

        if len(pairs) > 1:
            found_x = pairs[0][1] + 1
            found_y = y
            break

        if y % 10000 == 0:
            print("progress", y / max_coord * 100, "%")

    print("found at", found_x, found_y)
    print("goal frequency", (found_x * max_coord) + found_y)
