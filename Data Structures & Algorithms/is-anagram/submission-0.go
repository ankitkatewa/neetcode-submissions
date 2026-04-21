func isAnagram(s string, t string) bool {
	
	if len(s)!=len(t){
		return false
	}

	var freqcount [256]int

	for i:=0;i<len(s);i++{
		freqcount[s[i]]++
		
	}

	for i:=0;i<len(s);i++{
		freqcount[t[i]]--
		if(freqcount[t[i]]<0){
			return false
		}
	}

	return true

}
