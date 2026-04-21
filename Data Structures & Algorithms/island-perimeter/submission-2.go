func islandPerimeter(grid [][]int) int {
	row:=len(grid)
	col:=len(grid[0])

	var dfs func(i,j int) int
	dfs=func(i,j int)int{
		if i<0||j<0||i>=row||j>=col{
			return 1
		}

		if grid[i][j]==0{
			return 1
		}
		if grid[i][j]==-1{
			return 0
		}

		grid[i][j]=-1

		return dfs(i+1,j)+dfs(i-1,j)+dfs(i,j+1)+dfs(i,j-1)
		
	}

	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if grid[i][j]==1{
				return dfs(i,j)
			}
		}
	}

	

	return 0
}
