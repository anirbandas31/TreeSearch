package tree

// Tree Node
type Node struct {
	Value    int
	Children []*Node  
}

// CheckDuplicateIDs finds the shallowest duplicate node in the tree
func CheckDuplicateIDs(root *Node) (*int, int) {
	if root == nil {
		return nil, 0 // If root is nil, return nil and 0 as there are no nodes
	}

	queue := []*Node{root}          // init the tree root
	levelMap := make(map[int]int)   // map for storing duplicates
	level := 1                      // initialize lvel at 1
	nodeCount := len(queue)         // var for tracking node 

	for nodeCount > 0 { // start a while loop until nodeCount is 0
		for i := 0; i < nodeCount; i++ { // loop over all of the nodes in the current queue
			node := queue[i] // store the current queue
			if levelSeen, exists := levelMap[node.Value]; exists { // check if we came across this value already
				return &node.Value, levelSeen // return the result and current level if yes
			}
			levelMap[node.Value] = level //otherwise store the value and level in the map
			for _, child := range node.Children { // add every child of the node for processing later
				queue = append(queue, child) 
			}
		}
		queue = queue[nodeCount:] // remove processed nodes
		nodeCount = len(queue) // update nodecount for next level
		level++ // go to next level
	}

	return nil, 0 // return nil and 0 if no duplicates
}
