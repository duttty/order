# order

"Order" sorts string slices from small to large based on the first consecutive number in the string.
This allows you to quickly sort and classify files by name. It uses the idea of inserting sorting algorithm to make its time complexity O(n), and because of the classification, the spatial complexity is O(n).

## How To Use
```go
go get github.com/duttty/order
```

```go

package main

import github.com/duttty/order

func main(){
	s := []string{"usel ess2", "第1讲它的1的位置在1", "useless1",
        "uu1它的1的位置是2", "第1 2讲中途有空格会忽略也视为连续", 
        "第10讲后面多余的12忽略掉", "无关元素", "关元素", "元素",
		"第111讲", " 10开始"}
	Strings(s)
	fmt.Println(s)
}


/* 
    [ 10开始 第1讲它的1的位置在1 第10讲后面多余的12忽略掉 第1 2讲中途有空格会忽略也视为连续 第111讲 uu1它的
1的位置是2 useless1 usel ess2 元素 关元素 无关元素]
*/
```