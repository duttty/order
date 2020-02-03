package order

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	s := []string{"usel ess2", "第1讲它的1的位置在1", "useless1",
		"uu1它的1的位置是2", "第1 2讲中途有空格会忽略也视为连续", "第10讲后面多余的12忽略掉", "无关元素", "关元素", "元素",
		"第111讲", " 10开始", "01的开始位置是0", "011算做11"}
	Strings(s)
	fmt.Println(s)
}
