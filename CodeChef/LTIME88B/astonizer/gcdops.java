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
        while (t-- > 0) {
            int n = sc.nextInt();  
            int[] b = new int[n];
            boolean check = false;
            for (int i = 0; i < n; i++) {
                b[i] = sc.nextInt();
            }
            for (int i = 0; i < n; i++) {
                if ((i+1) % b[i] != 0) {
                    check = true;
                    break;
                }
            }
            if(!check) {
                pw.println("YES");
            } else {
                pw.println("NO");
            }
        }
        pw.flush();
    }
}
