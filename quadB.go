package main

import (
	"fmt"
	"os"
	"strconv"
)

// Print the head of the shape
func Head2(widthX int) {
	for i := 0; i < widthX; i++ {
		if i == 0 {
			fmt.Printf("/")
		} else if i == widthX-1 {
			fmt.Printf("\\")
		} else {
			fmt.Printf("*")
		}
	}
	fmt.Printf("\n")
}

// Print the footer of the shape
func Foot2(widthX int) {
	for i := 0; i < widthX; i++ {
		if i == 0 {
			fmt.Printf("\\")
		} else if i == widthX-1 {
			fmt.Printf("/")
		} else {
			fmt.Printf("*")
		}
	}
	fmt.Printf("\n")
}

// Print the body of the shape
func Vsides2(widthX, heightY int) {
	for i := 1; i < heightY-1; i++ {
		for j := 0; j < widthX; j++ {
			if j == 0 || j == (widthX-1) {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

// Function to build the whole shape
func BuildShape2(w, y int) {
	if y == 1 {
		Head2(w)
	} else {
		Head2(w)
		Vsides2(w, y)
		Foot2(w)
	}
}

// Function to handle input and call BuildShape2
func QuadB(x, y int) {
	if x > 0 && y > 0 {
		BuildShape2(x, y)
	} else {
		fmt.Println("error")
	}
}

// Main function to handle command-line arguments
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program1 width height")
		return
	}

	width, err1 := strconv.Atoi(os.Args[1])
	height, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input")
		return
	}

	QuadB(width, height)
}
