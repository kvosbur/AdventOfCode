from .graph import Edge
from pprint import pprint
from typing import List


def debug_graph(graph: List[List[Edge]]):
    for row in graph:
        for node in row:
            print(node.letter if node.visited else '.', end="")
        print("")


def main(lines):
    start = None

    graph = []
    for y, line in enumerate(lines):
        graph_row = []
        for x, char in enumerate([*line]):
            edge = Edge(char, x, y)
            if char == 'S':
                start = edge
            graph_row.append(edge)
        graph.append(graph_row)

    start.visited = True
    start.step_count = 0
    to_visit = start.get_neighbors(graph)
    while len(to_visit) != 0:
        current = to_visit.pop(0)

        print(current)
        if current.letter == 'E':
            print("found goal at ", current.step_count)
            break

        to_visit += current.get_neighbors(graph)

    debug_graph(graph)
