package main

import "fmt"

/*
机器人在一个无限大小的 XY 网格平面上行走，从点(0, 0) 处开始出发，面向北方。
该机器人可以接收以下三种类型的命令 commands ：
	-2 ：向左转90 度
	-1 ：向右转 90 度
	1 <= x <= 9 ：向前移动x个单位长度
在网格上有一些格子被视为障碍物obstacles 。第 i 个障碍物位于网格点 obstacles[i] = (xi, yi)
机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分
返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）

注意：
	北表示 +Y 方向。
	东表示 +X 方向。
	南表示 -Y 方向。
	西表示 -X 方向。

示例 1：
输入：commands = [4,-1,3], obstacles = []
输出：25

解释：
机器人开始位于 (0, 0)：
1. 向北移动 4 个单位，到达 (0, 4)
2. 右转
3. 向东移动 3 个单位，到达 (3, 4)
距离原点最远的是 (3, 4) ，距离为 32 + 42 = 25

示例2：
输入：commands = [4,-1,4,-2,4], obstacles = [[2,4]]
输出：65

解释：
机器人开始位于 (0, 0)：
1. 向北移动 4 个单位，到达 (0, 4)
2. 右转
3. 向东移动 1 个单位，然后被位于 (2, 4) 的障碍物阻挡，机器人停在 (1, 4)
4. 左转
5. 向北走 4 个单位，到达 (1, 8)
距离原点最远的是 (1, 8) ，距离为 12 + 82 = 65

提示：
commands[i] 只会存在于 [-2,-1,1,2,3,4,5,6,7,8,9] 中
*/

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func robotSim(commands []int, obstacles [][]int) int {
	// 判断障碍物在该位置是否存在
	obstaclesMap := make(map[[2]int]bool)
	for _, v := range obstacles {
		if len(v) == 2 {
			obstaclesMap[[2]int{v[0], v[1]}] = true
		}
	}
	// []int{北，东，南，西}
	dirX := [4]int{0, 1, 0, -1}
	dirY := [4]int{1, 0, -1, 0}
	x, y, dir, distance := 0, 0, 0, 0 // 当前位置，当前方向，欧式距离
	for _, v := range commands {
		switch v {
		case -2: // 左转
			dir = (dir + 3) % 4
		case -1: // 右转
			dir = (dir + 1) % 4
		default: // 前进
			for i := 0; i < v; i++ { // 向dir方向走v步
				nextX := x + dirX[dir]
				nextY := y + dirY[dir]
				// 判断是否存在障碍物
				if _, ok := obstaclesMap[[2]int{nextX, nextY}]; ok {
					break
				}
				x = nextX
				y = nextY
				// 计算最大欧氏距离
				distance = max(distance, x*x+y*y)
			}
		}

	}
	return distance
}

func main() {
	commands := []int{4, -1, 4, -2, 4}
	var obstacles [][]int
	obstacles = append(obstacles, []int{2, 4})
	sim := robotSim(commands, obstacles)
	fmt.Println(sim)
}
