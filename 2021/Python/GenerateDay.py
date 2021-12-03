import argparse
from os import path, rename
import shutil

parser = argparse.ArgumentParser(description="generate the base files for a given day")
parser.add_argument("day_number", metavar="day", type=int, nargs=1, help="The given day to create the files for")

if __name__ == "__main__":
    args = parser.parse_args()

    base_path = path.abspath(path.dirname(__file__))

    day = args.day_number[0]
    target_dir = path.join(base_path, f"day{day}")
    src_dir = path.join(base_path, "base")

    shutil.copytree(src_dir, target_dir)

    src_input = path.join(target_dir, "day_x_input.txt")
    dest_input = path.join(target_dir, f"day_{day}_input.txt")
    rename(src_input, dest_input)
