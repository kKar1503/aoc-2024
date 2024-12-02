#ifndef HELPER_UTILS_H
#define HELPER_UTILS_H

#include <stdio.h>
#include <stdlib.h>

// Split string by a delimiter
char **str_split(char *a_str, const char a_delim, size_t *count);

// Return a copy of array with ith element removed
char **remove_element_str(char **array, size_t size, size_t i);

#endif // HELPER_UTILS_H
