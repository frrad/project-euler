#include <iostream>
#include "number_theory.h"

int main() {
  for (int x = 1; x < 100; x++) {
    if (is_prime(x)) {

      std::cout << x << ", ";
    }
  }
}
