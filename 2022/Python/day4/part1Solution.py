
def contained_in(small, larger):
    return small[0] >= larger[0] and small[1] <= larger[1]


def main(lines):
    count = 0
    for line in lines:
        first_raw, second_raw = line.split(',')
        first = [int(x) for x in first_raw.split('-')]
        second = [int(x) for x in second_raw.split('-')]

        if contained_in(first, second):
            print(first, second)
            count += 1

        elif contained_in(second, first):
            print(first, second)
            count += 1

    print("solution is day1 .....", count)
