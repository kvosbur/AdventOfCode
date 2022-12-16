import re

max_coord = 4000000


def convert_to_int(val):
    return [int(x) for x in val]


def get_manhat_dist(sensor, beacon):
    return abs(sensor[0] - beacon[0]) + abs(sensor[1] - beacon[1])


def cover_full_manhat_dist(center, target_y, dist):
    global max_coord

    y_dist = target_y - center[1]
    x_dist = dist - y_dist
    if y_dist < 0:
        x_dist = dist + y_dist
    # x_dist = dist - abs(y_dist)
    # x_start = max(center[0] - x_dist, 0)
    x_start = center[0] - x_dist
    if x_start < 0:
        x_start = 0
    # x_end = min(center[0] + x_dist, max_coord)
    x_end = center[0] + x_dist
    if x_end > max_coord:
        x_end = max_coord

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
        combos.append([sensor, beacon, get_manhat_dist(sensor, beacon)])

    found_x = 0
    found_y = 0
    for y in range(max_coord + 1):
        pairs = []

        for sensor, beacon, manhat_dist in combos:
            if y < sensor[1] - manhat_dist or sensor[1] + manhat_dist < y:
                continue
            temp = cover_full_manhat_dist(sensor, y, manhat_dist)
            pairs += [temp]

        pairs.sort()
        first = pairs[0]
        results = []
        for pair in pairs[1:]:
            if pair[0] - 1 <= first[1] <= pair[1]:
                new = [first[0], pair[1]]
                first = new
            elif pair[1] < first[1]:
                pass
            else:
                results.append(first)
                results.append(pair)
                break

        if len(results) > 1:
            found_x = results[0][1] + 1
            found_y = y
            break

        if y % 20000 == 0:
            print("progress", y / max_coord * 100, "%")

    print("found at", found_x, found_y)
    print("goal frequency", (found_x * max_coord) + found_y)

    # found
    # at
    # 2655411
    # 3166538
    # goal
    # frequency
    # 10621647166538
