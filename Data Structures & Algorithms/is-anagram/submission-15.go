func isAnagram(s string, t string) bool {

   if len(s) != len(t){
    return  false
   }

   sMap := make(map[byte]int)

   for i := 0; i < len(s); i++{
      sMap[s[i]]++
   } 
  
   tMap := make(map[byte]int)
   for j := 0; j < len(t); j++{
    tMap[t[j]]++
   }


   for key, val := range sMap{
       tKey, ok := tMap[key]
       if  !ok{
        return false
       }

       if tKey != val{
        return false
       } 
   }


    return true

}


