package main

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	dirs := strings.Split(path, "/")

	for _, dir := range dirs {
		switch dir {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, dir)
		}
	}

	simplifiedPath := "/" + strings.Join(stack, "/")
	return simplifiedPath
}