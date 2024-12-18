#include "../../../include/runner/file_utils.h"
#include "../../../include/runner/helper_utils.h"
#include <stdio.h>
#include <stdlib.h>

int compare_ints(const void *a, const void *b) {
  return (*(int *)a - *(int *)b);
}

int main() {
  char *input = read_file_content("inputs/day1");
  if (input == NULL) {
    return 1;
  }

  size_t count;
  char **lines = str_split(input, '\n', &count);
  int sum = 0;
  int *left = malloc(sizeof(int) * count);
  int *right = malloc(sizeof(int) * count);

  if (lines) {
    int i;
    for (i = 0; *(lines + i); i++) {
      int num1, num2;
      if (sscanf(*(lines + i), "%d   %d", &num1, &num2) == 2) {
        *(left + i) = num1;
        *(right + i) = num2;
      } else {
        fprintf(stderr, "Failed to parse line %d: %s\n", i, *(lines + i));
        exit(1);
      }
      free(*(lines + i));
    }
    free(lines);
  }

  qsort(left, count, sizeof(int), compare_ints);
  qsort(right, count, sizeof(int), compare_ints);

  for (int i = 0; i < count; i++) {
    sum += abs(right[i] - left[i]);
  }

  char buf[100];
  sprintf(buf, "%d", sum);
  write_to_file("outputs/day1/output-part1", buf);

  free(input);
  return 0;
}
