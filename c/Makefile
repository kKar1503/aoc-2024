CC = clang
CFLAGS = -I../../include/runner
LDFLAGS =
BIN_DIR = bin
OUTPUT_DIR = outputs

# Default target
all:
	@echo "Usage: make run/dayX/partY (e.g., make run/day1/part1 or make run/day10/part2)"
	@echo "       make read/dayX/partY (e.g., make read/day1/part1 or make read/day10/part2)"

# Dynamic rule for any day/part
build/%:
	@day=$(shell dirname $*); \
	part=$(shell basename $*); \
	solution_dir=src/solutions/$$day; \
	if [ ! -f inputs/$$day ]; then \
		echo "Error: inputs/$$day not found"; \
		exit 1; \
	fi; \
	mkdir -p $(BIN_DIR); \
	mkdir -p $(BIN_DIR)/$$day; \
	echo "Building $(BIN_DIR)/$$day/$$part.out..."; \
	$(CC) $(CFLAGS) $$solution_dir/$$part.c src/runner/helper_utils.c src/runner/file_utils.c -o $(BIN_DIR)/$$day/$$part.out

run/%: build/%
	@day=$(shell dirname $*); \
	part=$(shell basename $*); \
	mkdir -p $(OUTPUT_DIR); \
	mkdir -p $(OUTPUT_DIR)/$$day; \
	bin_file=$(BIN_DIR)/$$day/$$part.out; \
	echo "Running $$bin_file..."; \
	$$bin_file

read/%: run/%
	@day=$(shell dirname $*); \
	part=$(shell basename $*); \
	output_file=$(OUTPUT_DIR)/$$day/output-$$part; \
	echo "Reading $$output_file..."; \
	cat $$output_file

# Clean up
clean:
	rm -rf $(BIN_DIR) $(OUTPUT_DIR)
