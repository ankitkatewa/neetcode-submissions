func topKFrequent(nums []int, k int) []int {

	mp:=make(map[int]int,len(nums))
	for _ ,value:=range nums{
		mp[value]++
	}

	result :=make([][]int,len(nums)+1)
	for key,value:=range mp{
		result[value]=append(result[value],key)
	}

	ans:=make([]int,0,k)
	for i:=len(result)-1;i>=0;i--{
		if len(result[i])>0{
			ans=append(ans,result[i]...)
		}
		if len(ans)>=k{
			return ans[:k]
		}
	}
		return ans
	

}
