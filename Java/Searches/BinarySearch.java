import java.io.*;
import java.util.*;
/*
* Code to search linearly in a give array
* Worst time Complexity - O(N) and Best time - O(1)
*
*/


public class BinarySearch {
    public int search(int arr[] , int key){
        int start = 0;
        int end = arr.length -1;
        while(start<=end){
            int mid = (start+end)/2;
            if(arr[mid]==key){
                return (mid+1);
            }
            else if(arr[mid]<key){
                end = mid-1;
            }
            else{
                start = mid+1;
            }
        }
        return -1;
    }
    public void searchResult(int arr[], int key){
        int pos = search(arr, key);
        if(pos!=-1){
            System.out.println("Found at pos "+ pos);
        }
        else{
            System.out.println("Not Found ;[ ");
        }
    }
    public static void main(String args[]) throws Exception{
        Scanner sc = new Scanner(System.in);
        BinarySearch bs = new BinarySearch();
        // Input for length of Array
        int N = sc.nextInt();
        int keys[] = new int[N+1];
        for(int i=0; i<N; ++i) keys[i] = sc.nextInt();
        int key = sc.nextInt(); // Key to be searched
        bs.searchResult(keys, key);
    }
}