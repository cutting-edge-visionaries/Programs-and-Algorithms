import java.io.*;
import java.util.*;

class ZAlgorithm {
    public static void main(String args[] ) throws Exception {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String c = br.readLine();
        String s = br.readLine();
        int m = c.length();
        int n = s.length();
        c += "#" + s;
        int count = 0;
        int[] z = new int[n+m+1];

        //System.out.println(c);
        int L = 0, R = 0;
        for (int i = 1; i < m+n+1; i++){
            if (i > R){
                L = R = i;
                while (R < n && c.charAt(R-L) == c.charAt(R)){
                    R++;
                }
                z[i] = R-L; 
                R--;
            } 
            else{
                int k = i-L;
                if (z[k] < R-i+1){
                    z[i] = z[k];
                } 
                else{
                    L = i;
                    while (R < n && c.charAt(R-L) == c.charAt(R)){
                        R++;
                    }
                    z[i] = R-L; 
                    R--;
                }
            }

            if(z[i] == m)
                count++;
        }

        // for(int x: z)
        //     System.out.print(x+"");

        System.out.println(count);
    }
}
