package main

/*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，
其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：
1 <= numCourses <= 105
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同
*/

/*
深度优先搜索
时间复杂度：O(m+n)
空间复杂度：O(m+n)
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges, visited, valid := make([][]int, numCourses), make([]int, numCourses), true
	var dfs func(int)

	dfs = func(i int) {
		visited[i] = 1
		for _, v := range edges[i] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[i] = 2
	}

	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0]) //组成邻接表的存储结构
	}
	for i := 0; i < numCourses && valid; i++ { //当valid=false时，直接退出
		if visited[i] == 0 {
			dfs(i)
		}
	}

	return valid
}

/*
广度优先搜索
时间复杂度：O(m+n)
空间复杂度：O(m+n)
*/
func canFinish2(numCourses int, prerequisites [][]int) bool {
	edges, indeg, result, queue := make([][]int, numCourses),
		make([]int, numCourses), make([]int, 0), make([]int, 0)
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0]) //组成邻接表的存储结构
		indeg[v[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return len(result) == numCourses
}

func main() {

}
