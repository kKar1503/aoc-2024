import os
from typing import Counter


class Solution:
    day = 2

    def __init__(self):
        self.input_file = f"inputs/day{self.day}"
        self.output_file_part1 = f"outputs/day{self.day}/output-part1"
        self.output_file_part2 = f"outputs/day{self.day}/output-part2"

    def read_input(self):
        if not os.path.exists(self.input_file):
            raise FileNotFoundError(f"Input file {self.input_file} not found")

        with open(self.input_file, "r") as f:
            return f.read()

    def write_output(self, part, result):
        if part == 1:
            output_file = self.output_file_part1
        elif part == 2:
            output_file = self.output_file_part2
        else:
            raise ValueError("Invalid part number")

        if not os.path.exists(os.path.dirname(output_file)):
            os.makedirs(os.path.dirname(output_file))
        with open(output_file, "w") as f:
            f.write(str(result))

    def sign(self, n):
        return 1 if n > 0 else -1 if n < 0 else 0

    def is_safe(self, report: list[int]) -> bool:
        prev = report[0]
        direction = report[1] - prev
        if direction == 0:
            return False
        if abs(direction) > 3:
            return False
        prev = report[1]
        for i in range(2, len(report)):
            delta = report[i] - prev
            if delta == 0:
                return False
            if abs(delta) > 3:
                return False
            if self.sign(delta) != self.sign(direction):
                return False
            prev = report[i]
        return True

    def part1(self):
        # Implement part 1 logic here
        input_data = self.read_input()
        reports = [
            list(map(int, line.split(" ")))
            for line in input_data.splitlines(keepends=False)
        ]

        safe = 0
        for report in reports:
            if self.is_safe(report):
                safe += 1

        self.write_output(1, str(safe))

    def part2(self):
        # Implement part 2 logic here
        input_data = self.read_input()
        reports = [
            list(map(int, line.split(" ")))
            for line in input_data.splitlines(keepends=False)
        ]

        safe = 0
        special_safe = 0
        for report in reports:
            if self.is_safe(report):
                safe += 1
                continue

            for i in range(len(report)):
                clone = report.copy()
                clone.pop(i)
                if self.is_safe(clone):
                    safe += 1
                    special_safe += 1
                    break

        self.write_output(2, str(safe))
