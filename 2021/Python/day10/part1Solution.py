opening = "([{<"
closing = ")]}>"
pairs = ["()", "[]", "{}", "<>"]
points = {')': 3, ']': 57, '}': 1197, '>': 25137}


def do_line(characters):
    stack = []
    for char in characters:
        if char in opening:
            stack.append(char)
        else:
            beginning = stack.pop()
            for pair in pairs:
                if char in pair and beginning not in pair:
                    return char
    return None


def main(lines):
    # do work here

    score = 0
    for line in lines:
        result = do_line(line)
        if result is not None:
            score += points[result]

    print("Score", score)
