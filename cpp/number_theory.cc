#include <vector>

using std::vector;

bool is_prime(int n) {
  if (n <= 1) {
    return false;
  }
  for (int a = 2; a * a <= n; a++) {
    if (n % a == 0) {
      return false;
    }
  }
  return true;
}

vector<int> primes_below(int n) {
  vector<int> answer;
  int x = 2;

  while (x < n) {
    if (is_prime(x)) {
      answer.push_back(x);
    }
    x++;
  }

  return answer;
}
