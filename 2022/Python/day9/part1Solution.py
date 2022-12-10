class Position:

    def __init__(self):
        self.x = 0
        self.y = 0

    def do_step(self, x, y):
        self.x += x
        self.y += y


def move_head(head: Position, direction):
    if direction == "U":
        head.do_step(0, 1)
    elif direction == "D":
        head.do_step(0, -1)
    elif direction == "R":
        head.do_step(1, 0)
    elif direction == "L":
        head.do_step(-1, 0)


def move_tail(tail: Position, head: Position):
    move_x = 0
    move_y = 0

    x_diff = head.x - tail.x
    y_diff = head.y - tail.y
    if abs(x_diff) == 2 or abs(y_diff) == 2:
        # tail will move
        if x_diff != 0:
            move_x = x_diff // abs(x_diff)
        if y_diff != 0:
            move_y = y_diff // abs(y_diff)

    tail.do_step(move_x, move_y)


def main(lines):
    head = Position()
    tail = Position()
    seen = set()
    seen.add("0,0")
    # print(seen)
    for line in lines:
        direction, str_value = line.split(' ')
        value = int(str_value)
        for iteration in range(value):
            move_head(head, direction)
            move_tail(tail, head)
            tail_position_str = f"{tail.x},{tail.y}"
            # print(direction, tail_position_str, seen)
            seen.add(tail_position_str)

    print(len(seen))
    # print(seen)
