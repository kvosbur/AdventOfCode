

def contained_in(small, larger):
    return small[0] >= larger[0] and small[1] <= larger[1]


def overlaps(first, second):
    return (first[0] <= second[0] <= first[1]) or (first[0] <= second[1] <= first[1]) or contained_in(first, second) or contained_in(second, first)


def main(lines):
    count = 0
    for line in lines:
        first_raw, second_raw = line.split(',')
        first = [int(x) for x in first_raw.split('-')]
        second = [int(x) for x in second_raw.split('-')]

        if overlaps(first, second):
            count += 1

    print("solution is day1 .....", count)
