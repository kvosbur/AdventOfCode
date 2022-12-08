from .shared import File, Directory


def update_current(root, cwd):
    result = root
    for path in cwd:
        result = result.files[path]

    return result


def generate_folder_structure(lines):
    cwd = []
    root = Directory("/")
    current = root

    for line in lines:
        parts = line.split(' ')
        if parts[0] == "$":
            # process command
            if parts[1] == "ls":
                # it is an ls command
                pass
            elif parts[1] == "cd":
                if parts[2] == "/":
                    current = root
                    cwd = []
                elif parts[2] == "..":
                    cwd.pop()
                    current = update_current(root, cwd)
                else:
                    cwd.append(parts[2])
                    current = update_current(root, cwd)
        elif parts[0] == "dir":
            # add new directory
            current.add_file(Directory(parts[1]))
        else:
            # add new file
            current.add_file(File(parts[1], int(parts[0])))
    return root


def get_small_folder_sum(root: Directory):
    sum = 0
    if root.size() <= 100000:
        sum += root.size()

    for file in root.files.values():
        if isinstance(file, Directory):
            sum += get_small_folder_sum(file)

    return sum

def main(lines):
    root = generate_folder_structure(lines)

    # do work here
    print("solution is day1 .....")
    print(get_small_folder_sum(root))
    # root.output()
