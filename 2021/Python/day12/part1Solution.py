def make_base_counts(graph, existing_counts=None):
    base = {}
    for key in graph.keys():
        if existing_counts is None:
            base[key] = 0
        else:
            base[key] = existing_counts[key]
    return base


def check_visit_count_okay(node, node_counts):
    if node.islower() and node_counts[node] > 0:
        return False
    return True


def path_hunt(start, graph, node_counts, count = 0):
    if start == "end":
        return 1

    paths = 0
    for edge in graph[start]:
        if check_visit_count_okay(edge, node_counts):
            new_node_counts = make_base_counts(graph, node_counts)
            new_node_counts[edge] += 1
            paths += path_hunt(edge, graph, new_node_counts, count + 1)

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
