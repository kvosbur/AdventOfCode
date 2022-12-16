import re


class Edge:
    def __init__(self, destination, weight=1):
        self.destination = destination
        self.weight = weight

    def __repr__(self):
        return f"dest:{self.destination} weight:{self.weight}"


class Valve:

    def __init__(self, raw_line):
        self.line = raw_line
        self._get_flow_rate()
        self._get_identifier()
        self._get_edges()

    def __repr__(self):
        return f"{self.identifier} flow:{self.flow} edges: {self.edges}"

    def _get_flow_rate(self):
            r = re.compile("rate=([0-9]+);")
            self.flow = int(r.findall(self.line)[0])

    def _get_identifier(self):
        r = re.compile("Valve ([A-Z]+) ")
        self.identifier = r.findall(self.line)[0]

    def _get_edges(self):
        pieces = self.line.replace(',', '').split(' ')
        self.edges = []
        for piece in pieces[::-1]:
            if len(piece) == 2:
                self.edges.append(Edge(piece))
            else:
                break
