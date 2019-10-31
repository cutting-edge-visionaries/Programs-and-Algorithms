#include<bits/stdc++.h>
using namespace std;
typedef long long int ll;
set<ll>s;
void precompute()
{
	for(auto it = s.begin(); it != s.end() ; it++)
	{
		if((*it)*10>(ll)1e18)
		{
			continue;
		}
		else
		{
			s.insert((*it)*10);
		}
		
		if((*it)*20>(ll)1e18)
		{
			continue;
		}
		else
		{
			s.insert((*it)*20);
		}
	}
}
int main(){
	s.insert(1);
	precompute();
	ll t;
	cin>>t;
	while(t--)
	{
	    ll n;
	    cin>>n;
	    if(s.find(n)==s.end()){
	        cout<<"No"<<endl;
	    }
	    else{
	        cout<<"Yes"<<endl;
	    }
	}
}
