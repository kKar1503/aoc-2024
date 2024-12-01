CC = clang
CFLAGS = -I../../include/runner
LDFLAGS =

# Default target
all:
	@echo "Usage: make dayX/partY (e.g., make day1/part1 or make day10/part2)"

# Dynamic rule for any day/part
%:
	@day=$(shell dirname $@); \
	part=$(shell basename $@); \
	solution_dir=src/solutions/$$day; \
	input_file=$$solution_dir/input.txt; \
	$(CC) $(CFLAGS) $$solution_dir/$$part.c src/runner/helper_utils.c src/runner/file_utils.c -o $$solution_dir/$$part.out; \
	echo "Running $$part.out..."; \
	$$solution_dir/$$part.out
