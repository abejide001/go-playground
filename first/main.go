package main

import (
	_ "flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	_ "reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var packet int

func main() {
	a, b := 5, 6
	fmt.Println(b, a)
	_, name := path.Split("cc/main.css")
	fmt.Println(name)
	force := 2.4
	speed := 11
	val := (speed / int(force))
	fmt.Println(val)
	fmt.Println(string(101))
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%+d\n", 30)
	arg := os.Args[2]
	x, _ := strconv.ParseFloat(arg, 16)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", os.Args[1])
	fmt.Sprint("yolo")
	y := strings.ToUpper(os.Args[1])
	fmt.Println(y + "!!!")
	fmt.Println(utf8.RuneCountInString("lolololo"))
	var city string
	city = "abuja"
	rev := strings.Split(city, "")
	sort.Sort(sort.Reverse(sort.StringSlice(rev)))
	fmt.Println(rev)
	const metre = iota
	fmt.Printf("%T\n", metre)

	switch value := 100; value {
	case 100, 300:
		fmt.Println("fmt")
	default:
		fmt.Println("sff")
	}
	words := "world"
	revi := strings.Split(words, "")
	var joined string
	for i := (len(rev) - 1); i >= 0; i-- {
		joined += revi[i]
	}
	fmt.Println(joined)
	i := 1
	var sum int
	for {
		if i > 3 {
			break
		}
		sum += i
		i++
	}
	fmt.Printf("%d\n", sum)
	rand.Seed(time.Now().UnixNano())
	pack := [...]string{
		1: "femi,",
		0: "dfgf",
		2: "sfdfdfd",
	}
	fmt.Printf("%v\n", pack)
	seed := rand.Intn(len(pack))
	fmt.Println(pack[seed])
	comp1 := [...]string{"sfdfd", "sdfdfdf", "femi"}
	comp2 := [...]string{"sfdfd", "sdfdfdf", "sdf"}
	third := len(comp1) - 1
	fmt.Println("len is,", comp1[third])
	fmt.Println(comp1 == comp2)
	// replacing items in an array
	thisArr := []int{1, 2, 3, 4, 5, 6}
	thisArr = append(thisArr[0:2])
	fmt.Println(thisArr)
	arr1 := []float32{1, 2, 3, 5, 6}
	arr1 = append(arr1[:4], arr1[1:]...) // this removes the item in the first index
	fmt.Println("returned array", arr1)
	var ac = 's'
	fmt.Println(ac)
	var run int
	run = 'a'
	fmt.Println(run)
	var byt = []byte{115, 97, 200}
	fmt.Println(string(byt))
	var arr5 *int
	var integer = 1
	arr5 = &integer
	fmt.Println(arr5)
	printThis, _ := tryFunc("dfdgdfg")
	fmt.Println(printThis)
	var mapper = map[interface{}]interface{}{1: true}
	for i, v := range mapper {
		fmt.Println(i, v)
	}
	litte := Person{
		name:  "sfdfd",
		color: map[string]int{"dsfsf": 22},
	}
	fmt.Println(litte.color["dsfsf"])
	const runner = "dsf"
	fmt.Println(runner)
	var getPerson []Person
	fmt.Printf("%#v\n", getPerson)
	for i, v := range getPerson {
		fmt.Println("sfdgd", i, v)
	}
	type eyan string
	var e eyan
	e = "dssf"
	info := fmt.Sprintf("%s", e)
	fmt.Println(info)
	initArr := []string{"20"}
	fmt.Println(initArr)
	mapped := make(map[string]string, 2)
	fmt.Printf("%#v\n", mapped)
	byteString := []byte{200, 211, 100}
	for _, value := range byteString {
		fmt.Println("string values", string(value))
	}
}

// Person struct
type Person struct {
	name  string
	color map[string]int
}

func tryFunc(catt string) ([]string, error) {
	return strings.Split(catt, " "), nil
}
