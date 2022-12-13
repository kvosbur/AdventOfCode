from typing import List


class Edge:

    def __init__(self, letter, x, y):
        self.visited = False
        self.step_count = -1
        self.letter = letter
        self.x = x
        self.y = y

    def __repr__(self):
        return f"{self.letter} y:{self.y} x:{self.x}"

    def letter_value(self):
        if self.letter == 'S':
            return 0
        elif self.letter == 'E':
            return 26
        else:
            return ord(self.letter) - ord('a')

    def get_neighbors(self, graph):
        to_visit = []
        # check left
        if self.x > 0 and not graph[self.y][self.x - 1].visited:
            node = graph[self.y][self.x - 1]
            if node.letter_value() - self.letter_value() <= 1:
                node.step_count = self.step_count + 1
                node.visited = True
                to_visit.append(node)

        # check right
        if self.x < len(graph[self.y]) - 1 and not graph[self.y][self.x + 1].visited:
            node = graph[self.y][self.x + 1]
            if node.letter_value() - self.letter_value() <= 1:
                node.step_count = self.step_count + 1
                node.visited = True
                to_visit.append(node)

        # check up
        if self.y > 0 and not graph[self.y - 1][self.x].visited:
            node = graph[self.y - 1][self.x]
            if node.letter_value() - self.letter_value() <= 1:
                node.step_count = self.step_count + 1
                node.visited = True
                to_visit.append(node)

        # check down
        if self.y < len(graph) - 1 and not graph[self.y + 1][self.x].visited:
            node = graph[self.y + 1][self.x]
            if node.letter_value() - self.letter_value() <= 1:
                node.step_count = self.step_count + 1
                node.visited = True
                to_visit.append(node)

        return to_visit
