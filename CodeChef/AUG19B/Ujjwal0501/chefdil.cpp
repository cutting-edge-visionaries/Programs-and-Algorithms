#include <iostream>

using namespace std;
int main(void) {

  long long int t;
  cin >> t;

  while (t--) {
    int count = 0;
    string a;
    cin >> a;

    for (int i = 0; a[i] != '\0'; i++)
      if (a[i] == '1') count++;

    if (count > 0 && count%2 == 1) cout << "WIN\n";
    else cout << "LOSE\n";
  }

  return 0;
}
