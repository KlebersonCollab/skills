package layout

import "testing"

func TestLayoutRowSpaceBetween(t *testing.T) {
	root := NewNode(Style{
		Direction: Row,
		Justify:   JustifySpaceBetween,
		Padding:   EdgeValues{Left: 1, Right: 1},
	})

	child1 := NewNode(Style{FlexGrow: 1})
	child2 := NewNode(Style{FlexGrow: 1})

	root.AddChild(child1)
	root.AddChild(child2)

	root.Calculate(100, 10)

	if root.Result.Width != 100 {
		t.Errorf("Expected root width 100, got %d", root.Result.Width)
	}

	// Content width = 100 - 1 (left) - 1 (right) = 98
	// FlexGrow 1+1 = 2. flexUnit = 98 / 2 = 49
	if child1.Result.Width != 49 {
		t.Errorf("Expected child1 width 49, got %d", child1.Result.Width)
	}
	if child1.Result.X != 1 {
		t.Errorf("Expected child1 X 1, got %d", child1.Result.X)
	}

	if child2.Result.X != 50 { // 1 + 49
		t.Errorf("Expected child2 X 50, got %d", child2.Result.X)
	}
}

func TestLayoutColumn(t *testing.T) {
	root := NewNode(Style{
		Direction: Column,
	})

	header := NewNode(Style{FlexGrow: 1})
	body := NewNode(Style{FlexGrow: 3})

	root.AddChild(header)
	root.AddChild(body)

	root.Calculate(80, 20)

	if header.Result.Height != 5 { // 1/4 of 20
		t.Errorf("Expected header height 5, got %d", header.Result.Height)
	}
	if body.Result.Height != 15 { // 3/4 of 20
		t.Errorf("Expected body height 15, got %d", body.Result.Height)
	}
}
