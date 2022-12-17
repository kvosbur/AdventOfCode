from .valve import Valve, Edge
from typing import List, Dict
from copy import deepcopy


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
        edges.sort(key=lambda x: x.weight)
        valves[v].edges = edges


def is_node_turned_on_in_visit(node: str, visited: set):
    return node in visited


# visited array hold info as [amount to wait, flow_to_add_after_wait, visited_node_list]
def brute_force(current: Valve, valves: Dict[str, Valve], first_visited: List,
                second_visited: List, flow: int, temp_flow: int, time_remaining: int):
    navigator = first_visited
    other_navigator = second_visited
    if second_visited[0] < first_visited[0]:
        navigator = second_visited
        other_navigator = first_visited

    if time_remaining <= 0 or time_remaining - navigator[0] < 0:
        # print("exceeds time")
        return [[flow, first_visited, second_visited]]

    new_time_remaining = time_remaining - navigator[0]
    new_flow = flow + (navigator[1] * new_time_remaining)
    new_temp_flow = temp_flow + navigator[1]
    new_other_navigator = deepcopy(other_navigator)
    new_other_navigator[0] -= navigator[0]

    results = []
    for edge in current.edges:
        # print("debug", flow, temp_flow, new_time_remaining, edge, first_visited, second_visited)
        if edge.destination in navigator[2] or edge.destination in other_navigator[2]:
            # print("rejected")
            continue
        navigator_copy = deepcopy(navigator)
        next_valve = valves[edge.destination]
        navigator_copy[0] = edge.weight
        navigator_copy[1] = next_valve.flow
        navigator_copy[2] = navigator_copy[2] + [edge.destination]

        results += brute_force(next_valve, valves, navigator_copy, new_other_navigator,
                               new_flow, new_temp_flow, new_time_remaining)

    # filter results based on visited list
    # filtered = []
    # for result in results:
    #     first_set = set(result[1][2])
    #     second_set = set(result[2][2])
    #     if not (first_set & second_set):
    #         filtered.append(result)

    if len(results) == 0:
        # print("no results")
        return [[flow, first_visited, second_visited]]

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
    res = brute_force(valves["AA"], valves, [0, 0, ["AA"]], [0, 0, ["AA"]], 0, 0, 26)
    print("best results")
    print(res)
