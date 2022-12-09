class Tree:
    def __init__(self, size, is_visible):
        self.size = size
        self.is_visible = is_visible

    def __repr__(self):
        return f"{self.size}:{'T' if self.is_visible else 'F'}"


def mark_edge_trees(forrest):
    for x in range(len(forrest[0])):
        forrest[0][x].is_visible = True
        forrest[-1][x].is_visible = True

    for y in range(len(forrest)):
        forrest[y][0].is_visible = True
        forrest[y][-1].is_visible = True


def is_tree_visible(forrest, x, y):
    is_visible = False
    this_tree = forrest[y][x]

    # check left
    temp_visible = True
    for current_x in range(0, x):
        tree = forrest[y][current_x]
        if tree.size >= this_tree.size:
            temp_visible = False
    is_visible = is_visible or temp_visible

    # check right
    temp_visible = True
    for current_x in range(x + 1, len(forrest[0])):
        tree = forrest[y][current_x]
        if tree.size >= this_tree.size:
            temp_visible = False
    is_visible = is_visible or temp_visible

    # check up
    temp_visible = True
    for current_y in range(0, y):
        tree = forrest[current_y][x]
        if tree.size >= this_tree.size:
            temp_visible = False
    is_visible = is_visible or temp_visible

    # check down
    temp_visible = True
    for current_y in range(y + 1, len(forrest)):
        tree = forrest[current_y][x]
        if tree.size >= this_tree.size:
            temp_visible = False
    is_visible = is_visible or temp_visible

    return is_visible


def mark_trees(forrest):
    for x in range(0, len(forrest[0])):
        for y in range(0, len(forrest)):
            forrest[y][x].is_visible = is_tree_visible(forrest, x, y)


def count_visible_trees(forrest):
    visible = 0
    for x in range(0, len(forrest[0])):
        for y in range(0, len(forrest)):
            if forrest[y][x].is_visible:
                visible += 1
    return visible


def main(lines):
    forrest = []
    for line in lines:
        raw_trees = [*line]
        row = [Tree(int(tree), False) for tree in raw_trees]
        forrest.append(row)

    mark_edge_trees(forrest)
    mark_trees(forrest)
    print(count_visible_trees(forrest))
