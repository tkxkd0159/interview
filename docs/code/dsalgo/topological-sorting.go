func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
		indegree[p[0]]++
	}

	q := make([]int, 0)
	for i := range numCourses {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

    visited := 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		visited++

		for _, neighbor := range adj[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	return visited == numCourses
}
