from string import ascii_lowercase, ascii_uppercase

def score_char(char):
    if char in ascii_lowercase:
        return ord(char) - ord('a') + 1
    elif char in ascii_uppercase:
        return ord(char) - ord('A') + 27


def get_same_char(first, second):
    return list(set(first) & set(second))[0]

def main(lines):
    score = 0
    for line in lines:
        length_of_line = len(line) // 2
        first = line[:length_of_line]
        second = line[length_of_line:]

        same_char = get_same_char(first, second)
        score += score_char(same_char)

    print(score)
    print("solution is day1 .....")
