#include "../../../include/runner/file_utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct MulData {
  int num1;
  int num2;
  char *ending_ptr;
} MulData;

MulData *new_mul_data(int num1, int num2, char *ending_ptr) {
  MulData *data = malloc(sizeof(MulData));
  data->num1 = num1;
  data->num2 = num2;
  data->ending_ptr = ending_ptr;
  return data;
}

typedef struct NumData {
  int num;
  char *ending_ptr;
} NumData;

NumData *new_num_data(int num, char *ending_ptr) {
  NumData *data = malloc(sizeof(NumData));
  data->num = num;
  data->ending_ptr = ending_ptr;
  return data;
}

NumData *parse_num(char *input) {
  size_t len = 0;
  while (input[len] != '\0' && input[len] >= '0' && input[len] <= '9') {
    len++;
  }

  char num_str[len + 1];
  memcpy(num_str, input, len);
  num_str[len] = '\0';
  int num = atoi(num_str);

  return new_num_data(num, input + len);
}

MulData *parse_mul(char *input) {
  if (strlen(input) < 4) {
    return NULL;
  }

  char *ptr = input;

  char cmp[5];
  memcpy(cmp, ptr, 4);
  cmp[4] = '\0';

  if (strcmp(cmp, "mul(") != 0) {
    return NULL;
  }
  ptr += 4;

  NumData *num1_data = parse_num(ptr);
  if (num1_data == NULL) {
    return NULL;
  }

  ptr = num1_data->ending_ptr;
  if (*(ptr++) != ',') {
    free(num1_data);
    return NULL;
  }

  NumData *num2_data = parse_num(ptr);
  if (num2_data == NULL) {
    free(num1_data);
    return NULL;
  }

  ptr = num2_data->ending_ptr;
  if (*ptr != ')') {
    free(num1_data);
    free(num2_data);
    return NULL;
  }

  return new_mul_data(num1_data->num, num2_data->num, ptr);
}

int main() {
  char *input = read_file_content("inputs/day3");
  if (input == NULL) {
    return 1;
  }

  int sum = 0;
  char *search = strchr(input, 'm');

  while (search != NULL) {
    MulData *mul_data = parse_mul(search);
    if (mul_data != NULL) {
      sum += mul_data->num1 * mul_data->num2;
      search = strchr(mul_data->ending_ptr, 'm');
    } else {
      search = strchr(search + 1, 'm');
    }

    free(mul_data);
  }

  char buf[100];
  sprintf(buf, "%d", sum);
  write_to_file("outputs/day3/output-part1", buf);

  free(input);
  return 0;
}
