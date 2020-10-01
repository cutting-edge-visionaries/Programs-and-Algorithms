/* package codechef; // don't place package name! */

import java.util.*;
import java.lang.*;
import java.io.*;

/* Name of the class has to be "Main" only if the class is public. */
class Codechef
{
    public static void main (String[] args) throws java.lang.Exception
    {
        Scanner sc = new Scanner(System.in);
        PrintWriter pw = new PrintWriter(System.out);
        int t = sc.nextInt();
        while(t-- > 0) {
            int n = sc.nextInt();
            int[] arr = new int[n];
            int flag = 0;
            for(int i = 0; i < n; i++) {
                arr[i] = sc.nextInt();
            }

            int sum = 0;
            for(int i = 0; i < n; i++) {
                sum += arr[i];
            }

            if(sum < 0) {
                System.out.println("NO");
            } else {
                System.out.println("YES");
            }
        }
        pw.flush();
    }
}
