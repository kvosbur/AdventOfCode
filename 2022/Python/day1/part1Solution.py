

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


    print("solution for day1 is.....", max(elves))
