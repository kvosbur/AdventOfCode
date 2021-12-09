def get_entry(lone_set):
    return list(lone_set)[0]


def get_9(sets, final):
    pre = final[7].union(final[4])  # union of 7 and 4
    # iterate through length 6 values only 9 will have diff of 1 inferring g
    for s in sets[3:6]:
        if len(s-pre) == 1:
            final[9] = s
            sets.remove(s)
            return


def get_0_and_6(sets, final):
    for s in sets[3:]:
        if len(final[1] - s) == 0:
            final[0] = s
            sets.remove(s)
            final[6] = sets[-1]
            sets.remove(sets[-1])
            return


def get_5(sets, final):
    pre = final[4] - final[1]  # bd
    for s in sets:
        if len(pre-s) == 0:
            final[5] = s
            sets.remove(s)
            return


def get_2_and_3(sets, final):
    pre = final[8] - final[9]  # e
    for s in sets:
        if len(pre - s) == 0:
            final[2] = s
            sets.remove(s)
            final[3] = sets[0]
            sets.remove(sets[0])
            return


def get_numbers(sets):
    final = [None for x in range(10)]
    final[1] = sets[0]
    final[4] = sets[2]
    final[7] = sets[1]
    final[8] = sets[9]
    sets = sets[3:9]

    get_9(sets, final)
    get_0_and_6(sets, final)
    get_5(sets, final)
    get_2_and_3(sets, final)

    return final


def evaluate_values(final, values):
    number = ''
    for val in values:
        for i, f in enumerate(final):
            if f == val:
                number += str(i)
    return int(number)


def main(lines):
    count = 0
    for line in lines:
        beg, end = line.split(' | ')
        combos = sorted(beg.split(' '), key=len)
        values = [set(x) for x in end.split(' ')]
        sets = [set(x) for x in combos]
        final = get_numbers(sets)
        count += evaluate_values(final, values)

    print("Sum of values: ", count)
