package offer
func myPow(x float64, n int) float64 {
	var res float64
	res = 1.0
	for {
		if n < 0 {
			
			if abs(n)&1==1 {
				res/=x	
			}
			n/=2
			x*=x
		}
		if n > 0 {
			if abs(n)&1==1 {
				res*=x
			}
			n>>=1
			x*=x
		} 
        if n==0{
            break
        }
	}
	return res
}

func abs(n int)int{
	if n>0 {
		return n
	}else{
		return 0-n
	}

}