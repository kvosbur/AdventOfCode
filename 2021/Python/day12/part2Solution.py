special_nodes = ["start", "end"]


def make_base_counts(graph, existing_counts=None):
    base = {}
    for key in graph.keys():
        if existing_counts is None:
            base[key] = 0
        else:
            base[key] = existing_counts[key]
    return base


def has_visited_small_cave_twice(node_counts):
    for key in node_counts.keys():
        if key.islower() and node_counts[key] == 2:
            return True


def check_visit_count_okay(node, node_counts):
    if node.islower():
        if node_counts[node] == 1 and not has_visited_small_cave_twice(node_counts) and node not in special_nodes:
            return True
        if node_counts[node] > 0:
            return False
    return True


def path_hunt(start, graph, node_counts):
    if start == "end":
        return 1

    paths = 0
    for edge in graph[start]:
        if check_visit_count_okay(edge, node_counts):
            new_node_counts = make_base_counts(graph, node_counts)
            new_node_counts[edge] += 1
            paths += path_hunt(edge, graph, new_node_counts)
    return paths


def main(lines):
    # do work here
    graph = {}
    for line in lines:
        first, second = line.split('-')
        if graph.get(first) is None:
            graph[first] = []
        if graph.get(second) is None:
            graph[second] = []

        graph[first].append(second)
        graph[second].append(first)

    node_counts = make_base_counts(graph)
    node_counts["start"] = 1
    paths = path_hunt("start", graph, node_counts)

    print("Paths", paths)
