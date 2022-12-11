from collections import defaultdict


class Monkey:
    def __init__(self, monkey_lines):
        identifier_parts = monkey_lines[0].replace(':', '').split(' ')
        self.identifier = int(identifier_parts[1])

        stripped_items = monkey_lines[1].replace(',', '').replace('  ', '').split(' ')
        # assuming that each monkey has some items
        self.items = [int(item) for item in stripped_items[2:]]

        split_operation = monkey_lines[2].split(' ')
        self.operation = split_operation[-2]
        self.operation_arg = split_operation[-1]

        self.divisibility = int(monkey_lines[3].split(' ')[-1])

        self.monkey_if_true = int(monkey_lines[4].split(' ')[-1])
        self.monkey_if_false = int(monkey_lines[5].split(' ')[-1])

        # variable to count the amount of times this monkey has inspected
        self.inspected = 0

        # print("init monkey")
        # print(self.identifier, self.items, self.operation, self.operation_arg, self.divisibility, self.monkey_if_true, self.monkey_if_false)

    def increase_worry(self, item):
        operation_value = item
        if self.operation_arg != "old":
            operation_value = int(self.operation_arg)

        if self.operation == "+":
            return item + operation_value
        elif self.operation == "*":
            return item * operation_value
        else:
            print("unsupported operation", self.operation)
            exit()

    def do_turn(self):
        item_locations = defaultdict(lambda: [])
        for item in self.items:
            self.inspected += 1
            # print(self.identifier, item)
            new_worry = self.increase_worry(item) // 3
            test_value = new_worry % self.divisibility
            if test_value == 0:
                item_locations[self.monkey_if_true].append(new_worry)
            else:
                item_locations[self.monkey_if_false].append(new_worry)

        self.items = []

        return item_locations

    def __repr__(self):
        resp = "Monkey: " + str(self.identifier)
        resp += " items:" + str(self.items)
        return resp
