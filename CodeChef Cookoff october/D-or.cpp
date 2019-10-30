#include<bits/stdc++.h>
using namespace std;
typedef long long int ll;
int main(){
    ll t;
    cin>>t;
    while(t--)
    {
        ll a,b;
        cin>>a>>b;
        ll p = 1;
        while(p<b)
        {
            p*=2;
        }
        p/=2;
        if(p-1>a)
        {
            cout<<(b|(p-1))<<endl;
        }
        else
        {
            ll p[64]={0},q[64]={0};
            ll c[64] = {0};
            swap(b,a);
            ll r = a,w = b;
            ll i = 0;
            while(i<=63)
            {
                if(a%2==0)
                {
                    p[i] = 0;
                }
                else
                {
                    p[i] = 1;
                }
                a/=2;
                i++;
            }
            i = 0;
            while(i<=63)
            {
                if(b%2==0)
                {
                    q[i] = 0;
                }
                else
                {
                    q[i] = 1;
                }
                b/=2;
                i++;
            }
            ll flag = 0;
            for(int i = 63 ; i >= 0; i--)
            {
                if(p[i]==q[i])
                {
                    c[i] = p[i];
                    continue;
                }
                else
                {
                    for(int j = i ; j >=0 ; j--)
                    {
                        if(p[j]==1)
                        {
                            for(int k = j ; k>=0 ; k--)
                            {
                                c[k] = 1;
                            }
                            break;
                        }
                        else
                        {
                            c[j]  = 0;
                        }
                    }
                    break;
                }
            }
            ll ans = 0;
            for(int i = 63 ; i >= 0 ; i--)
            {
                ans = ans*2+c[i];
            }
            cout<<ans<<endl;
        }
    }
}
