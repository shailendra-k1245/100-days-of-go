package main
import ("fmt")
func main(){
	var i = 15.5
	var j = 15
	var txt = "Hello world!"
	var t2 = "Hello"

	// General formatting verbs
	fmt.Printf("%v\n",i)
	fmt.Printf("%#v\n",i)
	fmt.Printf("%v%%\n",i)
	fmt.Printf("%T\n",i)

	fmt.Printf("%v\n",txt)
	fmt.Printf("%#v\n",txt)
	fmt.Printf("%T\n",txt)
	fmt.Print("\n")

	// Integer formatting verbs
	fmt.Printf("%b\n", j)
	fmt.Printf("%d\n", j)
	fmt.Printf("%+d\n", j)
	fmt.Printf("%o\n", j)
	fmt.Printf("%O\n", j)
	fmt.Printf("%x\n", j)
	fmt.Printf("%X\n", j)
	fmt.Printf("%#x\n", j)
	fmt.Printf("%4d\n", j)
	fmt.Printf("%-4d\n", j)
	fmt.Printf("%04d\n", j)

	// String formatting Vervs
	stringFormattingVerbs(t2)
}

func stringFormattingVerbs(txt string){
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

}

