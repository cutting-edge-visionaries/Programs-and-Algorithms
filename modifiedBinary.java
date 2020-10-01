import java.io.*;
import java.util.*;

class TestClass {
    public static void main(String args[] ) throws Exception {
        // Write your code here
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String[] str = br.readLine().split(" ");
        int n = Integer.parseInt(str[0]);
        int m = Integer.parseInt(str[1]);
        String[] an = br.readLine().split(" ");
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        int[] a = new int[n];
        for(int i=0;i<n;i++)
            a[i] = Integer.parseInt(an[i]);
        long l = 1, r = 10000000, ans = 10000000;
        while(l <= r){
            long mid = (l+r)/2;
            long s = 0;
            for(int i = 0; i < n; i++)
                s += mid/a[i];
            if(s >= m){
                ans = mid;
                r = mid - 1;
            }
            else
                l = mid + 1;
        }
        System.out.println(ans);
    }
}
