

def main(lines):
    characters = lines[0]

    # assume that there is a beginning message
    for i in range(len(characters)):
        group = characters[i: i + 4]
        if len(set(group)) == 4:
            print(i + 4)
            exit()
