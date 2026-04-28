package layout

import "math"

// Calculate computes the layout for the node and its children.
func (n *Node) Calculate(availableWidth, availableHeight int) {
	// Optimization: Skip if dimensions haven't changed and we have a result
	if n.Result.lastW == availableWidth && n.Result.lastH == availableHeight && n.Result.Width > 0 {
		return
	}
	n.Result.lastW = availableWidth
	n.Result.lastH = availableHeight

	// 1. Initial dimensions based on style or availability
	width := n.Style.Width
	if width == 0 || width > availableWidth {
		width = availableWidth
	}
	height := n.Style.Height
	if height == 0 || height > availableHeight {
		height = availableHeight
	}

	// Apply Margins (Reduce available content size)
	width -= n.Style.Margin.Left + n.Style.Margin.Right
	height -= n.Style.Margin.Top + n.Style.Margin.Bottom

	n.Result.Width = width
	n.Result.Height = height

	// Content area (excluding padding)
	contentW := width - n.Style.Padding.Left - n.Style.Padding.Right
	contentH := height - n.Style.Padding.Top - n.Style.Padding.Bottom

	if len(n.Children) == 0 {
		return
	}

	// 2. Measure children
	totalFlexGrow := 0.0
	for _, child := range n.Children {
		totalFlexGrow += child.Style.FlexGrow
		// For now, assume children take auto size if no flex grow
	}

	// 3. Position children based on Direction
	currentX := n.Style.Padding.Left + n.Style.Margin.Left
	currentY := n.Style.Padding.Top + n.Style.Margin.Top

	if n.Style.Direction == Row {
		// Calculate fixed sizes and flex factor
		fixedSize := 0
		for _, child := range n.Children {
			if child.Style.FlexGrow == 0 {
				fixedSize += child.Style.Width
			}
		}

		remainingW := contentW - fixedSize
		flexUnit := 0.0
		if totalFlexGrow > 0 {
			flexUnit = float64(remainingW) / totalFlexGrow
			remainingW = 0 // FlexGrow consumes all remaining space
		}

		// Calculate spacing based on JustifyContent
		spacing := 0
		offset := 0
		if remainingW > 0 {
			switch n.Style.Justify {
			case JustifyCenter:
				offset = remainingW / 2
			case JustifyFlexEnd:
				offset = remainingW
			case JustifySpaceBetween:
				if len(n.Children) > 1 {
					spacing = remainingW / (len(n.Children) - 1)
				}
			case JustifySpaceAround:
				spacing = remainingW / len(n.Children)
				offset = spacing / 2
			case JustifySpaceEvenly:
				spacing = remainingW / (len(n.Children) + 1)
				offset = spacing
			}
		}

		currentX += offset

		for _, child := range n.Children {
			childW := 0
			if child.Style.FlexGrow > 0 {
				childW = int(math.Floor(child.Style.FlexGrow * flexUnit))
			} else if child.Style.Width > 0 {
				childW = child.Style.Width
			} else {
				childW = contentW / len(n.Children)
			}

			// AlignItems (Cross Axis)
			childH := contentH
			childY := currentY
			if child.Style.Height > 0 && child.Style.Height < contentH {
				childH = child.Style.Height
				switch n.Style.Align {
				case AlignCenter:
					childY += (contentH - childH) / 2
				case AlignFlexEnd:
					childY += (contentH - childH)
				}
			}

			child.Result.X = currentX
			child.Result.Y = childY
			child.Result.Width = childW
			child.Result.Height = childH

			// Recurse
			child.Calculate(childW, childH)

			currentX += childW + spacing
		}
	} else {
		// Column logic (Vertical)
		remainingH := contentH
		flexUnit := 0.0
		if totalFlexGrow > 0 {
			flexUnit = float64(remainingH) / totalFlexGrow
			remainingH = 0
		}

		// Vertical Justify
		spacing := 0
		offset := 0
		if remainingH > 0 {
			switch n.Style.Justify {
			case JustifyCenter:
				offset = remainingH / 2
			case JustifyFlexEnd:
				offset = remainingH
			}
		}
		currentY += offset

		for _, child := range n.Children {
			childH := 0
			if child.Style.FlexGrow > 0 {
				childH = int(math.Floor(child.Style.FlexGrow * flexUnit))
			} else if child.Style.Height > 0 {
				childH = child.Style.Height
			} else {
				childH = contentH / len(n.Children)
			}

			// AlignItems (Cross Axis)
			childW := contentW
			childX := currentX
			if child.Style.Width > 0 && child.Style.Width < contentW {
				childW = child.Style.Width
				switch n.Style.Align {
				case AlignCenter:
					childX += (contentW - childW) / 2
				case AlignFlexEnd:
					childX += (contentW - childW)
				}
			}

			child.Result.X = childX
			child.Result.Y = currentY
			child.Result.Width = childW
			child.Result.Height = childH

			// Recurse
			child.Calculate(childW, childH)

			currentY += childH + spacing
		}
	}
}
