increased = 0

with open("part1.txt", "r") as f:
    numbers = [int(n) for n in f.readlines()]
    for i, val in enumerate(numbers[:-1]):
        if val < numbers[i + 1]:
            increased += 1
        if val == numbers[i + 1]:
            print('edge case found at line ${i}')

print(increased)
