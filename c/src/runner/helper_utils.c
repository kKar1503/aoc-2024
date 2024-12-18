#include "../../include/runner/helper_utils.h"
#include <assert.h>
#include <stdlib.h>
#include <string.h>

char **str_split(char *a_str, const char a_delim, size_t *count) {
  a_str = strdup(a_str);
  char **result = 0;
  *count = 0;
  char *tmp = a_str;
  char *last_comma = 0;
  char delim[2];
  delim[0] = a_delim;
  delim[1] = 0;

  /* Count how many elements will be extracted. */
  while (*tmp) {
    if (a_delim == *tmp) {
      (*count)++;
      last_comma = tmp;
    }
    tmp++;
  }

  /* Add space for trailing token. */
  *count += last_comma < (a_str + strlen(a_str) - 1);

  result = malloc(sizeof(char *) * *count);

  if (result) {
    size_t idx = 0;
    char *token = strtok(a_str, delim);

    while (token) {
      assert(idx < *count);
      *(result + idx++) = strdup(token);
      token = strtok(0, delim);
    }
    assert(idx == *count);
  }

  return result;
}

char **remove_element_str(char **array, size_t size, size_t i) {
  char **new_array = malloc((size - 1) * sizeof(char *));
  char **p = array;
  char **p_new = new_array;
  for (size_t j = 0; j < size; j++) {
    if (j != i) {
      *p_new = strdup(*p);
      p_new++;
    }
    p++;
  }
  return new_array;
}
