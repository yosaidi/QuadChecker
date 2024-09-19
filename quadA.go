package main

import (
	"fmt"
	"os"
	"strconv"
)

func HeadFoot(widthX int) string {
	mutVar1 := ""
	for i := 0; i < widthX; i++ {
		if i == 0 || i == (widthX-1) {
			mutVar1 += "o"
		} else {
			mutVar1 += "-"
		}
	}
	mutVar1 += "\n"
	return mutVar1
}

func Vsides(widthX, heightY int) string {
	mutVar1 := ""
	for i := 1; i < heightY-1; i++ {
		for j := 0; j < widthX; j++ {
			if j == 0 || j == (widthX-1) {
				mutVar1 += "|"
			} else {
				mutVar1 += " "
			}
		}
		mutVar1 += "\n"
	}
	return mutVar1
}

func BuildShape(w, y int) string {
	mutVar1 := ""
	if y == 1 {
		mutVar1 += HeadFoot(w)
	} else {
		mutVar1 += HeadFoot(w)
		mutVar1 += Vsides(w, y)
		mutVar1 += HeadFoot(w)
	}
	return mutVar1
}

func QuadA(x, y int) string {
	if x > 0 && y > 0 {
		return BuildShape(x, y)
	}
	return "error"
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./quadA width height")
		return
	}
	width, err1 := strconv.Atoi(os.Args[1])
	height, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input")
		return
	}
	fmt.Print(QuadA(width, height))
}
