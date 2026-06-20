func minEatingSpeed(piles []int, h int) int {

    l, r := 1, 0
    for _, pile := range piles{
        if pile > r{
            r = pile
        }
    }
    res := r
    for l <= r {
        totalHours := 0
        k := (l+r)/ 2
        for _, pile := range piles{
            totalHours += int(math.Ceil(float64(pile)/ float64(k)))
        }

        if totalHours <= h{
            res = k
            r = k-1
        }else{
            l = k+1
        }
    }
    return res

}
