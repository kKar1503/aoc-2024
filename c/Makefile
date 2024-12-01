CC = clang
CFLAGS = -I../../include/runner
LDFLAGS =

# Default target
all:
	@echo "Usage: make run/dayX/partY (e.g., make run/day1/part1 or make run/day10/part2)"
	@echo "       make read/dayX/partY (e.g., make read/day1/part1 or make read/day10/part2)"

# Dynamic rule for any day/part
run/%:
	@day=$(shell dirname $*); \
	part=$(shell basename $*); \
	solution_dir=src/solutions/$$day; \
	input_file=$$solution_dir/input.txt; \
	$(CC) $(CFLAGS) $$solution_dir/$$part.c src/runner/helper_utils.c src/runner/file_utils.c -o $$solution_dir/$$part.out; \
	echo "Running $$part.out..."; \
	$$solution_dir/$$part.out

read/%:
	@day=$(shell dirname $*); \
	part=$(shell basename $*); \
	solution_dir=src/solutions/$$day; \
	output_file=$$solution_dir/output-$$part; \
	echo "Reading $$output_file..."; \
	cat $$output_file
