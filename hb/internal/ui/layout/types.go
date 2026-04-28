package layout

// FlexDirection defines the direction of the main axis.
type FlexDirection int

const (
	Column FlexDirection = iota
	Row
)

// JustifyContent defines alignment along the main axis.
type JustifyContent int

const (
	JustifyFlexStart JustifyContent = iota
	JustifyCenter
	JustifyFlexEnd
	JustifySpaceBetween
	JustifySpaceAround
	JustifySpaceEvenly
)

// AlignItems defines alignment along the cross axis.
type AlignItems int

const (
	AlignAuto AlignItems = iota
	AlignFlexStart
	AlignCenter
	AlignFlexEnd
	AlignStretch
)

// EdgeValues holds values for top, right, bottom, left.
type EdgeValues struct {
	Top, Right, Bottom, Left int
}

// Style defines the Flexbox styling for a Node.
type Style struct {
	Direction     FlexDirection
	Justify       JustifyContent
	Align         AlignItems
	Padding       EdgeValues
	Margin        EdgeValues
	Width, Height int // Optional fixed dimensions (0 = auto)
	FlexGrow      float64
	FlexShrink    float64
}

// LayoutResult stores the computed position and size.
type LayoutResult struct {
	X, Y          int
	Width, Height int
	lastW, lastH  int // Cache for optimization
}

// Node is a layout element in the tree.
type Node struct {
	Style    Style
	Children []*Node
	Result   LayoutResult
}

// NewNode creates a new layout node with default styles.
func NewNode(style Style) *Node {
	return &Node{
		Style:    style,
		Children: []*Node{},
	}
}

// AddChild adds a child node to the container.
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}
