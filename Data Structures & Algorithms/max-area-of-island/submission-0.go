func maxAreaOfIsland(grid [][]int) int {
	rows:=len(grid)
	cols:=len(grid[0])
    visited:=make([][]int,rows)

	for i:=0;i<rows;i++{
		visited[i]=make([]int,cols)
	}

	var dfs func(i, j int) int
	dfs=func(i,j int)int{
		if i<0 || j<0 || i>=rows || j>=cols{
			return 0
		}
		if visited[i][j]==1{
			return 0
		}
		if grid[i][j]==0{
			return 0
		}

		visited[i][j]=1
		// if grid[i][j]==1{
		// 	return 1
		// }

		return 1+dfs(i+1,j)+dfs(i-1,j)+dfs(i,j+1)+dfs(i,j-1)
	}

	maxArea:=0
	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			if grid[i][j]==1 && visited[i][j]==0{
				area:=dfs(i,j)
				if area>maxArea{
					maxArea=area
				}
			}
		}

	}

	return maxArea


}
