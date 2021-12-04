import argparse
import datetime
import importlib
from os import path

parser = argparse.ArgumentParser(description="Run the challenge for a given day and part")
parser.add_argument("day_number", metavar="day", type=int, nargs=1, help="The given day to create the files for")
parser.add_argument("part_number", metavar="part", type=int, nargs=1, choices=[1,2], help="The given day to create the files for")

if __name__ == "__main__":
    args = parser.parse_args()

    day = args.day_number[0]
    target_dir = f"day{day}"
    before = datetime.datetime.now()

    part = args.part_number[0]
    print(f"Running program for day {day} part {part}")
    input_file = path.join(path.dirname(__file__), target_dir, f"day_{day}_input.txt")

    with open(input_file, "r") as f:
        lines = f.readlines()
        lines = [line.replace('\n', '') for line in lines]

        if part == 1:
            module_dir = f"day{day}.part1Solution"
        else:
            module_dir = f"day{day}.part2Solution"

        module = importlib.import_module(module_dir)
        module.main(lines)

    after = datetime.datetime.now()
    total = after - before

    print(f"total time taken: {total}")


