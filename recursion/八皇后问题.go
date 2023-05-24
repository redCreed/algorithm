package recursion

import "fmt"

/*
	在一个8*8大小的国际象棋棋盘上放置8个皇后棋子，使得所有的皇后都是安全的（即任意两个皇后都无法攻击到对方）。
	按照国际象棋的规则，皇后的攻击方式是横，竖和斜向。
	皇后可以攻击到同一列所有其它棋子，因此可推导出每1列只能存在1个皇后，即每个皇后分别占据一列。棋盘一共8列，刚好放置8个皇后。

	1. 在第0列找到一个位置放置第1个皇后；
	2. 在第1列找到一个位置放置第2个皇后；
	3. 在第2列找到一个位置放置第3个皇后；
	4. 对第3，4，5，6，7，8列都执行放置操作；
	5. 当执行完“在第7列找到一个放置第8个皇后”这一操作完毕后，问题求解完毕。
*/

func Queen() {
	//用一维数组 数组下标i来表示第i行来放置第i+1个皇后 value值表示多少列(从0开始)
	board := make([]int, 8)
	solve(board, 0)
}

//在棋盘上的第row放第i个皇后 从0开始
func solve(board []int, n int) {
	if n == 8 {
		printBoard(board)
		return
	}

	// 遍历当前行的每一列
	for col := 0; col < 8; col++ {
		////先把当前的子n 放在i这个位置上  isSafe(board, n)下面会判断是否冲突，如果冲突的话，col会+1，再赋值给board[n]，然后再判断是否冲突
		board[n] = col
		// 检查皇后是否可以放置在该位置
		if isSafe(board, n) {
			//如果n个皇后不冲突 继续放置下一个皇后
			solve(board, n+1)
		}
	}
}

// 检查皇后是否可以放置在该位置
func isSafe(board []int, n int) bool {
	for i := 0; i < n; i++ {
		//board[i] == board[n] 表示n+1个皇后不能和以前的放置皇后在同一列上
		// abs(board[n]-board[i]) == abs(n-i) 统一斜线表示2个皇后位置处于45度角,说明2个皇后的横坐标之差等于纵坐标之差
		if board[i] == board[n] || abs(board[n]-board[i]) == abs(n-i) {
			return false
		}
	}
	return true
}

// 打印棋盘
func printBoard(board []int) {
	for i := 0; i < 8; i++ {
		fmt.Print(board[i], "  ")
	}
	fmt.Println()
}

//goLang没有整数的取绝对值方法  自己写一个
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
