import os
import re


class Solution:
    day = 3

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

    def part1(self):
        # Implement part 1 logic here
        input_data = self.read_input()

        regex = "mul\\((\\d+),(\\d+)\\)"
        results = re.findall(regex, input_data)

        sum = 0
        for result in results:
            sum += int(result[0]) * int(result[1])

        self.write_output(1, str(sum))

    def part2(self):
        # Implement part 2 logic here
        input_data = self.read_input()

        regex = "(mul\\((\\d+),(\\d+)\\)|(do\\(\\))|(don't\\(\\)))"
        results = re.findall(regex, input_data)

        sum = 0
        should_add = True
        for result in results:
            if result[3] != "":
                should_add = True
            elif result[4] != "":
                should_add = False
            elif result[1] != "" and result[2] != "":
                if should_add:
                    sum += int(result[1]) * int(result[2])

        self.write_output(2, str(sum))
