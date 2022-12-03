from enum import Enum

class Option(Enum):
    Rock = 1
    Paper = 2
    Scissors = 3


def parse_option(text: str):
    if text in ['A', 'X']:
        return Option.Rock
    if text in ['B', 'Y']:
        return Option.Paper
    if text in ['C', 'Z']:
        return Option.Scissors

"""
-1 - second wins
0 - tie
1 - first wins
"""
def who_wins(first: Option, second: Option):
    # tie
    if first == second:
        return 0

    # first is "higher"
    if first.value > second.value:
        if first == Option.Scissors and second == Option.Rock:
            return -1
        else:
            return 1

    # second value is "higher"
    if second == Option.Scissors and first == Option.Rock:
        return 1
    return -1
