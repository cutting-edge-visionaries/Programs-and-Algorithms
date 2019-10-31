#include <iostream>
#include <algorithm>

using namespace std;
int main(void) {

  long long int t;
  cin >> t;

  while (t--) {
    long long int n, count = 0;
    int flag = 0;
    cin >> n;

    long long int arra[n+1] = {0}, arrh[n+1] = {0};

    for (long long int i = 1; i <= n; i++) {
      long long int temp, j, last;
      cin >> temp;
      j = i-temp>1?i-temp:1;
      if (i-temp > 1) {
        arra[i-temp]++;
      } else {
        arra[1]++;
      }
      if (i+temp < n) {
        arra[i+temp+1]--;
      }

      // for (j = 0; j <= n; j++) cout << arra[j] << " ";
      // cout << "\n";
    }

    for (long long int i = 1; i <= n; i++) arra[i] += arra[i-1];
    sort(arra, arra+n+1);

    // for (int j = 0; j <= n; j++) cout << arra[j] << " ";
    // cout << "\n";

    for (long long int i = 1; i <= n; i++) {
      cin >> arrh[i];
    }
    sort(arrh, arrh+n+1);

    for (long long int i = 1; i <= n; i++) {
      if (arrh[i] != arra[i]) {
        flag = 1;
        break;
      }
    }

    if (flag == 1) cout << "NO\n";
    else cout << "YES\n";

  }

  return 0;
}
