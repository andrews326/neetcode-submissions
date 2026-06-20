func groupAnagrams(strs []string) [][]string {
      
      anagramMap := make(map[string][]string)

      for _, v := range strs{
         hashKey := sortString(v)

         if _, ok := anagramMap[hashKey]; ok{
            anagramMap[hashKey] = append(anagramMap[hashKey], v)
         }else{
            anagramMap[hashKey] = []string{v}
         }

         
      }


      results := [][]string{}

      for _ , value := range anagramMap{
        results =  append(results,value)
      }

      return  results


}


func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
