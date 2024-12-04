#include "../../../include/runner/file_utils.h"
#include "../../../include/runner/helper_utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int check_mas_or_reverse(char ch1, char ch2, char ch3) {
  bool is_mas = ch1 == 'M' && ch2 == 'A' && ch3 == 'S';
  bool is_sam = ch1 == 'S' && ch2 == 'A' && ch3 == 'M';
  return is_mas || is_sam ? 1 : 0;
}

int count_mas_in_x(char **lines, int x, int y, size_t width, size_t height) {
  if (x - 1 < 0 || x + 1 >= width || y - 1 < 0 || y + 1 >= height) {
    return 0;
  }

  return check_mas_or_reverse(lines[y - 1][x - 1], lines[y][x],
                              lines[y + 1][x + 1]) *
         check_mas_or_reverse(lines[y - 1][x + 1], lines[y][x],
                              lines[y + 1][x - 1]);
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
      if (lines[y][x] == 'A') {
        count += count_mas_in_x(lines, x, y, width, height);
      }
    }
  }
  for (int i = 0; i < height; i++) {
    free(lines[i]);
  }
  free(lines);

  char buf[100];
  sprintf(buf, "%d", count);
  write_to_file("outputs/day4/output-part2", buf);

  return 0;
}
