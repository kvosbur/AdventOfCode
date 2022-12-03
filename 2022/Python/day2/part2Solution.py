from .game import parse_option, who_wins, all_options


def figure_out_ma_hand(other_hand, goal):
    for option in all_options:
        if who_wins(other_hand, option) == goal:
            return option


def main(lines):
    score = 0
    for line in lines:
        first_raw, second_raw = line.split(' ')
        first = parse_option(first_raw)

        # add points based on round status
        if second_raw == 'Y':
            score += 3
        elif second_raw == 'Z':
            score += 6

        # add points based on my hand
        goal = 1
        if second_raw == 'Y':
            goal = 0
        elif second_raw == 'Z':
            goal = -1
        ma_hand = figure_out_ma_hand(first, goal)
        score += ma_hand.value

    print("solution is day2", score)
