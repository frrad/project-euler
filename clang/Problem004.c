#include "stdio.h"
#include <stdbool.h>

#define FROM 100
#define TO 1000

int reverse(int n) {
    int ans = 0;
    int offset = 1;
    while (n > 0) {
        ans = 10 * ans + n % 10;
        offset *= 10;
        n /= 10;
    }
    return ans;
}

bool palindrome(int n) { return n == reverse(n); }

int main() {
    int largest = 0;

    for (int i = FROM; i < TO; i++) {
        for (int j = i; j < TO; j++) {
            if (i * j > largest && palindrome(i * j))
                largest = i * j;
        }
    }

    printf("%d\n", largest);

    return 0;
}
