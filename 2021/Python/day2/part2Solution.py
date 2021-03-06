
def main(lines):
    horizontal = 0
    depth = 0
    aim = 0

    for line in lines:
        direction, move = line.split(' ')
        moveInt = int(move)
        if direction == "forward":
            horizontal += moveInt
            depth += (aim * moveInt)
        elif direction == "up":
            aim -= moveInt
        elif direction == "down":
            aim += moveInt
        else:
            print("error occurred for:", line)

    print(horizontal)
    print(depth)
    print(horizontal * depth)
