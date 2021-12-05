
grid_size = 1000

grid = [[0 for column in range(grid_size)]for row in range(grid_size)]


def mark_line(first_point, second_point):
    xdirection = 0
    if first_point[0] != second_point[0]:
        xdirection = 1 if first_point[0] < second_point[0] else -1

    ydirection = 0
    if first_point[1] != second_point[1]:
        ydirection = 1 if first_point[1] < second_point[1] else -1

    x = first_point[0]
    y = first_point[1]
    while x != second_point[0] or y != second_point[1]:
        grid[x][y] += 1
        x += xdirection
        y += ydirection

    grid[x][y] += 1


def count_toxic():
    count = 0
    for x in range(len(grid)):
        for y in range(len(grid[0])):
            if grid[x][y] > 1:
                count += 1

    return count


def main(lines):
    # do work here
    for line in lines:
        first_point_str, second_point_str = line.split(' -> ')
        first_point = [int(x) for x in first_point_str.split(',')]
        second_point = [int(x) for x in second_point_str.split(',')]

        mark_line(first_point, second_point)

    count = count_toxic()
    print("toxic squares:", count)


# 13354 too low
# 19663 is correct
