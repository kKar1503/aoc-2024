import os
import re


class Solution:
    day = 4

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

    def is_xmas(self, s1: str, s2: str, s3: str, s4: str) -> int:
        return 1 if s1 == "X" and s2 == "M" and s3 == "A" and s4 == "S" else 0

    def is_mas_or_reverse(self, s1: str, s2: str, s3: str) -> int:
        return (
            1
            if (s1 == "M" and s2 == "A" and s3 == "S")
            or (s1 == "S" and s2 == "A" and s3 == "M")
            else 0
        )

    def check_xmas(self, input: list[str], coord: tuple[int, int]) -> int:
        count = 0
        x, y = coord
        hasSpaceLeft = x - 3 >= 0
        hasSpaceRight = x + 3 < len(input[y])
        hasSpaceUp = y - 3 >= 0
        hasSpaceDown = y + 3 < len(input)

        # Check left
        if hasSpaceLeft:
            count += self.is_xmas(
                input[y][x], input[y][x - 1], input[y][x - 2], input[y][x - 3]
            )
        # Check right
        if hasSpaceRight:
            count += self.is_xmas(
                input[y][x], input[y][x + 1], input[y][x + 2], input[y][x + 3]
            )

        # Check up
        if hasSpaceUp:
            count += self.is_xmas(
                input[y][x], input[y - 1][x], input[y - 2][x], input[y - 3][x]
            )

        # Check down
        if hasSpaceDown:
            count += self.is_xmas(
                input[y][x], input[y + 1][x], input[y + 2][x], input[y + 3][x]
            )

        # Check diagonal top left
        if hasSpaceLeft and hasSpaceUp:
            count += self.is_xmas(
                input[y][x],
                input[y - 1][x - 1],
                input[y - 2][x - 2],
                input[y - 3][x - 3],
            )

        # Check diagonal top right
        if hasSpaceRight and hasSpaceUp:
            count += self.is_xmas(
                input[y][x],
                input[y - 1][x + 1],
                input[y - 2][x + 2],
                input[y - 3][x + 3],
            )

        # Check diagonal bottom left
        if hasSpaceLeft and hasSpaceDown:
            count += self.is_xmas(
                input[y][x],
                input[y + 1][x - 1],
                input[y + 2][x - 2],
                input[y + 3][x - 3],
            )

        # Check diagonal bottom right
        if hasSpaceRight and hasSpaceDown:
            count += self.is_xmas(
                input[y][x],
                input[y + 1][x + 1],
                input[y + 2][x + 2],
                input[y + 3][x + 3],
            )

        return count

    def check_mas_in_x(self, input: list[str], coord: tuple[int, int]) -> int:
        x, y = coord
        if x - 1 < 0 or x + 1 >= len(input[y]):
            return 0
        if y - 1 < 0 or y + 1 >= len(input):
            return 0

        return self.is_mas_or_reverse(
            input[y - 1][x - 1], input[y][x], input[y + 1][x + 1]
        ) * self.is_mas_or_reverse(
            input[y - 1][x + 1], input[y][x], input[y + 1][x - 1]
        )

    def part1(self):
        # Implement part 1 logic here
        input_data = self.read_input()
        input = input_data.splitlines(keepends=False)

        count = 0
        for y, line in enumerate(input):
            for x, char in enumerate(line):
                if char == "X":
                    count += self.check_xmas(input, (x, y))

        self.write_output(1, str(count))

    def part2(self):
        # Implement part 2 logic here
        input_data = self.read_input()
        input = input_data.splitlines(keepends=False)

        count = 0
        for y, line in enumerate(input):
            for x, char in enumerate(line):
                if char == "A":
                    count += self.check_mas_in_x(input, (x, y))

        self.write_output(2, str(count))
