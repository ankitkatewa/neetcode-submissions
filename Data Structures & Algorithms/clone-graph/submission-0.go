/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {

	if node==nil{
		return nil
	}
	visited :=make(map[*Node]*Node)
	return dfs(node,visited)


}

func dfs(node *Node,visited map[*Node]*Node)*Node{

	if value,ok:= visited[node]; ok{
		return value
	}

	newNode:=&Node{
	Val:node.Val,
	Neighbors:[]*Node{},
	}

   visited[node]=newNode

   for _,nei:=range node.Neighbors{
		newNode.Neighbors=append(newNode.Neighbors,dfs(nei,visited))
   }

   return newNode
}
