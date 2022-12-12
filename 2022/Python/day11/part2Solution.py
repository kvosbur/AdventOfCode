from .monkey import Monkey


def spread_around_items(monkeys, monkey_results):
    for key, value in monkey_results.items():
        monkeys[key].items += value
        # print("passing:", value, "to:", key)


def main(lines):
    rounds = 10000

    monkeys = []
    current_monkey_lines = []
    for line in lines:
        if line != "":
            current_monkey_lines.append(line)
        else:
            monkeys.append(Monkey(current_monkey_lines))
            current_monkey_lines = []

    monkeys.append(Monkey(current_monkey_lines))

    for i in range(rounds):
        for monkey in monkeys:
            # print("doing work for monkey:", monkey.identifier)
            results = monkey.do_turn(is_worried=False)
            spread_around_items(monkeys, results)

        # print("\nresults for round:", i + 1)
        # for monkey in monkeys:
        #     print(monkey)
        # for monkey in monkeys:
        #     print(monkey.identifier, monkey.inspected)

        # if i == 19:
        #     print("\nresults for round:", i + 1)
        #     for monkey in monkeys:
        #         print(monkey.identifier, monkey.inspected)
        if (i + 1) % 1000 == 0:
            print("\nresults for round:", i + 1)
            for monkey in monkeys:
                print(monkey.identifier, monkey.inspected)

    monkeys.sort(key= lambda monkey: monkey.inspected)

    print("")
    for monkey in monkeys:
        print(monkey.identifier, monkey.inspected)

    monkey_business = monkeys[-2].inspected * monkeys[-1].inspected
    print("monkey business:", monkey_business)

    # do work here
    print("solution is day1 .....")
