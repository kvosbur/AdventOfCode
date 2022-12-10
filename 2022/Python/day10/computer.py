from typing import Callable, Dict, List


class Register:
    value: int
    identifier: str

    def __init__(self, identifier: str):
        self.value = 1
        self.identifier = identifier


class Instruction:
    cycles: int

    def __init__(self, cycles: int, action: Callable[[Dict[str, Register]], None]):
        self.cycles = cycles
        self.action = action


class CPU:

    def __init__(self, registers: Dict[str, Register],  instructions: List[Instruction]):
        self.registers = registers
        self.instructions = instructions
        self.current_cycle = 1
        self.current_instruction_index = 0

    def increment_cycle(self) -> int:
        self.current_cycle += 1
        current_instruction = self.instructions[self.current_instruction_index]
        current_instruction.cycles -= 1
        if current_instruction.cycles == 0:
            current_instruction.action(self.registers)
            self.current_instruction_index += 1
        return self.current_cycle

    def is_finished(self) -> bool:
        return len(self.instructions) == self.current_instruction_index


class Screen:

    def __init__(self, registers: Dict[str, Register]):
        self.registers = registers
        self.current_position = 0
        self.current_row = ""
        self.rows = []

    def draw_next(self):
        width_pos = self.current_position % 40
        sprite_pos = self.registers['x'].value
        if sprite_pos - 1 <= width_pos <= sprite_pos + 1:
            self.current_row += "#"
        else:
            self.current_row += "."

        self.current_position += 1
        if self.current_position % 40 == 0:
            self.rows.append(self.current_row)
            self.current_row = ""

