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

 int IterCheck(PyObject *o) {
     return PyIter_Check(o);
 }

 int IndexCheck(PyObject *o) {
   return PyIndex_Check(o);
 }
 PyObject * Mapping_Keys(PyObject* o) {
   return PyMapping_Keys(o);
 }
*/
import "C";
import "unsafe";

/*
 * TODO:
 *  - Py_ssize_t
 */
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

func (self *Object) HasAttr(attr_name *Object) int {
	result := C.PyObject_HasAttr(self.cptr, attr_name.cptr);
	return int(result);
}

func (self *Object) GetAttr(attr_name *Object) *Object {
	result := C.PyObject_GetAttr(self.cptr, attr_name.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) SetAttrString(attr_name string, v *Object) int {
	result := C.PyObject_SetAttrString(self.cptr, C.CString(attr_name), v.cptr);
	return int(result);
}

func (self *Object) SetAttr(attr_name *Object, v *Object) int {
	result := C.PyObject_SetAttr(self.cptr, attr_name.cptr, v.cptr);
	return int(result);
}

func (self *Object) DelAttrString(attr_name string) int {
	result := C.PyObject_SetAttrString(self.cptr, C.CString(attr_name), nil);
	return int(result);
}

func (self *Object) DelAttr(attr_name *Object) int {
	result := C.PyObject_SetAttr(self.cptr, attr_name.cptr, nil);
	return int(result);
}

// XXX:
// do we need 
// PyAPI_FUNC(int) PyObject_Cmp(PyObject *o1, PyObject *o2, int *result);

func (self *Object) Compare(o2 *Object) int {
	result := C.PyObject_Compare(self.cptr, o2.cptr);
	return int(result);
}

func (self *Object) Repr() *Object {
	result := C.PyObject_Repr(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Str() *Object {
	result := C.PyObject_Str(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Unicode() *Object {
	result := C.PyObject_Unicode(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Callable_Check() int {
	result := C.PyCallable_Check(self.cptr);
	return int(result);
}

func (self *Object) Call(args *Object, kw *Object) *Object {
	result := C.PyObject_Call(self.cptr, args.cptr, kw.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) CallObject(args *Object) *Object {
	result := C.PyObject_CallObject(self.cptr, args.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

// TODO:
// PyAPI_FUNC(PyObject *) PyObject_CallFunction(PyObject *callable_object,
//                                                 char *format, ...);
//      PyAPI_FUNC(PyObject *) PyObject_CallMethod(PyObject *o, char *m,
//                                               char *format, ...);
//      PyAPI_FUNC(PyObject *) _PyObject_CallFunction_SizeT(PyObject *callable,
// 							 char *format, ...);
//      PyAPI_FUNC(PyObject *) _PyObject_CallMethod_SizeT(PyObject *o,
// 						       char *name,
// 						       char *format, ...);

//      PyAPI_FUNC(PyObject *) PyObject_CallFunctionObjArgs(PyObject *callable,
//                                                         ...);
//      PyAPI_FUNC(PyObject *) PyObject_CallMethodObjArgs(PyObject *o,
//                                                       PyObject *m, ...);


func (self *Object) Hash() int {
	result := C.PyObject_Hash(self.cptr);
	return int(result);
}

func (self *Object) IsTrue() int {
	result := C.PyObject_IsTrue(self.cptr);
	return int(result);
}

func (self *Object) Not() int {
	result := C.PyObject_Not(self.cptr);
	return int(result);
}

func (self *Object) Type() *Object {
	result := C.PyObject_Type(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Size() int {
	result := C.PyObject_Size(self.cptr);
	return int(result);
}

func (self *Object) Length() int {
	return self.Size();
}

func (self *Object) GetItem(key *Object) *Object {
	result := C.PyObject_GetItem(self.cptr, key.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) SetItem(key *Object, v *Object) int {
	result := C.PyObject_SetItem(self.cptr, key.cptr, v.cptr);
	return int(result);
}

func (self *Object) DelItemString(key string) int {
	result := C.PyObject_DelItemString(self.cptr, C.CString(key));
	return int(result);
}

func (self *Object) DelItem(key *Object) int {
	result := C.PyObject_DelItem(self.cptr, key.cptr);
	return int(result);
}
// TODO: buffer APIs

// iterator
func (self *Object) GetIter() *Object {
	result := C.PyObject_GetIter(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Iter_Check() int {
	return int(C.IterCheck(self.cptr));
}

func (self *Object) Iter_Next() *Object {
	result := C.PyIter_Next(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

// Number Protocol

func (self *Object) Number_Check() int {
	result := C.PyNumber_Check(self.cptr);
	return int(result);
}

func (self *Object) Number_Add(o2 *Object) *Object {
	result := C.PyNumber_Add(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Subtract(o2 *Object) *Object {
	result := C.PyNumber_Subtract(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Multiply(o2 *Object) *Object 
{
	result := C.PyNumber_Multiply(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Divide(o2 *Object) *Object {
	result := C.PyNumber_Divide(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_FloorDivide(o2 *Object) *Object {
	result := C.PyNumber_FloorDivide(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_TrueDivide(o2 *Object) *Object {
	result := C.PyNumber_TrueDivide(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Remainder(o2 *Object) *Object {
	result := C.PyNumber_Remainder(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Divmod(o2 *Object) *Object {
	result := C.PyNumber_Divmod(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Power(o2 *Object, o3 *Object) *Object {
	result := C.PyNumber_Power(self.cptr, o2.cptr, o3.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Negative() *Object {
	result := C.PyNumber_Negative(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Positive() *Object {
	result := C.PyNumber_Positive(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Absolute() *Object {
	result := C.PyNumber_Absolute(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Invert() *Object {
	result := C.PyNumber_Invert(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Lshift(o2 *Object) *Object {
	result := C.PyNumber_Lshift(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Rshift(o2 *Object) *Object {
	result := C.PyNumber_Rshift(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_And(o2 *Object) *Object {
	result := C.PyNumber_And(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Xor(o2 *Object) *Object {
	result := C.PyNumber_Xor(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Or(o2 *Object) *Object {
	result := C.PyNumber_Or(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Index_Check() int {
	return int(C.IndexCheck(self.cptr));
}

func (self *Object) Number_Index() *Object {
	result := C.PyNumber_Index(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_AsSsize(exc *Object) int {
	return int(C.PyNumber_AsSsize_t(self.cptr, exc.cptr));
}

func (self *Object) Number_Int() *Object {
	result := C.PyNumber_Int(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Long() *Object {
	result := C.PyNumber_Long(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_Float() *Object {
	result := C.PyNumber_Float(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceAdd(o2 *Object) *Object {
	result := C.PyNumber_InPlaceAdd(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceSubtract(o2 *Object) *Object {
	result := C.PyNumber_InPlaceSubtract(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceMultiply(o2 *Object) *Object {
	result := C.PyNumber_InPlaceMultiply(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceDivide(o2 *Object) *Object {
	result := C.PyNumber_InPlaceDivide(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceFloorDivide(o2 *Object) *Object {
	result := C.PyNumber_InPlaceFloorDivide(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceTrueDivide(o2 *Object) *Object {
	result := C.PyNumber_InPlaceTrueDivide(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceRemainder(o2 *Object) *Object {
	result := C.PyNumber_InPlaceRemainder(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlacePower(o2 *Object, o3 *Object) *Object {
	result := C.PyNumber_InPlacePower(self.cptr, o2.cptr, o3.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceLshift(o2 *Object) *Object {
	result := C.PyNumber_InPlaceLshift(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceRshift(o2 *Object) *Object {
	result := C.PyNumber_InPlaceRshift(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceAnd(o2 *Object) *Object {
	result := C.PyNumber_InPlaceAnd(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceXor(o2 *Object) *Object {
	result := C.PyNumber_InPlaceXor(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_InPlaceOr(o2 *Object) *Object {
	result := C.PyNumber_InPlaceOr(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Number_ToBase(base int) *Object {
	result := C.PyNumber_ToBase(self.cptr, C.int(base));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

// Sequence protocol

func (self *Object) Sequence_Check() int {
	return int(C.PySequence_Check(self.cptr));
}

func (self *Object) Sequence_Size() int {
	return int(C.PySequence_Size(self.cptr));
}

func (self *Object) Sequence_Concat(o2 *Object) *Object {
	result := C.PySequence_Concat(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Sequence_Repeat(count int) *Object {
	result := C.PySequence_Repeat(self.cptr, C.Py_ssize_t(count));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Sequence_GetItem(i int) *Object {
	result := C.PySequence_GetItem(self.cptr, C.Py_ssize_t(i));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Sequence_GetSlice(i1 int, i2 int) *Object {
	result := C.PySequence_GetSlice(self.cptr, C.Py_ssize_t(i1), C.Py_ssize_t(i2));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Sequence_SetItem(i int, v *Object) int {
	return int(C.PySequence_SetItem(self.cptr, C.Py_ssize_t(i), v.cptr));
}

func (self *Object) Sequence_DelItem(i int) int {
	return int(C.PySequence_DelItem(self.cptr, C.Py_ssize_t(i)));
}

func (self *Object) Sequence_SetSlice(i1 int, i2 int, v *Object) int {
	return int(C.PySequence_SetSlice(self.cptr, C.Py_ssize_t(i1), C.Py_ssize_t(i2), v.cptr));
}

func (self *Object) Sequence_DelSlice(i1 int, i2 int) int {
	return int(C.PySequence_DelSlice(self.cptr, C.Py_ssize_t(i1), C.Py_ssize_t(i2)));
}

func (self *Object) Sequence_Tuple() *Object {
	result := C.PySequence_Tuple(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Sequence_List() *Object {
	result := C.PySequence_List(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Sequence_Fast(m string) *Object {
	result := C.PySequence_Fast(self.cptr, C.CString(m));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Sequence_Count(value *Object) int {
	return int(C.PySequence_Count(self.cptr, value.cptr));
}

func (self *Object) Sequence_Contains(ob *Object) int {
	return int(C.PySequence_Contains(self.cptr, ob.cptr));
}

func (self *Object) Sequence_In(value *Object) int {
	return int(C.PySequence_In(self.cptr, value.cptr));
}

func (self *Object) Sequence_Index(value *Object) int {
	return int(C.PySequence_Index(self.cptr, value.cptr));
}

func (self *Object) Sequence_InPlaceConcat(o2 *Object) *Object {
	result := C.PySequence_InPlaceConcat(self.cptr, o2.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

func (self *Object) Sequence_InPlaceRepeat(count int) *Object {
	result := C.PySequence_InPlaceRepeat(self.cptr, C.Py_ssize_t(count));
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

// Mapping protocol
func (self *Object) Mapping_Check() int {
	return int(C.PyMapping_Check(self.cptr));
}

func (self *Object) Mapping_Size() int {
	return int(C.PyMapping_Size(self.cptr));
}

func (self *Object) Mapping_DelItemString(key string) int {
	result := C.PyObject_DelItemString(self.cptr, C.CString(key));
	return int(result);
}

func (self *Object) Mapping_DelItem(key *Object) int {
	result := C.PyObject_DelItem(self.cptr, key.cptr);
	return int(result);
}

func (self *Object) Mapping_HasKeyString(key string) int {
	return int(C.PyMapping_HasKeyString(self.cptr, C.CString(key)));
}

func (self *Object) Mapping_HasKey(key *Object) int {
	return int(C.PyMapping_HasKey(self.cptr, key.cptr));
}

func (self *Object) Mapping_Keys() *Object {
	result := C.Mapping_Keys(self.cptr);
	return newObject((*C.PyObject)(unsafe.Pointer(result)));
}

// TODO: More Generic Object Interface



func (self *Object) DecRef() {
	C.Py_DecRef(self.cptr);
}

func XDecRef(self *Object) {
	C.Py_XDecRef(self.cptr);
}

// tupleobject.h

func Tuple_New(size int) *Object {
	self := new(Object);
	result := C.PyTuple_New(C.Py_ssize_t(size));
	self.cptr = result;
	return self;
}

func (self *Object) Tuple_SetItem(index int, item *Object) int {
	result := C.PyTuple_SetItem(self.cptr, C.Py_ssize_t(index), item.cptr);
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

