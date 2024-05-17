package main

import (
	"fmt"
	"tree-search/tree"
)

func main() {
	// Example 1: Tree with duplicates at different levels
	root1 := &tree.Node{Value: 1}
	root1.Children = []*tree.Node{
		{Value: 2, Children: []*tree.Node{
			{Value: 4, Children: nil},
			{Value: 5, Children: nil},
		}},
		{Value: 3, Children: []*tree.Node{
			{Value: 4, Children: nil}, // duplicate at level 3
		}},
	}

	value, level := tree.CheckDuplicateIDs(root1)
	if value != nil {
		fmt.Printf("Example 1: Found duplicate value %d at level %d\n", *value, level)
	} else {
		fmt.Println("Example 1: No duplicates found")
	}

	// Example 2: Tree with no duplicates
	root2 := &tree.Node{Value: 1}
	root2.Children = []*tree.Node{
		{Value: 2, Children: nil},
		{Value: 3, Children: nil},
	}

	value, level = tree.CheckDuplicateIDs(root2)
	if value != nil {
		fmt.Printf("Example 2: Found duplicate value %d at level %d\n", *value, level)
	} else {
		fmt.Println("Example 2: No duplicates found")
	}

	// Example 3: Tree with duplicate at the root level
	root3 := &tree.Node{Value: 1}
	root3.Children = []*tree.Node{
		{Value: 1, Children: nil}, // duplicate at level 1
	}

	value, level = tree.CheckDuplicateIDs(root3)
	if value != nil {
		fmt.Printf("Example 3: Found duplicate value %d at level %d\n", *value, level)
	} else {
		fmt.Println("Example 3: No duplicates found")
	}

	// Example 4: Larger tree with multiple levels and duplicates
	root4 := &tree.Node{Value: 10}
	root4.Children = []*tree.Node{
		{Value: 20, Children: []*tree.Node{
			{Value: 30, Children: nil},
			{Value: 40, Children: nil},
		}},
		{Value: 20, Children: []*tree.Node{ // duplicate at level 2
			{Value: 50, Children: nil},
			{Value: 60, Children: nil},
		}},
		{Value: 30, Children: []*tree.Node{ // duplicate at level 2
			{Value: 70, Children: []*tree.Node{
				{Value: 30, Children: nil}, // duplicate at level 4
			}},
		}},
	}

	value, level = tree.CheckDuplicateIDs(root4)
	if value != nil {
		fmt.Printf("Example 4: Found duplicate value %d at level %d\n", *value, level)
	} else {
		fmt.Println("Example 4: No duplicates found")
	}
}
