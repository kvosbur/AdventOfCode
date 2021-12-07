

def main(lines):
    # do work here
    fishies = [int(x) for x in lines[0].split(',')]
    day_due = [0 for x in range(300)]
    for fish in fishies:
        day_due[fish] += 1

    for day in range(256):
        new_fish = day_due[day]

        # set due date of existing fish
        day_due[day + 7] += day_due[day]

        # set due date of new fish
        day_due[day + 9] += new_fish
    print("Total Fishies: ", sum(day_due[256:]))

