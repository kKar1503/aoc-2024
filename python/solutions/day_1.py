import os
from typing import Counter


class Solution:
    day = 1

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
        numbers = [
            (int(nums[0]), int(nums[1]))
            for nums in (
                line.split("   ") for line in input_data.splitlines(keepends=False)
            )
        ]
        left = [nums[0] for nums in numbers]
        right = [nums[1] for nums in numbers]
        left.sort()
        right.sort()
        combined = zip(left, right)
        total = sum([abs(right - left) for left, right in combined])

        self.write_output(1, str(total))

    def part2(self):
        # Implement part 2 logic here
        input_data = self.read_input()
        numbers = [
            (int(nums[0]), int(nums[1]))
            for nums in (
                line.split("   ") for line in input_data.splitlines(keepends=False)
            )
        ]
        count = Counter([nums[1] for nums in numbers])
        score = 0
        for left, _ in numbers:
            score += left * count[left]

        self.write_output(2, str(score))
