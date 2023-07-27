package logic

func Exer_cal(shour int,smin int,ehour int,emin int,stype int) int {

	//游泳每分钟8
	//跑步每分钟10
	//跳绳15

	//分钟计算
	var minsum int


	if shour==ehour {
		minsum = smin-emin
	}else {
		hourcout := shour - ehour
		minsum = hourcout*60
		minsum = minsum + (60-smin)
		minsum = minsum + emin
	}
	switch stype {
	case 1:{
		return minsum*10
	}
	case 2:{
		return minsum*15
	}
	case 3:{
		return minsum*8
	}
	}
	return 0
}

func Diet_cal(dtype int,foodt int,weight int)  int{
	var t float32
	t = float32(weight / 100)
	var a int
	a = int(t)
	return a
}