#include "stdio.h"

int main() {
    long long int n = 600851475143;
    int i = 2;

    while (i * i <= n) {
        if (n % i == 0) {
            n /= i;
        } else {
            i++;
        }
    }

    printf("%d\n", n);

    return 0;
}
