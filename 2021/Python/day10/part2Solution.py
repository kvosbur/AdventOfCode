opening = "([{<"
closing = ")]}>"
pairs = ["()", "[]", "{}", "<>"]
points = {')': 1, ']': 2, '}': 3, '>': 4}


def do_line(characters):
    stack = []
    for char in characters:
        if char in opening:
            stack.append(char)
        else:
            beginning = stack.pop()
            for pair in pairs:
                if char in pair and beginning not in pair:
                    return char, stack
    # print(stack)
    return None, stack


def repair_line(stack):
    fix = []
    while len(stack) > 0:
        beginning = stack.pop()
        for pair in pairs:
            if beginning in pair:
                fix.append(pair[1])
    return fix


def score_fix(fix):
    score = 0
    for char in fix:
        score = (score * 5) + points[char]
    return score


def main(lines):
    # do work here

    scores = []
    for line in lines:
        result, stack = do_line(line)
        if result is None:
            fix = repair_line(stack)
            inter = score_fix(fix)
            scores.append(inter)

    scores = sorted(scores)
    middle = scores[len(scores) // 2]
    print("Middle value", middle)

