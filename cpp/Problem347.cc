#include <iostream>
#include <vector>
#include <math.h>
#include "number_theory.h"

using std::vector;
using std::cout;

long M(long p, long q, long N) {
  // Assume p,q prime because who wants to check that?
  if (p < 1 || q < 1 || p * q > N || p == q) {
    return 0;
  }

  long best = p * q;

  long p_max = p;
  long q_max = q;

  while (p_max * q <= N) {

    long target = N / p_max;
    q_max = 1;

    while (target >= q) {
      target /= q;
      q_max *= q;
    }

    if (q_max * p_max <= N) {
      if (q_max * p_max > best) {
        best = q_max * p_max;
      }
    }
    p_max *= p;
  }

  return best;
}

long S(long n) {
  vector<int> primes = primes_below(static_cast<int>(n));
  long answer = 0;

  for (uint a = 0; a < primes.size(); a++) {
    for (uint b = a + 1; b < primes.size(); b++) {
      if (primes[a] * primes[b] > n) {
        break;
      }
      answer += M(primes[a], primes[b], n);
    }
  }

  return answer;
}

int main() {
  long size = pow(10, 7);
  cout << "S(" << size << ") = " << S(size) << "\n";
}
