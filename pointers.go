package main

import (
	"fmt"
)

func passvalue(ival int){
	ival += 1;
}

func passref(ival *int){
	fmt.Println(ival);
	fmt.Println(*ival);
	*ival += 1;
}

func PointerProgram(){
	i := 0
    fmt.Println("initial:", i)	

    passvalue(i)
    fmt.Println("zeroval:", i)

    passref(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)

}