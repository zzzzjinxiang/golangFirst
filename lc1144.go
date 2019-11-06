package main

func zigzag(nums []int) int {
	n := len(nums);
	odd := [2]int{0, 1};
	ans := [2]int{0, 0};
	for i := range odd {
		for j := i; j < n; j = j + 2 {
			d := 0;
			if j > 0 {
				d = max(d, nums[i]-nums[i-1]+1);
			}
			if j+1 < n {
				d = max(d, nums[i]-nums[i+1]+1);
			}
			ans[i] += d;
		}
	}
	if (ans[0] == max(ans[0], ans[1])) {
		return ans[1];
	}
	return ans[0];
}

//func Max(a int, b int) int {
//	if (a > b) {
//		return a;
//	} else {
//		return b;
//	}
//}
//
//func Min(a int, b int) int {
//	if (a < b) {
//		return a;
//	} else {
//		return b;
//	}
//}
