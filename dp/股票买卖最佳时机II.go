package dp

/*
	给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

	在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

	返回 你能获得的 最大 利润 。
*/

func MaxProfit(prices []int) int {
	var len = len(prices)
	var profit int
	//当天买当天卖没有意义的
	for i := range prices {
		if i+1 < len {
			if prices[i] < prices[i+1] {
				profit += prices[i+1] - prices[i]
			}
		}
	}
	return profit
}
