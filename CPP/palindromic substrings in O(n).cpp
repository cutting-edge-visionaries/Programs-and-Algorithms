#include <bits/stdc++.h>
using namespace std;

int main()
{
    cout << "no. of test cases:";
    int t;
    cin >> t;
    // scanf("%i",&t) >> t;
    for(int a=0;a<t;a++)
    {
        long sum=0;
        
        int n;
        cout <<"string length:";
        cin >> n;
        // scanf("%d",&n);
        // sum+=n;
        cout << "enter string: ";
        string s;
        cin >> s;
        vector<int> odd(n);
    for (int i = 0, l = 0, r = -1; i < n; i++) 
    {
        int k = (i > r) ? 1 : min(odd[l + r - i], r - i + 1);
        while (0 <= i - k && i + k < n && s[i - k] == s[i + k]) 
        {
            k++;
        }
        odd[i] = k--;
        if (i + k > r) 
        {
            l = i - k;
            r = i + k;
        }
    }
    vector<int> even(n);
    for (int i = 0, l = 0, r = -1; i < n; i++) 
    {
        int k = (i > r) ? 0 : min(even[l + r - i + 1], r - i + 1);
        while (0 <= i - k - 1 && i + k < n && s[i - k - 1] == s[i + k]) 
        {
        k++;
        }
        even[i] = k--;
        if (i + k > r) 
        {
            l = i - k - 1;
            r = i + k ;
        }

    }
        for(int i=0;i<n;i++)
        sum+=odd[i]+even[i];
        cout << "no. of palinedromic substrings: " << sum << "\n";
        // printf("%ld\n",sum);
    }
}