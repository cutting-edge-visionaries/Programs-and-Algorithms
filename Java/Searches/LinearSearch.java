import java.io.*;
import java.util.*;
/*
* Code to search linearly in a give array
* Worst time Complexity - O(N) and Best time - O(1)
*
*/


class LinearSearch {
    public int search(int arr[] , int key){
        for(int i=0; i<arr.length; i++){
            if(arr[i] == key){
                return (i+1);
            }
        }
        return -1;
    }
    public void searchResult(int arr[], int key){
        int pos = search(arr, key);
        if(pos!=-1){
            System.out.println("Found at pos"+pos);
        }
        else{
            System.out.println("Not Found ;[ ");
        }
    }
    public static void main(String args[]) throws Exception{
        Scanner sc = new Scanner(System.in);
        LinearSearch ls = new LinearSearch();
        // Input for length of Array
        int N = sc.nextInt();
        int keys[] = new int[N+1];
        for(int i=0; i<N; ++i) keys[i] = sc.nextInt();
        int key = sc.nextInt(); // Key to be searched
        ls.searchResult(keys, key);
    }
}