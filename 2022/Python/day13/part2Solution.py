from string import digits
from functools import cmp_to_key


def parse_line(line):
    stack = []
    line = line.replace(',', ' ').replace('[', '[ ').replace(']', ' ]').replace('  ', ' ').split(' ')
    for token in line:
        # print(char, stack)
        if token[0] in digits:
            stack.append(int(token))
        elif token == "[":
            stack.append(token)
        elif token == "]":
            parts = []
            popping = True
            while popping:
                next_part = stack.pop()
                if next_part == '[':
                    popping = False
                else:
                    parts.insert(0, next_part)
            stack.append(parts)

    return stack[0]


def is_right_order(left, right):
    if type(left) == int and type(right) == int:
        if left < right:
            return -1
        elif left == right:
            return 0
        else:
            return 1

    if type(left) == int:
        left = [left]
    elif type(right) == int:
        right = [right]

    index = 0
    while index < len(left) and index < len(right):
        cmp = is_right_order(left[index], right[index])
        if cmp != 0:
            return cmp

        index += 1

    if index == len(left) and index < len(right):
        return -1
    elif index < len(left) and index == len(right):
        return 1
    elif index == len(left) and index == len(right):
        return 0


def main(lines):

    line_index = 0
    parsed = []
    while line_index < len(lines):
        first = lines[line_index]
        second = lines[line_index + 1]

        line_index += 3
        first_parsed = parse_line(first)
        second_parsed = parse_line(second)
        parsed.append(first_parsed)
        parsed.append(second_parsed)

    # add distress signal packets
    parsed.append([[2]])
    parsed.append([[6]])

    parsed.sort(key=cmp_to_key(is_right_order))

    product = 1
    for index, packet in enumerate(parsed):
        if packet == [[2]] or packet == [[6]]:
            product = product * (index + 1)

    # do work here
    print("solution is day2 .....", product)
