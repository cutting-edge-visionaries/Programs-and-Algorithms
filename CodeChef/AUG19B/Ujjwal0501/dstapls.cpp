#include <iostream>

using namespace std;
int main(void) {

  long long int t;
  cin >> t;

  while (t--) {
    long long n, k;
    cin >> n >> k;

    if (k == 1) cout << "NO\n";
    else {
      long long int a = n/k;
      if (a%k == 0) cout << "NO\n";
      else cout << "YES\n";
    }
  }

  return 0;
}
