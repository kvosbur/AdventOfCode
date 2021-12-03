
def main(lines):
    increased = 0

    numbers = [int(n) for n in lines]
    for i, val in enumerate(numbers[:-1]):
        if val < numbers[i + 1]:
            increased += 1
        if val == numbers[i + 1]:
            print('edge case found at line ${i}')

    print(increased)
