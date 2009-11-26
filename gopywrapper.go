package py
/*
#include <Python.h>

 void Py_DecRef(PyObject* obj) { Py_DECREF(obj); }
 void Py_XDecRef(PyObject* obj) { Py_XDECREF(obj); }
*/
import "C";
import "unsafe";

// object.h

type Object struct {
	cptr *C.PyObject;
}

func newObject(p *C.PyObject) *Object {
	if p == nil {
		return nil;
	}
	self := new(Object);
	self.cptr = p;
	return self;
}

func (self *Object) GetAttrString(value string) *Object {
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

func (self *Object) DecRef() {
//	self.cptr.ob_refcnt--;
//	if self.cptr.ob_refcnt == 0 {
		C.Py_DecRef(self.cptr);
//	}
}

func XDecRef(self *Object) {
	C.Py_XDecRef(self.cptr);
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
	result := C.PyInt_FromLong(C.long(value));
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

func Err_Occurred() *Object {
	result := C.PyErr_Occurred();
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

// import.h

func Import_Import(name *Object) *Object {
	result := C.PyImport_Import(name.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

