package dac

import (
	"fmt"
	"strconv"
)

/*
	分治算法 汉诺塔事件
*/

//Dac num 盘的个数  a, b,c 代表3个塔的顺序 表示从a塔移动到c塔
func Dac(num int, a, b, c string) {
	if num == 1 {
		fmt.Println("第1个盘从" + a + "移动到" + c)
		return
	}

	//当n大于等于2时，可以看作2个盘，最下面是一个盘，上面所有的盘是一个盘
	//将上面所有的盘从a移动到b盘，移动过程中使用c
	Dac(num-1, a, c, b)
	fmt.Println("第" + strconv.Itoa(num) + "个盘从" + a + "移动到" + c)
	//将b中的盘移动到c塔，移动过程使用到a
	Dac(num-1, b, a, c)
}
