import os
import re


class Solution:
    day = 5

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

        parts = input_data.split("\n\n")
        part1 = [line.split("|") for line in parts[0].splitlines(keepends=False)]

        after_map = {}
        for n in part1:
            before, after = int(n[0]), int(n[1])
            if before not in after_map:
                after_map[before] = [after]
            else:
                after_map[before].append(after)

        part2 = [line.split(",") for line in parts[1].splitlines(keepends=False)]
        sum = 0
        for num_strs in part2:
            nums = [int(num) for num in num_strs]
            is_valid = True
            for i, num in reversed(list(enumerate(nums))):
                after_this = after_map.get(num)
                if after_this is None:
                    continue

                overlap = set(after_this).intersection(set(nums[:i]))
                if len(overlap) > 0:
                    is_valid = False
                    break

            if is_valid:
                sum += nums[(len(nums) - 1) // 2]

        self.write_output(1, str(sum))

    def part2(self):
        # Implement part 2 logic here
        input_data = self.read_input()

        parts = input_data.split("\n\n")
        part1 = [line.split("|") for line in parts[0].splitlines(keepends=False)]

        after_map = {}
        for n in part1:
            before, after = int(n[0]), int(n[1])
            if before not in after_map:
                after_map[before] = [after]
            else:
                after_map[before].append(after)

        part2 = [line.split(",") for line in parts[1].splitlines(keepends=False)]
        sum = 0
        for num_strs in part2:
            nums = [int(num) for num in num_strs]
            swapped_once = False
            swapped = True
            while swapped:
                swapped = False
                for i in range(1, len(nums)):
                    after_this = after_map.get(nums[i])
                    if after_this is None:
                        continue
                    if nums[i - 1] in after_this:
                        swapped = True
                        swapped_once = True
                        nums[i], nums[i - 1] = nums[i - 1], nums[i]

            if swapped_once:
                print(nums)
                sum += nums[(len(nums) - 1) // 2]

        self.write_output(2, str(sum))
