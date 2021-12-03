
def main(lines):
    horizontal = 0
    depth = 0

    for line in lines:
        direction, move = line.split(' ')
        moveInt = int(move)
        if direction == "forward":
            horizontal += moveInt
        elif direction == "up":
            depth -= moveInt
        elif direction == "down":
            depth += moveInt
        else:
            print("error occurred for:", line)

    print(horizontal)
    print(depth)
    print(horizontal * depth)
