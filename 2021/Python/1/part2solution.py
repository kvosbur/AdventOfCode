increased = 0

with open("part1.txt", "r") as f:
    numbers = [int(n) for n in f.readlines()]
    prev = numbers[0] + numbers[1] + numbers[2]
    for x in zip(numbers[1:], numbers[2:], numbers[3:]):
        newSum = sum(x)
        if newSum > prev:
            increased += 1
        prev = newSum

print(increased)
