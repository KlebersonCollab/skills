package main

import (
	"fmt"
	"strings"
	"github.com/samber/lo"
)

func main() {
	// Sample data.
	names := []string{"gopher", "ferris", "moby", "tux", "gopher"}

	// 1. Uniq: Remove duplicates.
	uniqueNames := lo.Uniq(names)
	fmt.Printf("Unique: %v\n", uniqueNames)

	// 2. Map: Transform elements.
	upperNames := lo.Map(uniqueNames, func(name string, _ int) string {
		return strings.ToUpper(name)
	})
	fmt.Printf("Upper: %v\n", upperNames)

	// 3. Filter: Keep elements matching a condition.
	longNames := lo.Filter(upperNames, func(name string, _ int) bool {
		return len(name) > 4
	})
	fmt.Printf("Long names (>4): %v\n", longNames)

	// 4. GroupBy: Group elements by a key.
	grouped := lo.GroupBy(names, func(name string) string {
		return string(name[0])
	})
	fmt.Printf("Grouped by initial: %v\n", grouped)

	// 5. Find: Find the first element matching a condition.
	found, ok := lo.Find(names, func(name string) bool {
		return name == "moby"
	})
	if ok {
		fmt.Printf("Found: %s\n", found)
	}
}
