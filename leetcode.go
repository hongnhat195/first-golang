package main

import (
	"fmt"
	"reflect"
	"sort"
)
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func threeSum(nums []int) [][]int {
	
   sort.Ints(nums)
   a := make([][]int, 0 )
	for i:= 0; i< len(nums) -2 ; i++ { 
		flag := true
		if nums[i] > 0 {
			break;
		}
		for j:=  i+1 ; j< len(nums) -1 ; j++ {  
			if nums[i] + nums[j]  <= 0 {
				for k:= j+1 ; k< len(nums)  ; k++ {
					if 0 - (nums[i] + nums[j]) == nums [k]  {
						x := []int { nums[i], nums[j] , nums[k]}
                        fmt.Println(x)
						for n := 0; n < len(a); n++ {
							if Equal(a[n], x)  {
								flag = false
                                break
							}
						}
						if flag {
							a = append(a, x)
                            flag = true
						}
						
					}
				}
			} 
				
		}
		
	}
	return a
}

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
	var (
		flag = true
		a = make([][]int, 0)
		i =0
		j= 1
		k = len(nums)-1
		h = len(nums) -2
	)
    if len(nums) < 4 {
		val := [][]int {}
        return val
    }
  
	for i < k {
		fmt.Println(i,j, h, k)
		sum := nums[i] + nums[j] + nums[h] + nums[k]
		if ( sum == target) {
			t := []int {nums[i], nums[j] , nums[h] , nums[k] }
            sort.Ints(t)
			for n :=0; n < len(a); n++ {
				if  reflect.DeepEqual(a[n], t) {
					flag = false
					break
				}
			}
			if flag {
				a = append(a, t )
			}
			flag = true
			if j +1 ==h {
				i +=1
				j = i+1
			} else {
				j++
			}
		} else if sum < target { 
			if j +1 ==h {
				i +=1
				j = i+1
			} else {
				j++
			}
		}else if sum > target { 
			if h-1 == j {
				k -=1
				h = k-1
			} else {
				h--
			}
		}
	};
	return  a
}