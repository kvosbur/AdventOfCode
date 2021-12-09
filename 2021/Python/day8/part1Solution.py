
base = {'a': '', 'b': '', 'c':'', 'd':'', 'e':'', 'f':''}


def get_entry(lone_set):
    return list(lone_set)[0]


def get_g_and_e(sets, final):
    pre = sets[1].union(sets[2])  # union of 7 and 4
    # iterate through length 6 values only 9 will have diff of 1 inferring g
    for s in sets[3:6]:
        if len(s-pre) == 1:
            # also use 9 to infer e
            final[9] = s
            sets.remove(s)
            return get_entry(s-pre), get_entry(final[8] - s)  # return g, e


def get_0_and_6(sets, final):
    for s in sets[3:]:
        if len(final[1] - s) == 0:
            final[0] = s
            sets.remove(s)
            final[6] = sets[-1]
            sets.remove(sets[-1])
            return


def get_plugs(sets):
    final = [None for x in range(10)]
    final[1] = sets[0]
    final[4] = sets[2]
    final[7] = sets[1]
    final[8] = sets[9]
    sets = sets[3:9]
    plug_actual = base.copy()
    # plug_actual['a'] = get_entry(sets[1] - sets[0])  # 7 - 1
    # plug_actual['g'], plug_actual['e'] = get_g_and_e(sets, final)
    # get_0_and_6(sets, final)
    # print(final)
    # print(plug_actual)
    # print(sets)

    return final, plug_actual

    # know 0,1,4,6,7,8,9


def evaluate_values(final, values):
    count = 0
    for val in values:
        for f in final:
            if f is not None and f == val:
                count += 1

    return count



def main(lines):
    # do work here
    count = 0
    for line in lines:
        beg, end = line.split(' | ')
        combos = sorted(beg.split(' '), key=len)
        values = [set(x) for x in end.split(' ')]
        sets = [set(x) for x in combos]
        final, plug_actual = get_plugs(sets)
        count += evaluate_values(final, values)

    print("Count: ", count)
