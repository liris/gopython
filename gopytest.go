package main

import (
	"py";
	"fmt";
	)

func runSimpleString() {
	fmt.Printf("--- runSimpleString ------------------------------\n");
	result := py.Run_SimpleString("print 'hello'");
	fmt.Printf("result %d\n\n", result);
	fmt.Printf("--- runSimpleString done ---\n\n");
}

func setArgs(pArgs *py.Object, index int, value int) bool {
	pValue := py.Int_FromInt(value);
	if pValue == nil {
		fmt.Printf("Value : %d\n", value);
		py.Err_Print();
		return false;
	}
	pArgs.Tuple_SetItem(index, pValue);
	return true;
}

func runFunc(pythonfile string, funcname string, arg []int) {
	fmt.Printf("--- runFunc ------------------------------\n");
	pName := py.String_FromString(pythonfile);
	pModule := py.Import_Import(pName);
	fmt.Printf("import done \n");
	pName.DecRef();
	if pModule != nil {
		fmt.Printf("pModule %t\n", pModule);
		pFunc := pModule.GetAttrString(funcname);
		if pFunc != nil && pFunc.Callable_Check() != 0 {
			pArgs := py.Tuple_New(2);
			if setArgs(pArgs, 0, arg[0]) == false {
				pArgs.DecRef();
				pModule.DecRef();
				fmt.Printf("Cannot Convert Argument\n");
				return;
			}
			if setArgs(pArgs, 1, arg[1]) == false {
				pArgs.DecRef();
				pModule.DecRef();
				fmt.Printf("Cannot Convert Argument\n");
				return;
			}
			pValue := pFunc.CallObject(pArgs);
			pArgs.DecRef();
			if pValue != nil {
				fmt.Printf("Result of call: %d\n", pValue.Int_AsInt64());
				pValue.DecRef();
			} else {
				pFunc.DecRef();
				pModule.DecRef();
				py.Err_Print();
				fmt.Printf("Call Failed\n");
			}
		} else {
			if py.Err_Occurred() != nil {
				py.Err_Print();
			}
			fmt.Printf("Cannot find function %s\n", funcname);
		}
		py.XDecRef(pFunc);
		pModule.DecRef();
	} else {
		py.Err_Print();
		fmt.Print("error\n\n");
	}
	fmt.Printf("--- runFunc done ---\n\n");
}


func main() {
	py.Initialize();
		
	runSimpleString();
	runFunc("multiply", "multiply", []int{4, 5});
	
	py.Finalize();
	fmt.Print("finalize done\n");
}

