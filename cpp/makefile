CC = g++
FLAGS = -Wall -Werror
CFLAGS = -c

347: number_theory.o Problem347.cc
	$(CC) $(FLAGS) Problem347.cc

number_theory.o: number_theory.h number_theory.cc
	$(CC) $(FLAGS) $(CFLAGS) number_theory.cc

clean:
	\rm *.o

