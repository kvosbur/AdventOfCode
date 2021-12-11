def get_risk(rows):
    risks = []
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
                risks.append(current)

    risks = [x + 1 for x in risks]
    return sum(risks)


def main(lines):
    rows = []
    for line in lines:
        rows.append([int(x) for x in line])
    temp = get_risk(rows)
    print("Risk Sums:", temp)
