func removeElement(nums []int, val int) int {

	k:=0
    for _,number :=range nums{
		if number!=val{
			nums[k]=number
			k++
		}
	}

	return k
}
