func search(nums []int, target int) int {
// nums = [3,4,5,6,1,2], target = 1

  l, r := 0, len(nums)-1

  for l <= r{
    mid := (l+r)/ 2

    if nums[mid] == target{
        return mid
    }
    // left array
    if nums[l] <= nums[mid]{
        if target > nums[mid] || target < nums[l]{
            l = mid+1
        }else{
            r = mid-1
        }
    }else{
        // right array
        if target < nums[mid] || target > nums[r]{
            r = mid-1
        }else{
            l = mid+1
        }
    }  
  }

  return -1

}
