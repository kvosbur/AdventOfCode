
def get_bitcounts(lines):
    total_bits_per = len(lines[0].replace("\n", ""))
    bit_counts = [[0, 0] for bit in range(total_bits_per)]

    for line in lines:
        for index, bit in enumerate(line.replace('\n', '')):
            bit_val = int(bit)
            bit_counts[index][bit_val] += 1

    return bit_counts


def get_preferred_bit(bitcounts, index):
    if bitcounts[index][0] > bitcounts[index][1]:
        return "0"
    else:
        return "1"


def does_line_pass_filter(line, index, target_bit, use_most_common):
    if use_most_common:
        return line[index] == target_bit
    else:
        return line[index] != target_bit


def filter_lines(lines, bitcounts, index, use_most_common):
    target_bit = get_preferred_bit(bitcounts, index)

    new_lines = []
    for line in lines:
        if does_line_pass_filter(line, index, target_bit, use_most_common):
            new_lines.append(line)

    return new_lines


def main(lines):

    iterative_lines = lines
    index = 0
    while len(iterative_lines) > 1:
        bitcounts = get_bitcounts(iterative_lines)
        iterative_lines = filter_lines(iterative_lines, bitcounts, index, True)
        index += 1

    oxygen = int(iterative_lines[0], 2)

    iterative_lines = lines
    index = 0
    while len(iterative_lines) > 1:
        bitcounts = get_bitcounts(iterative_lines)
        iterative_lines = filter_lines(iterative_lines, bitcounts, index, False)
        index += 1

    co2_scrubber = int(iterative_lines[0], 2)

    print("oxygen", oxygen)
    print("CO2 scruuber", co2_scrubber)
    print("Life support Rating", (oxygen * co2_scrubber))

