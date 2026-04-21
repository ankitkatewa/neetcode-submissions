func groupAnagrams(strs []string) [][]string {
	

	 mp:= make(map[string][]string,len(strs))

	 for _, str:=range strs {
		sortedstr:= sortString(str)
		value, ok:=mp[sortedstr]; if ok{
			mp[sortedstr]=append(value,str)
		}else{
			mp[sortedstr]=[]string{str}
		}
	 }
	 result := make([][]string, 0, len(mp))

for _, value := range mp {
    result = append(result, value)
}
return result

}
func sortString(s string) string{
	runes:=[]rune(s)

	sort.Slice(runes,func(i,j int)bool{
		return runes[i]<runes[j]
	})

	return string(runes)
}
