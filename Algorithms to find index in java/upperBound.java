public int upperBound(int[] arr, int key){
    int s = 0;
    int e = arr.length;
    while(s != e){
        int mid = (s+e)>>1;
        if(arr[mid] >= k)                  // same as lower bound only diff. is in this line
            s = mid+1;
        else
            e = mid;
    }
    if(s == arr.length)
        return -1;

    return s;
}