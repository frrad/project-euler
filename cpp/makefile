CC = g++ -std=c++14
FLAGS = -Wall -Werror
CFLAGS = -c

PROBLEMS = 347

# problems

347: Problem347.o number_theory.o
	$(CC) $(FLAGS) Problem347.o number_theory.o -o 347

Problem347.o: number_theory.o Problem347.cc
	$(CC) $(FLAGS) $(CFLAGS) Problem347.cc

# libraries

number_theory.o: number_theory.h number_theory.cc
	$(CC) $(FLAGS) $(CFLAGS) number_theory.cc

# other

clean:
	\rm *.o $(PROBLEMS)
