
grid_size = 1000

grid = [[0 for column in range(grid_size)]for row in range(grid_size)]


def filter_line(first_point, second_point):
    if first_point[0] == second_point[0] or first_point[1] == second_point[1]:
        return False
    return True


def mark_line(first_point, second_point):
    if first_point[0] == second_point[0]:
        start = min(first_point[1], second_point[1])
        end = max(first_point[1], second_point[1])
        for y in range(start, end + 1, 1):
            grid[first_point[0]][y] += 1

    elif first_point[1] == second_point[1]:
        start = min(first_point[0], second_point[0])
        end = max(first_point[0], second_point[0])
        for x in range(start, end + 1, 1):
            grid[x][first_point[1]] += 1


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
        if filter_line(first_point, second_point):
            continue

        mark_line(first_point, second_point)

    count = count_toxic()
    print("toxic squares:", count)
