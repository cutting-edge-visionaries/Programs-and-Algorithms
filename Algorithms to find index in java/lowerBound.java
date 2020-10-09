public int lowerBound(ArrayList<Integer> ar,int k)
{
    int s=0;
    int e=ar.size();
    while (s !=e)
    {
        int mid = s+e>>1;
        if (ar.get(mid) <k)
        {
            s=mid+1;
        }
        else
        {
            e=mid;
        }
    }
    if(s==ar.size())
    {
        return -1;
    }
    return s;
}