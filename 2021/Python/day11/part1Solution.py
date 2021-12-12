def increase_energy(locations):
    for row in range(len(locations)):
        for col in range(len(locations[0])):
            increment_location(locations, row, col)


def increment_location(locations, row, col):
    if locations[row][col] != -1:
        locations[row][col] += 1


def do_explode(locations, row, col):
    locations[row][col] = -1
    if row > 0:
        increment_location(locations, row - 1, col)
        if col > 0:
            increment_location(locations, row - 1, col - 1)
        if col < len(locations[row]) - 1:
            increment_location(locations, row - 1, col + 1)

    if row < len(locations) - 1:
        increment_location(locations, row + 1, col)
        if col > 0:
            increment_location(locations, row + 1, col - 1)
        if col < len(locations[row]) - 1:
            increment_location(locations, row + 1, col + 1)

    if col > 0:
        increment_location(locations, row, col - 1)
    if col < len(locations[row]) - 1:
        increment_location(locations, row, col + 1)


def reset(locations):
    for row in range(len(locations)):
        for col in range(len(locations[0])):
            if locations[row][col] == -1:
                locations[row][col] = 0


def explode(locations):
    explosion_seen = True
    flashes = 0
    while explosion_seen:
        explosion_seen = False
        for row in range(len(locations)):
            for col in range(len(locations[0])):
                if locations[row][col] > 9:
                    do_explode(locations, row, col)
                    flashes += 1
                    explosion_seen = True

    reset(locations)
    return flashes


def print_locations(locations):
    for row in locations:
        for col in row:
            print(col, end=" ")
        print("")


def main(lines):
    locations = []
    for line in lines:
        locations.append([int(x) for x in line])

    flashes = 0
    for step in range(100):
        increase_energy(locations)
        flashes += explode(locations)

    print_locations(locations)
    print("Total Flashes:", flashes)

