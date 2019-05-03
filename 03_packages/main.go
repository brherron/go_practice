package main

import (
	"fmt"
	"math"

	"github.com/beauherron/go-crash-course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(3.5))
	fmt.Println(strutil.Reverse("olleH"))
}
