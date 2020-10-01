#include<bits/stdc++.h>
using namespace std;
typedef long long int ll;
int main(){
	ll t;
	cin>>t;
	while(t--)
	{
		ll n;
		cin>>n;
		ll a[n];
		for(int i = 0 ; i< n ; i++)
		{
			cin>>a[i];
		}
		ll indx,indy;
		indx = 0;
		indy = 0;
		ll max = a[0],min = a[0];
		for(int i = 0 ;i < n ; i++)
		{
			if(a[i]>max)
			{
				max = a[i];
				indx = i;
			}
			if(a[i]<min)
			{
				min = a[i];
				indy = i;
			}
		}
		if(indx>indy)
		{
			cout<<a[indy]<<" "<<a[indx];
		}
		else
		{
			cout<<a[indx]<<" "<<a[indy];
		}
		cout<<endl;
	}
}
