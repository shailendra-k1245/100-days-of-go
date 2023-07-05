package main
import ("fmt")

func main() {
	var i,j string = "Hello", "World!"
	var x,y = 10,20
	fmt.Print(i)
	fmt.Print(j)
	fmt.Print("\n")
	// adding a new line
	fmt.Print(i,"\n",j)
	fmt.Print("\n")
	// adding a space
	fmt.Print(i," ",j)
	fmt.Print("\n")
	// Print adds a space if varibales are int
	fmt.Print(x,y)
	fmt.Print("\n")
	// Whitespace is added when using Println function
	fmt.Println(i,j)
	fmt.Printf("Type of i is %T and value is %v",i,i)
	fmt.Print("\n")
	fmt.Printf("Type of j is %T and value is %v",j,j)
}