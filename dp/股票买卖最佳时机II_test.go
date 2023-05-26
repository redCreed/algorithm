package dp

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("股票最大盈利：%d", MaxProfit(prices))

}
