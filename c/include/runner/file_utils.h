#ifndef FILE_UTILS_H
#define FILE_UTILS_H

#include <stdio.h>
#include <stdlib.h>

// Reads the content of a file into a dynamically allocated string.
// Returns NULL if the file can't be opened or read.
char *read_file_content(const char *file_path);

// Writes a string to a file. Overwrites the file if it exists.
void write_to_file(const char *file_path, const char *content);

#endif // FILE_UTILS_H
