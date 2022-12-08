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


def get_folder_sizes(root: Directory):
    folders = [root]

    for file in root.files.values():
        if isinstance(file, Directory):
            folders += get_folder_sizes(file)

    return folders


def main(lines):
    total_allowed = 70000000
    goal = 30000000
    root = generate_folder_structure(lines)

    space_left = total_allowed - root.size()
    target_space = goal - space_left

    all_folders = get_folder_sizes(root)
    all_folders.sort(key=lambda x: x.size())

    for folder in all_folders:
        if folder.size() > target_space:
            print("answer", folder.size())
            break

