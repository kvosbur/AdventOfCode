def do_step(template, pairs):
    new_str = ""
    for index in range(len(template)):
        pair = template[index: index + 2]
        new_str += template[index]
        if pair in pairs:
            new_str += pairs[pair]
    return new_str


def count_characters(template):
    counts = {}
    for char in template:
        if char not in counts:
            counts[char] = 0
        counts[char] += 1
    return counts


def main(lines):
    template = lines[0]

    pairs = {}
    for line in lines[2:]:
        pair, result = line.split(' -> ')
        pairs[pair] = result

    print(pairs)
    for step in range(10):
        template = do_step(template, pairs)

    print(len(template))
    counts = count_characters(template)
    largest = max(counts.values())
    smallest = min(counts.values())
    print("Target Diff", (largest - smallest))
