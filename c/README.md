## Advent of Code 2024 Attempts in C

This directory will contain all the attempts that I had in C.

## Motivation

This year I decided to try writing some of the attempts in C since
I have enrolled in SMU this year, and I note that CS101 was taught in C,
and now that I have gained newfound knowledge, so why not put it to good
use.

## How to Run Code

To run the code, you will need to have a C compiler and Make installed
on your machine. I am using Clang, but you should be able to use other
compiler that you have. But my Makefile is written with Clang, so you
will have to adjust that.

With Clang and Make installed, you can run the following in sequence:

1. First, add your input file into the `inputs/` directory. The file
   should be named `dayX` where `X` is the day of the challenge.

2. Run following to generate an output file into the `bin/` directory:

   ```bash
   make build/dayX/partY # where X is the day and Y is the part
   ```

3. Then, run the following to run the output binary to generate the
   program's output into the `outputs/` directory:

   ```bash
   make run/dayX/partY # where X is the day and Y is the part
   ```

4. Finally, run the following to read the program's output:

   ```bash
   make read/dayX/partY # where X is the day and Y is the part
   ```

5. If you want to clean up the output files, you can run the following:

   ```bash
   make clean
   ```

You can also just run the following:

```bash
make read/dayX/partY # where X is the day and Y is the part
```

Since within the Makefile the target `read` is dependent on the `build`
and `run` targets, it will automatically build and run the program
for you.
