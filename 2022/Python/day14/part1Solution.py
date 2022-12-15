

def print_cavern(cavern):
    for y in range(0, len(cavern)):
        for x in range(0, len(cavern[0])):
            print(cavern[y][x], end="")
        print("")


def place_line(cavern, x_lower, start, end):
    if start[0] == end[0]:
        lower = min(start[1], end[1])
        upper = max(start[1], end[1])
        # vertical line
        for y in range(lower, upper + 1):
            place_character(cavern, x_lower, [start[0], y], '#')
    else:
        # horizontal line
        lower = min(start[0], end[0])
        upper = max(start[0], end[0])
        for x in range(lower, upper + 1):
            place_character(cavern, x_lower, [x, start[1]], '#')


def place_sand(cavern, lower_x):
    coord = [500, 0]
    moving = True
    while moving:
        # print("debug", lower_x, coord, len(cavern), len(cavern[0]))
        if coord[1] == len(cavern) - 1 or coord[0] == lower_x or coord[0] - lower_x == len(cavern[0]) - 1:
            print("reached the abyss")
            return True
        below = peek_character(cavern, lower_x, [coord[0], coord[1] + 1])
        bottom_left = peek_character(cavern, lower_x, [coord[0] - 1, coord[1] + 1])
        bottom_right = peek_character(cavern, lower_x, [coord[0] + 1, coord[1] + 1])

        if below == '.':
            coord[1] += 1
        elif bottom_left == '.':
            coord[0] -= 1
            coord[1] += 1
        elif bottom_right == '.':
            coord[0] += 1
            coord[1] += 1
        else:
            place_character(cavern, lower_x, coord, 'o')
            return False


def place_character(cavern, x_lower, coord, character):
    x, y = coord
    new_x = x - x_lower
    cavern[y][new_x] = character


def peek_character(cavern, x_lower, coord):
    x, y = coord
    new_x = x - x_lower
    return cavern[y][new_x]


def main(lines):
    all_x = []
    all_y = []
    all_lines = []
    for line in lines:
        parts = line.split(' -> ')
        parsed = []
        for part in parts:
            x, y = part.split(',')
            parsed.append([int(x), int(y)])
            all_x.append(int(x))
            all_y.append(int(y))

        all_lines.append(parsed)

    lower_x = min(all_x) - 1
    max_x = max(all_x) + 1
    max_y = max(all_y) + 1

    cavern = [['.' for x in range(lower_x, max_x + 1)] for y in range(0, max_y + 1)]

    place_character(cavern, lower_x, [500, 0], '+')

    for lines in all_lines:
        for index in range(len(lines) - 1):
            first = lines[index]
            second = lines[index + 1]
            place_line(cavern, lower_x, first, second)

    finished = False
    count = 0
    while not finished:
        finished = place_sand(cavern, lower_x)
        count += 1
        print("sand count:", count)
        # print_cavern(cavern)

    print('result:', count - 1)
    # print_cavern(cavern)
