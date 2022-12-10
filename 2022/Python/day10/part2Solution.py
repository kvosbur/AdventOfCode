from .computer import Instruction, CPU, Register, Screen
from typing import Callable, Dict, List


def noop(registers):
    pass


def add_to_register(registers: Dict[str, Register], register_name, value):
    registers[register_name].value += value


def addx(value):
    return lambda registers: add_to_register(registers, 'x', value)


def main(lines):
    instructions = []
    for line in lines:
        instruction_parts = line.split(' ')
        if line == "noop":
            instructions.append(Instruction(1, noop))
        elif instruction_parts[0] == "addx":
            value = int(instruction_parts[1])
            instructions.append(Instruction(2, addx(value)))

    registers = {'x': Register('x')}
    cpu = CPU(registers, instructions)
    screen = Screen(registers)

    while not cpu.is_finished():
        screen.draw_next()
        cpu.increment_cycle()

    print("\n".join(screen.rows))
