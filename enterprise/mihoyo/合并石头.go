package main

/*
	n堆石头，第i堆石头的颜色是ai，按照编号依次摆放。
	合并相邻的石头，新的一堆石头放到原来两堆石头的位置，其他石头位置不变
	假设合并两堆石头，一堆石头的颜色数量为u，另一堆石头的颜色数量为v，那么这次合并的代价为uxv

	将n堆石头合并为1对石头，最小代价的总和是多少？

	5
	3 3 2 2 4
	5
*/

// import java.util.HashSet;
// import java.util.LinkedList;
// import java.util.Scanner;
// import java.util.Set;
//
/*
	java ac 20%
*/

// public class Main {
//
//    public static void main(String[] args) {
//        Scanner in = new Scanner(System.in);
//        int n = in.nextInt();
//        int[] color = new int[n];
//        for (int i = 0; i < n; i++) {
//            color[i] = in.nextInt();
//        }
//        LinkedList<Set<Integer>> list = new LinkedList<>();
//        int i =0;
//        int ans=0;
//        while(i<n){
//            Set<Integer> set = new HashSet<>();
//            set.add(color[i]);
//            int j = i+1;
//            boolean flag = false;
//            for (;j<n;j++){
//                if(!set.contains(color[j])){
//                    break;
//                }else {
//                    flag = true;
//                }
//            }
//            if (flag){
//                ans++;
//            }
//            i=j;
//            list.add(set);
//        }
//        for (int k = 1; k < list.size(); k++) {
//            ans+=list.get(k-1).size() * list.get(k).size();
//            Set<Integer> cur = list.get(k);
//            cur.addAll(list.get(k-1));
//            list.set(k,cur);
//        }
//        System.out.println(ans);
//    }
// }
