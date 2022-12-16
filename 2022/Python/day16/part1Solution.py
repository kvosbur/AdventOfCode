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


def is_node_turned_on_in_visit(node: Valve, visited: List[List]):
    for visit in visited:
        if node.identifier == visit[0] and visit[1]:
            return True
    return False


def brute_force(current: Valve, valves: Dict[str, Valve], visited: List[List], flow: int, time_remaining: int):
    if time_remaining <= 0:
        return [[flow, visited]]

    results = []
    for edge in current.edges:
        next_valve = valves[edge.destination]
        new_time_after_travel = time_remaining - edge.weight

        results += brute_force(next_valve, valves, visited + [[edge.destination, False]], flow, new_time_after_travel)
        if not is_node_turned_on_in_visit(next_valve, visited):
            new_time_after_open = new_time_after_travel - 1
            new_flow = flow + (next_valve.flow * new_time_after_open)
            results += brute_force(next_valve, valves, visited + [[edge.destination, True]], new_flow, new_time_after_open)

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

    # BRUTE FORCE!!!!!!!
    res = brute_force(valves["AA"], valves, [["AA", False]], 0, 16)
    print(res)

    for valve in valves.values():
        print(valve)


