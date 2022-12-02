import pprint


def do_step(template, pairs):
    new_str = []
    for index in range(len(template)):
        pair = template[index: index + 2]
        if pair in pairs:
            new_str.append(pairs[pair][:-1])
        else:
            new_str.append(pair[0])
    return "".join(new_str)


def next_level(levels, use_level_index):
    pairs = levels[use_level_index]
    upgraded_pairs = {}
    for key, value in pairs.items():
        print(use_level_index, key)
        out = do_step(value, pairs)
        upgraded_pairs[key] = out

    levels.append(upgraded_pairs)


def count_characters(template):
    counts = {}
    for char in template:
        if char not in counts:
            counts[char] = 0
        counts[char] += 1
    return counts


def main(lines):
    template = lines[0]

    levels = []
    pairs = {}
    for line in lines[2:]:
        pair, result = line.split(' -> ')
        pairs[pair] = pair[0] + result + pair[1]

    levels.append(pairs)

    for i in range(6):
        print("level", i)
        next_level(levels, i)

    out = do_step(template, levels[3])
    template = do_step(out, levels[5])
    # print(pairs)
    # for step in range(40):
    #     print(step)
    #     template = do_step(template, pairs)
    #
    print(len(template))
    counts = count_characters(template)
    largest = max(counts.values())
    smallest = min(counts.values())
    print("Target Diff", (largest - smallest))
