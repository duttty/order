# order

"Order" sorts string slices from small to large based on the first consecutive number in the string.

“Order”根据字符串中的第一个连续数字从小到大对字符串切片进行排序。

## Descripe
"Order" sorts the string slices from small to large based on the first consecutive number in the string. It ignores the space between two numbers and sorts according to the rule: "where the numbers appear < number size < no number", which is very helpful for quick sorting and classification of files. it draws on the idea of insertion sort algorithm, making its time complexity O (n). Because of the classification, the spatial complexity is O (n).

“Order”根据字符串中的第一个连续数字从小到大对字符串切片进行排序。它会忽略掉两个连续数字之间的空格，并按照: “数字出现的位置<数字大小<没有数字” 这种规则排序,这对文件进行快速排序和分类是非常有帮助的。它借鉴了插入排序算法的思想，使其时间复杂度O（n）。由于进行了分类，所以空间复杂度为O（n）。

## How To Use
```go
go get github.com/duttty/order
```

```go

package main

import github.com/duttty/order

func main(){
	s := []string{"usel ess2", "第1讲它的1的位置在1", "useless1",
		"uu1它的1的位置是2", "第1 2讲中途有空格会忽略也视为连续", "第10讲后面多余的12忽略掉", "无关元素", "关元素", "元素",
		"第111讲", " 10开始", "01的开始位置是0", "011算做11"}
	Strings(s)
	fmt.Println(s)
}


/* 
    [01的开始位置是0  10开始 011算做11 第1讲它的1的位置在1 第10讲后面多余的12忽略掉 第1 2讲中途有空格会忽略也视为连续 第111 讲 uu1它的1的位置是2 useless1 usel ess2 元素 关元素 无关元素]
*/
```