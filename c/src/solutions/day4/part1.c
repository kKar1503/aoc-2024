#include "../../../include/runner/file_utils.h"
#include "../../../include/runner/helper_utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int check_xmas(char ch1, char ch2, char ch3, char ch4) {
  return ch1 == 'X' && ch2 == 'M' && ch3 == 'A' && ch4 == 'S' ? 1 : 0;
}

int count_xmas(char **lines, int x, int y, size_t width, size_t height) {
  bool has_space_left = x - 3 >= 0;
  bool has_space_right = x + 3 < width;
  bool has_space_up = y - 3 >= 0;
  bool has_space_down = y + 3 < height;

  int count = 0;

  if (has_space_left) {
    count += check_xmas(lines[y][x], lines[y][x - 1], lines[y][x - 2],
                        lines[y][x - 3]);
  }

  if (has_space_right) {
    count += check_xmas(lines[y][x], lines[y][x + 1], lines[y][x + 2],
                        lines[y][x + 3]);
  }

  if (has_space_up) {
    count += check_xmas(lines[y][x], lines[y - 1][x], lines[y - 2][x],
                        lines[y - 3][x]);
  }

  if (has_space_down) {
    count += check_xmas(lines[y][x], lines[y + 1][x], lines[y + 2][x],
                        lines[y + 3][x]);
  }

  if (has_space_left && has_space_up) {
    count += check_xmas(lines[y][x], lines[y - 1][x - 1], lines[y - 2][x - 2],
                        lines[y - 3][x - 3]);
  }

  if (has_space_left && has_space_down) {
    count += check_xmas(lines[y][x], lines[y + 1][x - 1], lines[y + 2][x - 2],
                        lines[y + 3][x - 3]);
  }

  if (has_space_right && has_space_up) {
    count += check_xmas(lines[y][x], lines[y - 1][x + 1], lines[y - 2][x + 2],
                        lines[y - 3][x + 3]);
  }

  if (has_space_right && has_space_down) {
    count += check_xmas(lines[y][x], lines[y + 1][x + 1], lines[y + 2][x + 2],
                        lines[y + 3][x + 3]);
  }

  return count;
}

int main() {
  char *input = read_file_content("inputs/day4");
  if (input == NULL) {
    return 1;
  }

  int count = 0;
  size_t height;
  char **lines = str_split(input, '\n', &height);
  size_t width = strlen(lines[0]);
  free(input);

  for (int y = 0; y < height; y++) {
    for (int x = 0; x < width; x++) {
      if (lines[y][x] == 'X') {
        count += count_xmas(lines, x, y, width, height);
      }
    }
  }
  for (int i = 0; i < height; i++) {
    free(lines[i]);
  }
  free(lines);

  char buf[100];
  sprintf(buf, "%d", count);
  write_to_file("outputs/day4/output-part1", buf);

  return 0;
}
