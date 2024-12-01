#include "../../../include/runner/file_utils.h"
#include "../../../include/runner/helper_utils.h"
#include <stdio.h>
#include <stdlib.h>

int compare_ints(const void *a, const void *b) {
  return (*(int *)a - *(int *)b);
}

int main() {
  char *input = read_file_content("src/solutions/day1/input");
  if (input == NULL) {
    return 1;
  }

  size_t count;
  char **lines = str_split(input, '\n', &count);
  int sum = 0;
  int *left = malloc(sizeof(int) * count);
  int *right = malloc(sizeof(int) * 100000);

  if (lines) {
    int i;
    for (i = 0; *(lines + i); i++) {
      int num1, num2;
      if (sscanf(*(lines + i), "%d   %d", &num1, &num2) == 2) {
        *(left + i) = num1;
        right[num2]++;
      } else {
        fprintf(stderr, "Failed to parse line %d: %s\n", i, *(lines + i));
        exit(1);
      }
      free(*(lines + i));
    }
    printf("\n");
    free(lines);
  }

  qsort(left, count, sizeof(int), compare_ints);
  qsort(right, count, sizeof(int), compare_ints);

  for (int i = 0; i < count; i++) {
    sum += left[i] * right[left[i]];
  }

  char buf[100];
  sprintf(buf, "%d", sum);
  write_to_file("src/solutions/day1/output-part2", buf);

  free(input);
  return 0;
}
