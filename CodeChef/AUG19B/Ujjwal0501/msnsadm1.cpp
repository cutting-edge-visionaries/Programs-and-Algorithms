#include <iostream>
#include <algorithm>

using namespace std;
int main(void) {

  long long int t;
  cin >> t;

  while (t--) {
    long long int n, max = 0;
    cin >> n;

    long long int arr[n], arra[n];
    for (long long int i = 0; i < n; i++) cin >> arr[i];
    long long int temp = 0;
    for (long long int i = 0; i < n; i++) {
      cin >> temp;
      arra[i] = 20*arr[i] - 10*temp;

      if (max < arra[i]) max = arra[i];
    }

    // sort(arra, arra+n);
    if (max < 0) cout << 0 << "\n";
    else cout << max << "\n";
  }

  return 0;
}
