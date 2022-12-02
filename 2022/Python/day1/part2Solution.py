

def main(lines):
    # do work here
    elves = []
    calSum = 0
    for line in lines:
        if line != '':
            calSum += int(line)
        else:
            elves.append(calSum)
            calSum = 0

    elves.sort()

    goal = elves[-1] + elves[-2] + elves[-3]
    print("solution is day2 .....", goal)
