
def main(lines):
    increased = 0

    numbers = [int(n) for n in lines]
    prev = numbers[0] + numbers[1] + numbers[2]
    for x in zip(numbers[1:], numbers[2:], numbers[3:]):
        newSum = sum(x)
        if newSum > prev:
            increased += 1
        prev = newSum

    print(increased)
