type MinStack struct {
    stack [][2]int
}

func Constructor() MinStack {
   return MinStack{
    stack : [][2]int{},
   }
}

func (this *MinStack) Push(val int) {
   if len(this.stack) == 0{
     this.stack = append(this.stack, [2]int{val,val})
   }else{
    currMin :=  this.stack[len(this.stack)-1][1]
    if val < currMin {
        currMin = val
    }
   
    this.stack = append(this.stack, [2]int{val,currMin})

   }
}

func (this *MinStack) Pop() {
    if len(this.stack) == 0{
        return
    }else{
        this.stack = this.stack[:len(this.stack)-1]
    }
}

func (this *MinStack) Top() int {
    if len(this.stack) == 0{
        return 0
    }else{
        top := this.stack[len(this.stack)-1][0]
        //this.stack = this.stack[:len(this.stack)-1]

        return top
    }
}

func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1][1]
}
