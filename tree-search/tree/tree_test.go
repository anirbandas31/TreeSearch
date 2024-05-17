package tree

import (
	"testing"
)

func TestCheckDuplicateIDs(t *testing.T) {
	// Test case 1: Tree with duplicates at different levels
	root1 := &Node{Value: 1}
	root1.Children = []*Node{
		{Value: 2, Children: []*Node{
			{Value: 4, Children: nil},
			{Value: 5, Children: nil},
		}},
		{Value: 3, Children: []*Node{
			{Value: 4, Children: nil}, // duplicate at level 3
		}},
	}

	value, level := CheckDuplicateIDs(root1)
	if value == nil || *value != 4 || level != 3 {
		if value == nil {
			t.Errorf("Expected (4, 3), got (nil, %d)", level)
		} else {
			t.Errorf("Expected (4, 3), got (%d, %d)", *value, level)
		}
	}

	// Test case 2: Tree with no duplicates
	root2 := &Node{Value: 1}
	root2.Children = []*Node{
		{Value: 2, Children: nil},
		{Value: 3, Children: nil},
	}

	value, level = CheckDuplicateIDs(root2)
	if value != nil || level != 0 {
		if value == nil {
			t.Errorf("Expected (nil, 0), got (nil, %d)", level)
		} else {
			t.Errorf("Expected (nil, 0), got (%d, %d)", *value, level)
		}
	}

	// Test case 3: Tree with duplicate at the root level
	root3 := &Node{Value: 1}
	root3.Children = []*Node{
		{Value: 1, Children: nil}, // duplicate at level 1
	}

	value, level = CheckDuplicateIDs(root3)
	if value == nil || *value != 1 || level != 1 {
		if value == nil {
			t.Errorf("Expected (1, 1), got (nil, %d)", level)
		} else {
			t.Errorf("Expected (1, 1), got (%d, %d)", *value, level)
		}
	}
}
