#include "stdio.h"

int fib(int n) {
    if (n == 1 || n == 2)
        return 1;
    return fib(n - 1) + fib(n - 2);
}

int main() {

    int i = 1;
    int sum = 0;

    while (fib(i) < 4000000) {
        if (fib(i) % 2 == 0)
            sum += fib(i);
        i++;
    }
    printf("%d\n", sum);

    return 0;
}
