package main

import (
	"fmt"
	"os"
	"strconv"
)

const program = `package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}
`

func generatePackage(name string) {
	if err := os.Mkdir(name, os.FileMode(0755)); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	f, err := os.Create(name + "/" + name + ".go")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	f.WriteString(program)
	if err := f.Close(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	switch len(os.Args) {
	case 2:
		packageName := os.Args[1]
		generatePackage(packageName)
	case 3:
		name := os.Args[1]
		count, _ := strconv.Atoi(os.Args[2])
		digitNum := 0
		n := len(name)
		for i := n - 1; i >= 0; i-- {
			if name[i] >= '0' && name[i] <= '9' {
				digitNum++
			} else {
				break
			}
		}
		format := fmt.Sprintf("%%s%%0%dd", digitNum)
		start, _ := strconv.Atoi(name[n-digitNum:])
		for i := 0; i < count; i++ {
			packageName := fmt.Sprintf(format, name[:n-digitNum], start+i)
			generatePackage(packageName)
		}
	default:
		fmt.Println("gdp <package name> [count]")
	}
}
