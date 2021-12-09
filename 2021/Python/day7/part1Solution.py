def score_position(crabs, position):
    total = 0
    for crab in crabs:
        total += abs(position - crab)
    return total


def main(lines):
    crabs = [int(x) for x in lines[0].split(',')]
    crabs = sorted(crabs)

    best_position = 0
    best_value = score_position(crabs, 0)
    for i in range(1, crabs[-1]):
        temp = score_position(crabs, i)
        if temp < best_value:
            best_position = i
            best_value = temp

    print("Best postion: ", best_position)
    print("Best value: ", best_value)
