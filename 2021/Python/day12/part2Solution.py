class Node:
    def __init__(self, key):
        self.key = key
        self.is_lower = key.islower()
        self.is_start = key == "start"
        self.is_end = key == "end"
        self.count = 0

        self.edges = []

    def add_edge(self, node):
        self.edges.append(node)


def path_hunt(start, graph, seen_small_cave_twice=False):
    paths = 0
    for edge in start.edges:
        if edge.is_end:
            paths += 1
        elif not edge.is_lower:
            edge.count += 1
            paths += path_hunt(edge, graph, seen_small_cave_twice)
            edge.count -= 1
        else:
            if edge.count == 1 and not seen_small_cave_twice:
                edge.count += 1
                paths += path_hunt(edge, graph, True)
                edge.count -= 1
            elif edge.count == 0:
                edge.count += 1
                paths += path_hunt(edge, graph, seen_small_cave_twice)
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
        if not graph[first].is_start:
            graph[second].add_edge(graph[first])

    start = graph["start"]
    start.count = 1
    paths = path_hunt(start, graph)

    print("Paths", paths)
    # 122880
