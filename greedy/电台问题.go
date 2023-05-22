package greedy

import "fmt"

/*
	存在下面需要付费的广播电台，以及广播电台可以覆盖的地区。如何选择最少的电台，能实现区域的全覆盖
    广播台                           覆盖地区
    k1                     北京 上海 天津
	k2                     广州 北京 深圳
	k3                     成都 上海 杭州
    k4                     上海 天津
	k5                     杭州 大连

	* 贪心算法（贪婪算法）是指在对问题进行求解时，在每一步的选择中都选择最优解，从而期望能够导致结果是最优解的算法
	* 贪心算法所得到的结果不一定是最优解，但是一定是相对近似最优解的结果
	* 注意：此处虽然在每一步取得了最优解，但是最终结果不一定是最优解，这也是贪心算法的特性
*/

//已知存在多少电台，及电台对应的覆盖城市集合；并且各个电台所覆盖城市存在部分重复，需要最少几部电台可实现全覆盖
//首先汇总需要覆盖的城市，取各个电台对应城市的并集，作为汇总城市，并全部表示为未覆盖城市
//然后遍历各个电台，统计各个电台在未覆盖的城市中可以覆盖几座城市
//统计完成后，取覆盖城市最多的电台为最优解，即为先用电台，添加到选用集合；同时，从未覆盖城市集合中删除该电台覆盖的城市
//再剩余未覆盖的城市中，重复第3和第4步，直到未覆盖城市数为0，即表示城市已经被全部覆盖
//最终返回选用集合，表示最终选择的电台

func SelectRadioStation() {
	//添加所有电台
	radioMap := make(map[string][]string)
	radioMap["k1"] = []string{"北京", "上海", "天津"}
	radioMap["k2"] = []string{"广州", "北京", "深圳"}
	radioMap["k3"] = []string{"成都", "上海", "杭州"}
	radioMap["k4"] = []string{"上海", "天津"}
	radioMap["k5"] = []string{"杭州", "大连"}

	//获取所有未覆盖的城市 就是所有电台可以覆盖城市的集合
	allCity := getAllCity(radioMap)

	//统计选择的电台
	selected := make([]string, 0)

	for len(allCity) > 0 {
		//记录当前匹配批次中最大城市的电台和数量
		maxKey := ""
		maxCount := 0
		currentCount := 0
		//遍历所有的电台数据
		for k, r := range radioMap {
			//判断当前电台在所有未覆盖的城市的数量，求交集
			currentCount = getIntersection(allCity, r)
			// 如果当前匹配数量大于0, 说明匹配到, 有资格进行记录
			// 大于当前最大值, 则进行记录
			if currentCount > 0 && currentCount > maxCount {
				maxKey = k
				maxCount = currentCount
			}
		}
		//一次处理完成后, 统计出需要记录的key进行记录, 并移除匹配的城市
		if maxKey != "" {
			selected = append(selected, maxKey)
			deleteCity(allCity, radioMap[maxKey])
		}
	}

	fmt.Println("selected:", selected)
}

//删除切片的元素
func deleteCity(allCity map[string]struct{}, city []string) {
	for _, v := range city {
		delete(allCity, v)
	}
}

//求2个切片的交集
func getIntersection(allCity map[string]struct{}, radioCity []string) int {
	count := 0
	for k, _ := range allCity {
		for _, v2 := range radioCity {
			if k == v2 {
				count++
			}
		}
	}
	return count
}

func getAllCity(radioData map[string][]string) map[string]struct{} {
	radioMap := make(map[string]struct{})
	for _, r := range radioData {
		for _, v := range r {
			if _, ok := radioMap[v]; !ok {
				radioMap[v] = struct{}{}
			}
		}
	}

	return radioMap
}
