special_nodes = ["start", "end"]


class Node:
    def __init__(self, key):
        self.is_lower = key.islower()
        self.is_start = key == "start"
        self.is_end = key == "end"
        self.count = 0

        self.edges = []

    def add_edge(self, node):
        self.edges.append(node)


def check_visit_count_okay(node, seen_small_cave_twice):
    if node.is_lower:
        if node.count == 1 and not seen_small_cave_twice and not node.is_start:
            return True, True
        if node.count > 0:
            return False, seen_small_cave_twice
    return True, seen_small_cave_twice


def path_hunt(start, graph, seen_small_cave_twice=False):
    paths = 0
    for edge in start.edges:
        can_continue, new_seen_small_cave_twice = check_visit_count_okay(edge, seen_small_cave_twice)
        if can_continue:
            if edge.is_end:
                paths += 1
            else:
                edge.count += 1
                paths += path_hunt(edge, graph, new_seen_small_cave_twice)
                edge.count -= 1
    return paths


def main(lines):
    # do work here
    graph = {}
    for line in lines:
        first, second = line.split('-')
        if graph.get(first) is None:
            graph[first] = Node(first)
        if graph.get(second) is None:
            graph[second] = Node(second)

        graph[first].add_edge(graph[second])
        graph[second].add_edge(graph[first])

    start = graph["start"]
    start.count = 1
    paths = path_hunt(start, graph)

    print("Paths", paths)
    # 122880
