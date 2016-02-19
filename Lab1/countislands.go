package main

func CountIslands(grid [][]int) int {

	count:= 0
	rows:= len(grid)
	col:= len(grid[0])
	visited:= make ([][]int, rows)

	for i:= range visited {
		visited[i] = make([]int, col)
	}

	for i:=0 ; i< rows ;i++ {
		for j:=0; j< col ; j++ {
				visited[i][j] = 0
			}
	}

	for i:=0 ; i< rows ;i++ {
		for j:=0; j< col ; j++ {
				if (visited[i][j]==0 && grid[i][j] == 1) {
					check(grid,visited,i,j,rows,col)
					count ++
				}
			}
		}

	return count

}

func check(grid [][]int, visited [][]int, i int, j int, r int, c int) {

  if i < 0 {
    return
  }
  if i >= r {
    return
  }
  if j < 0 {
    return
  }
  if j >= c {
    return
  }
  if visited[i][j] == 1 {
    return
  }
  if grid[i][j] == 0 {
    return
  }
  visited[i][j] = 1
  check(grid,visited,i+1,j,r,c)
  check(grid,visited,i,j+1,r,c)
  check(grid,visited,i-1,j,r,c)
  check(grid,visited,i,j-1,r,c)
}
