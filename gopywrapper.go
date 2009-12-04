package py
/*
#include <Python.h>

 void Py_DecRef(PyObject* obj) { Py_DECREF(obj); }
 void Py_XDecRef(PyObject* obj) { Py_XDECREF(obj); }
 int Arg_ParseTuple(PyObject * self, const char * name) {
     return PyArg_ParseTuple(self, name);
 }
 
 PyObject* BuildIntValue(const char * name, int v) {
     return Py_BuildValue(name, v);
 }
 PyObject* BuildStringValue(const char * name, const char* v) {
     return Py_BuildValue(name, v);
 }
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

func (self *Object) HasAttrString(attr_name string) int {
	result := C.PyObject_HasAttrString(self.cptr, C.CString(attr_name));
	return int(result);
}

func (self *Object) GetAttrString(attr_name string) *Object {
	result := C.PyObject_GetAttrString(self.cptr, C.CString(attr_name));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) HasAttr(attr_name string) int {
	result := C.PyObject_HasAttr(self.cptr, C.CString(attr_name));
	return int(result);
}

// TODO: More Generic Object Interface

func (self *Object) GetAttr(attr_name *Object) *Object {
	result := C.PyObject_GetAttr(self.cptr, attr_name.cptr);
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
	C.Py_DecRef(self.cptr);
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


// intobject.h
func Int_FromInt64(value int64) *Object {
	result := C.PyInt_FromLong(C.long(value));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}	

func (self *Object) Int_AsInt64() int64 {
	result := C.PyInt_AsLong(self.cptr);
	return int64(result);
}

func Int_FromInt(value int) *Object {
	result := C.PyInt_FromLong(C.long(value));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}	

func (self *Object) Int_AsInt() int {
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
// modsupport.h
type CallbackFunction func(self *Object, args *Object) *Object

type MethodDef struct {
	Ml_name string;
	Ml_meth CallbackFunction;
	Ml_flags int;
	Ml_doc string;
}

func (self *Object) Arg_ParseTuple(name string) int {
	result := C.Arg_ParseTuple(self.cptr, C.CString(name));
	return int(result);
}

func BuildIntValue(name string, v int) *Object {
	result := C.BuildIntValue(C.CString(name), C.int(v));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func BuildStringValue(name string, v string) *Object {
	result := C.BuildStringValue(C.CString(name), C.CString(v));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

