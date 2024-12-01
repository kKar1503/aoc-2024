"""
All the days of solutions will be exporting a Solution class.
From the Solution class, there always exists part 1 and part 2.
Each part 1 and part 2 are void functions that basically output
the answer into an output file that is in the outputs/ directory.

We use the command line argument to capture the day and part and run
the respective function.
"""

import importlib
import os
import sys


def main():
    if len(sys.argv) != 3:
        print("Usage: python main.py <day> <part>")
        sys.exit(1)

    day = sys.argv[1]
    part = sys.argv[2]

    solution_file = f"solutions/day_{day}.py"
    if not os.path.exists(solution_file):
        print(f"Solution for day {day} does not exist.")
        sys.exit(1)

    module = importlib.import_module(f"solutions.day_{day}")

    solution_class = module.Solution()

    if part == "1":
        solution_class.part1()
    elif part == "2":
        solution_class.part2()
    else:
        print("Invalid part. Use '1' or '2'.")
        sys.exit(1)


if __name__ == "__main__":
    main()
