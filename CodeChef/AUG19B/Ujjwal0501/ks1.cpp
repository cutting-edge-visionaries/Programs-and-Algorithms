#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

void func(long long int arr[], long long int n) {
  long long int sum = 0;
  unordered_map <long long int, vector<long long int>> mp;

  // for (int i = 0; i < n; i++) cout << arr[i] << " ";
  // cout << "\n";
  // mp[0].push_back(0);
  for (long long int i = 0; i < n; i++) {
    long long int s, temp;

    if (arr[i] == 0) sum += i;
    if (mp[arr[i]].size() > 0) {
      // cout << arr[i] << " - [" << mp[arr[i]][0] << mp[arr[i]][1] << mp[arr[i]][2] << "] ";
      temp = (mp[arr[i]][1])*(i) - mp[arr[i]][0];
      sum += temp;
      // cout << temp << " ";
      mp[arr[i]][0] += i+1;
      mp[arr[i]][1]++;
      // cout << "[" << mp[arr[i]][0] << mp[arr[i]][1] << mp[arr[i]][2] << "]\n";
    } else {
      mp[arr[i]].push_back(i+1);
      mp[arr[i]].push_back(1);
    }
    // mp[arr[i]].push_back(0);//cout << mp[arr[i]][0] << " ";
  }//cout << "\n";

  cout << sum << "\n";
  return ;
}

int main(void) {

  long long int t;
  cin >> t;

  while (t--) {
    long long int n;
    cin >> n;

    long long int arr[n+1];

    arr[0] = 0;
    for (long long int i = 1; i <= n; i++) {
      long long int temp;
      cin >> temp;
      arr[i] = temp ^ arr[i-1];
      // cout << arr[i] << " ";
    }
// cout << "\n";
    func(arr+1, n);
  }

  return 0;
}
