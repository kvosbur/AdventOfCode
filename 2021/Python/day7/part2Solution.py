def score_position(crabs, costs, position):
    total = 0
    for crab in crabs:
        total += costs[abs(position - crab)]
    return total


def main(lines):
    crabs = [int(x) for x in lines[0].split(',')]
    crabs = sorted(crabs)

    costs = [0]
    for i in range(1, crabs[-1] + 1):
        costs.append(costs[i-1] + i)

    best_position = 0
    best_value = score_position(crabs, costs, 0)
    for i in range(1, crabs[-1]):
        temp = score_position(crabs, costs, i)
        if temp < best_value:
            best_position = i
            best_value = temp

    print("Best postion: ", best_position)
    print("Best value: ", best_value)

