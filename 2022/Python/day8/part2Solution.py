class Tree:
    def __init__(self, size, scenic_score):
        self.size = size
        self.scenic_score = scenic_score

    def __repr__(self):
        return f"{self.size}:({self.scenic_score})"


def how_scenic(forrest, x, y):
    current_score = 1
    this_tree = forrest[y][x]

    # check left
    count = 0
    for current_x in range(x - 1, -1, -1):
        count += 1
        tree = forrest[y][current_x]
        if tree.size >= this_tree.size:
            break
    current_score = current_score * count

    # check right
    count = 0
    for current_x in range(x + 1, len(forrest[0])):
        count += 1
        tree = forrest[y][current_x]
        if tree.size >= this_tree.size:
            break
    current_score = current_score * count

    # check up
    count = 0
    for current_y in range(y - 1, -1, -1):
        count += 1
        tree = forrest[current_y][x]
        if tree.size >= this_tree.size:
            break
    current_score = current_score * count

    # check down
    count = 0
    for current_y in range(y + 1, len(forrest)):

        count += 1
        tree = forrest[current_y][x]
        if tree.size >= this_tree.size:
            break
    current_score = current_score * count

    return current_score


def mark_trees(forrest):
    for x in range(0, len(forrest[0])):
        for y in range(0, len(forrest)):
            forrest[y][x].scenic_score = how_scenic(forrest, x, y)


def best_scenic_score(forrest):
    best = 0
    for x in range(0, len(forrest[0])):
        for y in range(0, len(forrest)):
            if forrest[y][x].scenic_score > best:
                best = forrest[y][x].scenic_score
    return best


def main(lines):
    forrest = []
    for line in lines:
        raw_trees = [*line]
        row = [Tree(int(tree), False) for tree in raw_trees]
        forrest.append(row)

    mark_trees(forrest)
    print(best_scenic_score(forrest))
