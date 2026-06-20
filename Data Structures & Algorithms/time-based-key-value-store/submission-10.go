type TimeMap struct {
  store map[string][]Store
}

type Store struct{
    value  string
    time int
}

func Constructor() TimeMap {
    return TimeMap{
        store: make(map[string][]Store),
    }
}

func (this *TimeMap) Set(key string, timestampvalue string, timestamp int) {
    this.store[key] = append(this.store[key],Store{
                        value:timestampvalue,
                        time:timestamp,
                      })
}

func (this *TimeMap) Get(key string, timestamp int) string {
     ans := ""
     if sre, ok := this.store[key]; !ok{
        return ""
     }else{
        left, right := 0, len(sre)-1
        for left <= right{
            mid := right + left / 2
            if sre[mid].time <= timestamp{
                 ans = sre[mid].value
                 left = mid+1
            }else{
                right = mid-1
            }
        }
     }

     return ans
}
