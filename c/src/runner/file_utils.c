#include "../../include/runner/file_utils.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Reads the content of the file into a dynamically allocated string
char *read_file_content(const char *file_path) {
  FILE *file = fopen(file_path, "r");
  if (file == NULL) {
    perror("Failed to open file");
    return NULL;
  }

  // Find the file size
  fseek(file, 0, SEEK_END);
  long file_size = ftell(file);
  rewind(file);

  // Allocate memory for the content and read it
  char *content = (char *)malloc(file_size + 1); // +1 for null terminator
  if (content == NULL) {
    perror("Failed to allocate memory");
    fclose(file);
    return NULL;
  }

  fread(content, 1, file_size, file);
  content[file_size] = '\0'; // Null-terminate the string

  fclose(file);
  return content;
}

// Writes a string to a file, overwriting if it exists
void write_to_file(const char *file_path, const char *content) {
  FILE *file = fopen(file_path, "w");
  if (file == NULL) {
    perror("Failed to open file for writing");
    return;
  }

  fwrite(content, sizeof(char), strlen(content), file);
  fclose(file);
}
