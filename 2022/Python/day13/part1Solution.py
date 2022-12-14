from string import digits
import ast


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
        # print("comparing:", left, right)
        if left < right:
            print("IN ORDER", left, ":", right)
            return -1
        elif left == right:
            print("SAME", left, ":", right)
            return 0
        else:
            print("OUT OF ORDER", left, ":", right)
            return 1

    if type(left) == int:
        left = [left]
    elif type(right) == int:
        right = [right]

    index = 0
    # print("comparing:", left, right)
    while index < len(left) and index < len(right):
        cmp = is_right_order(left[index], right[index])
        if cmp != 0:
            return cmp

        index += 1

    # print("debug", index, len(left), len(right), left, right)
    if index == len(left) and index < len(right):
        print("IN ORDER", left, ":", right)
        return -1
    elif index < len(left) and index == len(right):
        print("OUT OF ORDER", left, ":", right)
        return 1
    elif index == len(left) and index == len(right):
        print("SAME", left, ":", right)
        return 0


def main(lines):

    line_index = 0
    num_correct = 0
    while line_index < len(lines):
        first = lines[line_index]
        second = lines[line_index + 1]

        line_index += 3
        first_parsed = parse_line(first)
        second_parsed = parse_line(second)

        print(first_parsed)
        print(second_parsed)
        is_correct = is_right_order(first_parsed, second_parsed)
        if is_correct == -1:
            # print("adding", line_index // 3)
            num_correct += (line_index // 3)

        print(line_index // 3, "right" if is_correct == -1 else "wrong")
        print("")

    # do work here
    print("solution is day1 .....", num_correct)
