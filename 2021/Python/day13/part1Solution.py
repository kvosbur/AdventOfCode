def get_max(l, index):
    most = l[0][index]
    for item in l:
        if item[index] > most:
            most = item[index]
    return most


def place_dots(dots, sheet):
    for dot in dots:
        x, y = dot
        sheet[y][x] = '#'


def pretty_print_sheet(sheet):
    for y in sheet:
        for x in y:
            print(x, end="")
        print("")


def do_fold(sheet, fold):
    fold_var, fold_value = fold
    if fold_var == "x":
        for y in sheet:
            for index, x in enumerate(y[fold_value:]):
                if x == '#':
                    y[fold_value - index] = '#'
                    y[fold_value + index] = '.'
    else:
        for index, y in enumerate(sheet[fold_value:]):
            for x in range(len(y)):
                if y[x] == '#':
                    sheet[fold_value - index][x] = '#'
                    sheet[fold_value + index][x] = '.'


def count_dots_on_sheet(sheet):
    count = 0
    for y in sheet:
        for x in y:
            if x == '#':
                count += 1
    return count


def main(lines):
    dots = []
    for line in lines:
        if line == "":
            break
        dots.append([int(x) for x in line.split(',')])

    folds = []
    for line in lines:
        if "fold along" in line:
            var, value = line[11:].split('=')
            folds.append([var, int(value)])

    max_x = get_max(dots, 0)
    max_y = get_max(dots, 1)

    sheet = [['.' for j in range(max_x + 1)] for i in range(max_y + 1)]
    place_dots(dots, sheet)
    do_fold(sheet, folds[0])

    print("Dots:", count_dots_on_sheet(sheet))

