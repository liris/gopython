package main

import (
	"py";
	"fmt";
	)

func main() {
	fmt.Print("init\n");
	py.Initialize();
	fmt.Print("init done\n");
	result := py.Run_SimpleString("print 'hello'");
	fmt.Printf("script done %d\n", result);
	sobj := py.String_FromString("hoge");
	fmt.Printf("sobj %t\n", sobj);
	py.Finalize();
	fmt.Print("finalize done\n");
}

