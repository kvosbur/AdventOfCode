stacks = [
    ['R', 'S', 'L', 'F', 'Q'],
    ['N', 'Z', 'Q', 'G', 'P', 'T'],
    ['S', 'M', 'Q', 'B'],
    ['T', 'G', 'Z', 'J', 'H', 'C', 'B', 'Q'],
    ['P', 'H', 'M', 'B', 'N', 'F', 'S'],
    ['P', 'C', 'Q', 'N', 'S', 'L', 'V', 'G'],
    ['W', 'C', 'F'],
    ['Q', 'H', 'G', 'Z', 'W', 'V', 'P', 'M'],
    ['G', 'Z', 'D', 'L', 'C', 'N', 'R']
]


def main(lines):

    for line in lines:
        parts = line.split(' ')
        amount = int(parts[1])

        # subtract 1 to mimic 0 indexed lists
        from_stack = int(parts[3]) - 1
        to_stack = int(parts[5]) - 1

        print(amount, from_stack, to_stack)
        for i in range(amount):
            popped = stacks[from_stack].pop()
            stacks[to_stack].append(popped)

    sol = ''.join([stack[-1] for stack in stacks])
    print("solution is.....", sol)
