#include "../../../include/runner/file_utils.h"
#include "../../../include/runner/helper_utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int sign(int x) { return x > 0 ? 1 : x < 0 ? -1 : 0; }

bool is_safe(char **levels, size_t level_count) {
  int prev = atoi(levels[0]);
  int first = atoi(levels[1]);
  int direction = first - prev;
  if (direction == 0 || abs(direction) > 3) {
    return false;
  }
  prev = first;
  for (int j = 2; *(levels + j); j++) {
    int current = atoi(levels[j]);
    int delta = current - prev;
    if (delta == 0 || abs(delta) > 3) {
      return false;
    }
    if (sign(delta) != sign(direction)) {
      return false;
    }
    prev = current;
  }
  return true;
}

int main() {
  char *input = read_file_content("inputs/day2");
  if (input == NULL) {
    return 1;
  }

  size_t count;
  char **lines = str_split(input, '\n', &count);
  int safe = 0;

  if (lines) {
    for (int i = 0; i < count; i++) {
      size_t level_count;
      char **levels = str_split(lines[i], ' ', &level_count);
      if (levels) {
        if (is_safe(levels, level_count)) {
          safe++;
        }
      }
      for (int j = 0; j < level_count; j++) {
        free(levels[j]);
      }
      free(levels);
      free(lines[i]);
    }
    free(lines);
  }

  char buf[100];
  sprintf(buf, "%d", safe);
  write_to_file("outputs/day2/output-part1", buf);

  free(input);
  return 0;
}
