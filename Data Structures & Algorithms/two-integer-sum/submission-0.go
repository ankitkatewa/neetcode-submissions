func twoSum(nums []int, target int) []int {
	mp:=make(map[int]int)
	

	for i:=0;i<len(nums);i++{
		index, ok := mp[target-nums[i]];if ok{
			return []int{index,i}
		}
		mp[nums[i]]=i
	}

	return []int{}


    
}
