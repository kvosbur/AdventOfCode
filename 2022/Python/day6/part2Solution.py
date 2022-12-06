

def main(lines):
    characters = lines[0]

    # assume that there is a beginning message
    for i in range(len(characters)):
        group = characters[i: i + 14]
        if len(set(group)) == 14:
            print(i + 14)
            exit()
