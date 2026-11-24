package main

import "fmt"

// LargestRectangle returns the area of the largest axis-aligned rectangle
// fitting under a histogram. One pass with a monotonic stack: a bar is
// finalized the moment a shorter bar appears, its width spanning back to
// the previous shorter bar.
func LargestRectangle(heights []int) int {
	best := 0
	stack := []int{} // indexes of bars with strictly increasing heights
	for i := 0; i <= len(heights); i++ {
		h := 0 // sentinel bar of height 0 flushes the stack at the end
		if i < len(heights) {
			h = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			if area := heights[top] * width; area > best {
				best = area
			}
		}
		stack = append(stack, i)
	}
	return best
}

func main() {
	fmt.Println(LargestRectangle([]int{2, 1, 5, 6, 2, 3}))
}
