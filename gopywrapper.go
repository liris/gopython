package py
/*
#include <Python.h>
*/
import "C";
import "unsafe";


/* TODO:
 * Py_DECREF macro in object.h
 * Py_XDECREF
 */

// object.h

type Object struct {
	cptr *C.PyObject;
}

func newObject(p *C.PyObject) *Object {
	self := new(Object);
	self.cptr = p;
	return self;
}

func (self *Object) GetAttrString(object *Object, value string) *Object {
	result := C.PyObject_GetAttrString(self.cptr, C.CString(value));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) CallObject(args *Object) *Object {
	result := C.PyObject_CallObject(self.cptr, args.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Callable_Check() int {
	result := C.PyCallable_Check(self.cptr);
	return int(result);
}

// tupleobject.h

func Tuple_New(size int) *Object {
	self := new(Object);
	result := C.PyTuple_New((C.Py_ssize_t)(size));
	self.cptr = result;
	return self;
}

func (self *Object) Tuple_SetItem(index int, item *Object) int {
	result := C.PyTuple_SetItem(self.cptr, (C.Py_ssize_t)(index), item.cptr);
	return int(result);
}

// stringobject.h

func String_FromString(s string) *Object {
	result := C.PyString_FromString(C.CString(s));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func Int_FromLong(value int) *Object {
	result := C.PyInt_FromLong((C.long)(value));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}	

// intobject.h

func (self *Object) Int_AsLong() int {
	result := C.PyInt_AsLong(self.cptr);
	return int(result);
}

// pythonrun.h
func Initialize() {
	C.Py_Initialize();
}

func Finalize() {
	C.Py_Finalize();
}


func Run_SimpleString(script string) int {
	return int(C.PyRun_SimpleStringFlags(C.CString(script), nil));
}

func Err_Print() {
	C.PyErr_Print();
}

// import.h

func Import_Import(name *Object) *Object {
	result := C.PyImport_Import(name.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

/* TODO:
 * Py_DECREF
 * Py_XDECREF
 */

