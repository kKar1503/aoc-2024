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

typedef struct DoDontData {
  bool doo;
  bool dont;
  char *ending_ptr;
} DoDontData;

DoDontData *new_do_dont_data(bool doo, bool dont, char *ending_ptr) {
  DoDontData *data = malloc(sizeof(DoDontData));
  data->doo = doo;
  data->dont = dont;
  data->ending_ptr = ending_ptr;
  return data;
}

DoDontData *parse_do_dont(char *input) {
  char *ptr = input;
  if (strlen(ptr) < 2 || ptr[0] != 'd' || ptr[1] != 'o') {
    return NULL;
  }

  ptr += 2;

  // check for "do()"
  if (strlen(ptr) >= 2 && ptr[0] == '(' && ptr[1] == ')') {
    return new_do_dont_data(true, false, ptr + 2);
  }

  // check for "don't()"
  if (strlen(ptr) >= 5 && ptr[0] == 'n' && ptr[1] == '\'' && ptr[2] == 't' &&
      ptr[3] == '(' && ptr[4] == ')') {
    return new_do_dont_data(false, true, ptr + 5);
  }

  return NULL;
}

char *get_earlier_ptr(char *ptr1, char *ptr2) {
  if (ptr1 == NULL) {
    return ptr2;
  }
  if (ptr2 == NULL) {
    return ptr1;
  }
  return ptr1 < ptr2 ? ptr1 : ptr2;
}

int main() {
  char *input = read_file_content("inputs/day3");
  if (input == NULL) {
    return 1;
  }

  int sum = 0;
  char *search_m = strchr(input, 'm');
  char *search_d = strchr(input, 'd');
  bool doo = true;
  bool dont = false;

  while (search_m != NULL || search_d != NULL) {
    char *search = get_earlier_ptr(search_m, search_d);

    if (search == search_d) {
      DoDontData *do_dont_data = parse_do_dont(search);
      if (do_dont_data != NULL) {
        doo = do_dont_data->doo;
        dont = do_dont_data->dont;
        search_m = strchr(do_dont_data->ending_ptr, 'm');
        search_d = strchr(do_dont_data->ending_ptr, 'd');
      } else {
        search_m = strchr(search + 1, 'm');
        search_d = strchr(search + 1, 'd');
      }
      free(do_dont_data);
      continue;
    }

    MulData *mul_data = parse_mul(search);
    if (mul_data != NULL) {
      if (doo) {
        sum += mul_data->num1 * mul_data->num2;
      }
      search_m = strchr(mul_data->ending_ptr, 'm');
      search_d = strchr(mul_data->ending_ptr, 'd');
    } else {
      search_m = strchr(search + 1, 'm');
      search_d = strchr(search + 1, 'd');
    }
    free(mul_data);
  }

  char buf[100];
  sprintf(buf, "%d", sum);
  write_to_file("outputs/day3/output-part2", buf);

  free(input);
  return 0;
}
