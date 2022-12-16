from .valve import Valve, Edge
from typing import List, Dict


def remove_unnecessary_edges(valve, removed):
    node = valve.identifier
    new_edges = []
    for edge in valve.edges:
        if edge.destination != node and edge.destination != removed:
            new_edges.append(edge)

    valve.edges = new_edges


def update_sources(valves: dict, dest_node, current_valve: Valve):
    sources = []
    for identifier, valve in valves.items():
        new_edges = []
        for edge in valve.edges:
            if edge.destination == dest_node:
                for current_edge in current_valve.edges:
                    new_edges.append(Edge(current_edge.destination, current_edge.weight + edge.weight))

        valve.edges += new_edges
        remove_unnecessary_edges(valve, dest_node)

    del valves[dest_node]

    return sources


def floyd_warshall(valves: Dict[str, Valve]):
    keys = list(valves.keys())
    n = len(keys)
    # initialize adjacency matrix
    path = [[None for x in range(n)] for y in range(n)]

    # use known edges to put in adjacency matrix
    for v in range(n):
        valve = valves[keys[v]]
        for u in range(n):
            dest = keys[u]
            if v == u:
                path[v][u] = 0
            elif dest in valve.edges:
                edge = valve.edges.index(dest)
                path[v][u] = valve.edges[edge].weight

    # run Floyd-Warshall
    for k in range(n):
        for i in range(n):
            for j in range(n):
                if path[i][k] is not None and path[k][j] is not None \
                        and (path[i][j] is None or path[i][k] + path[k][j] < path[i][j]):
                    path[i][j] = path[i][k] + path[k][j]

    # carry over results into valves
    for v_index, v in enumerate(keys):
        edges = []
        for u_index, u in enumerate(keys):
            if u != v:
                edges.append(Edge(u, path[v_index][u_index]))
        valves[v].edges = edges


def is_node_turned_on_in_visit(node: str, visited: set):
    return node in visited


def brute_force(current: Valve, valves: Dict[str, Valve], visited: List[str], flow: int, time_remaining: int):
    if time_remaining <= 0:
        return [[flow, visited]]

    results = []
    for edge in current.edges:
        if edge.destination in visited:
            continue
        next_valve = valves[edge.destination]
        new_time_after_open = time_remaining - edge.weight - 1
        new_flow = flow + (next_valve.flow * new_time_after_open)

        if new_time_after_open < 0:
            continue

        results += brute_force(next_valve, valves, visited + [edge.destination], new_flow, new_time_after_open)

    if len(results) == 0:
        return [[flow, visited]]

    best = results[0]
    for result in results[1:]:
        if result[0] > best[0]:
            best = result
    return [best]


def main(lines):
    # parse input
    valves = {}
    for line in lines:
        valve = Valve(line)
        valves[valve.identifier] = valve

    # shrink graph
    valves_copy = valves.copy()
    for identifier, valve in valves_copy.items():
        # print(identifier, valve)
        if identifier == 'AA':
            continue

        if valve.flow == 0:
            # found node to remove
            update_sources(valves, identifier, valve)

    floyd_warshall(valves)

    print("after")
    for valve in valves.values():
        print(valve)

    # BRUTE FORCE!!!!!!!
    res = brute_force(valves["AA"], valves, ["AA"], 0, 30)
    print(res)


