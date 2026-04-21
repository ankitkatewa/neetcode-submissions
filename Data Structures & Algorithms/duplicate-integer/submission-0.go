func hasDuplicate(nums []int) bool {
	mp:=make(map[int]int)
   for _,num:= range nums{
	mp[num]++
	 value,_:=mp[num]
	 if value>1{
		return true
	 }
	

   }

   return false
                      
}