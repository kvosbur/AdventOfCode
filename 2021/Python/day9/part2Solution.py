def add_to_stack_if_not_seen(stack, row, col, seen):
    if seen[row][col] != 1:
        stack.append([row, col])
        seen[row][col] = 1


def get_basin_size(rows, seen, start_row, start_col):
    stack = [(start_row, start_col)]
    seen[start_row][start_col] = 1
    found = 0
    while len(stack) > 0:
        row, column = stack.pop()
        if row > 0 and rows[row - 1][column] < 9:
            add_to_stack_if_not_seen(stack, row - 1, column, seen)
        if row < len(rows) - 1 and rows[row + 1][column] < 9:
            add_to_stack_if_not_seen(stack, row + 1, column, seen)
        if column > 0 and rows[row][column - 1] < 9:
            add_to_stack_if_not_seen(stack, row, column - 1, seen)
        if column < len(rows[0]) - 1 and rows[row][column + 1] < 9:
            add_to_stack_if_not_seen(stack, row, column + 1, seen)
        found += 1
    return found


def get_risk(rows):
    sizes = []
    seen = [[0 for y in range(len(rows[0]))] for x in range(len(rows))]
    for row in range(len(rows)):
        for column in range(len(rows[0])):
            current = rows[row][column]
            lower_than_all = True
            if row > 0 and rows[row - 1][column] <= current:
                lower_than_all = False
            if row < len(rows) - 1 and rows[row + 1][column] <= current:
                lower_than_all = False
            if column > 0 and rows[row][column - 1] <= current:
                lower_than_all = False
            if column < len(rows[0]) - 1 and rows[row][column + 1] <= current:
                lower_than_all = False
            if lower_than_all:
                size = get_basin_size(rows, seen, row, column)
                sizes.append(size)

    sorted_sizes = sorted(sizes)
    return sorted_sizes[-1] * sorted_sizes[-2] * sorted_sizes[-3]


def main(lines):
    rows = []
    for line in lines:
        rows.append([int(x) for x in line])
    temp = get_risk(rows)
    print("Product:", temp)
