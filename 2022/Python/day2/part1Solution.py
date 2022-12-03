from .game import parse_option, who_wins

def main(lines):
    score = 0
    for line in lines:
        first_raw, second_raw = line.split(' ')
        first = parse_option(first_raw)
        second = parse_option(second_raw)
        score += second.value

        results = who_wins(first, second)
        if results == 0:
            score += 3
        elif results == -1:
            score += 6
    # do work here
    print(score)
    print("solution is day1 .....")
