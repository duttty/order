package order

import (
	"strconv"
	"strings"
)

type tpStr struct {
	Num int
	Val string
	Idx int
}

//Strings according to the first consecutive number inside the string and
//sorts slices from small to large.
func Strings(s []string) {
	temp := classify(s)
	for i := range temp {
		s[i] = temp[i].Val
	}

}

func classify(s []string) []*tpStr {
	l := len(s)
	yNum := make([]*tpStr, l, l)
	n := l - 1
	y := 0
	for _, v := range s {
		num, idx := FdNum(v)
		temp := &tpStr{
			Num: num,
			Val: v,
			Idx: idx,
		}
		//不符合条件
		if num == -1 {
			yNum[n] = temp
			n--
		} else {
			i := y
			for ; i > 0; i-- {
				if idx < yNum[i-1].Idx {
					//右移
					yNum[i] = yNum[i-1]
				} else if idx > yNum[i-1].Idx {
					break
				} else {
					for ; i > 0; i-- {
						if num < yNum[i-1].Num {
							//右移
							yNum[i] = yNum[i-1]
						} else {
							break
						}
					}
					break
				}
			}
			//插入
			yNum[i] = temp
			y++
		}
	}
	return yNum
}

//FdNum return consecutive numbers and where they start
func FdNum(s string) (o, idx int) {
	str := ""
	fIdx := -1
	l := len(s) - 1
	for k, v := range s {
		//数字0-9
		if v >= '0' && v <= '9' {
			if fIdx == -1 {
				fIdx = k
			}
			if k == l {
				str = s[fIdx:]
			}
		} else {
			if v == ' ' {
				continue
			}
			if fIdx != -1 {
				str = s[fIdx:k]
				break
			}
			idx++
		}
	}
	o, err := strconv.Atoi(strings.Replace(str, " ", "", -1))
	if err != nil {
		return -1, idx
	}
	return
}
