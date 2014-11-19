package unqlite
//#cgo linux CFLAGS: -I. -I.. -I../..
//#cgo CFLAGS: -I. -I.. -I../..
//#cgo linux LDFLAGS: -lunqlite
// void free(void *);
//#include <stdlib.h>
//#include "unqlite.h"
//#include "callback.h"
import "C"
import (
    "fmt"
    "unsafe"
)

type Jx9VmConfig int
const (
    Vm_OUTPUT    = Jx9VmConfig(C.UNQLITE_VM_CONFIG_OUTPUT)
    Vm_IMPORT_PATH =Jx9VmConfig(C.UNQLITE_VM_CONFIG_IMPORT_PATH)
    Vm_ERR_REPORT =Jx9VmConfig(C.UNQLITE_VM_CONFIG_ERR_REPORT)
    Vm_RECURSION_DEPTH=Jx9VmConfig(C.UNQLITE_VM_CONFIG_RECURSION_DEPTH)
    Vm_OUTPUT_LENGTH =Jx9VmConfig(C.UNQLITE_VM_OUTPUT_LENGTH)
    Vm_CREATE_VAR =Jx9VmConfig(C.UNQLITE_VM_CONFIG_CREATE_VAR)
    Vm_HTTP_REQUEST =Jx9VmConfig(C.UNQLITE_VM_CONFIG_HTTP_REQUEST)
    Vm_SERVER_ATTR=Jx9VmConfig(C.UNQLITE_VM_CONFIG_SERVER_ATTR)
    Vm_ENV_ATTR=Jx9VmConfig(C.UNQLITE_VM_CONFIG_ENV_ATTR)
    Vm_EXEC_VALUE=Jx9VmConfig(C.UNQLITE_VM_CONFIG_EXEC_VALUE)
    Vm_IO_STREAM=Jx9VmConfig(C.UNQLITE_VM_CONFIG_IO_STREAM)
    Vm_ARGV_ENTRY=Jx9VmConfig(C.UNQLITE_VM_CONFIG_ARGV_ENTRY)
    Vm_EXTRACT_OUTPUT=Jx9VmConfig(C.UNQLITE_VM_CONFIG_EXTRACT_OUTPUT)
)
type LibConfig int
const (
    Lib_USER_MALLOC = (C.UNQLITE_LIB_CONFIG_USER_MALLOC)
)
type SyncFlags uint
const (
    Sync_NORMAL     = SyncFlags(C.UNQLITE_SYNC_NORMAL)
    Sync_FULL       = SyncFlags(C.UNQLITE_SYNC_FULL)
    Sync_DATAONLY   = SyncFlags(C.UNQLITE_SYNC_DATAONLY)
)
type OpenFlags uint
const (
    Open_READONLY   OpenFlags = C.UNQLITE_OPEN_READONLY
    Open_CREATE OpenFlags = C.UNQLITE_OPEN_CREATE
    Open_RDWR   OpenFlags = C.UNQLITE_OPEN_READWRITE
    Open_MMAP   OpenFlags = C.UNQLITE_OPEN_MMAP
//    Open_TEMP   OpenFlags = C.UNQLITE_OPEN_TEMP
    Open_MEM    OpenFlags = C.UNQLITE_OPEN_IN_MEMORY
    Open_NO_JOURNALING OpenFlags = C.UNQLITE_OPEN_OMIT_JOURNALING
    Open_NO_MUTEX OpenFlags = C.UNQLITE_OPEN_NOMUTEX
)

// file locking levels
type LockLevel int
const (
    Lock_NONE   = LockLevel(C.UNQLITE_LOCK_NONE)
)
type Unqlite struct {
    db *C.unqlite
}
// Data base Engine handle
func Open(file string, mode OpenFlags) (*Unqlite,error) {
    f := C.CString(file)
    defer C.free(unsafe.Pointer(f))
    var db Unqlite
    e := C.unqlite_open(&db.db, f, C.uint(mode))
    if e == C.UNQLITE_OK {
        return &db, nil
    }
    return nil, code2Error(e)
}
func (u *Unqlite)Config() {
}
func (u *Unqlite)Close() error{
    e := C.unqlite_close(u.db)
    return code2Error(e)
}

// K-V store
func (u *Unqlite)KvStore(key, value []byte) error {
    e := C.unqlite_kv_store(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&value[0]), C.unqlite_int64(len(value)))
    return code2Error(e)
}
func (u *Unqlite)KvStoreFmt(key []byte, format string, a...interface{}) error {
    val := fmt.Sprintf(format, a...)
    return u.KvStore(key,[]byte(val))
}
func (u *Unqlite)KvAppend(key, value []byte) error{
    e := C.unqlite_kv_append(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&value[0]), C.unqlite_int64(len(value)))
    return code2Error(e)
}
func (u *Unqlite)KvAppendFmt(key []byte, format string, a... interface{})error{
    v := fmt.Sprintf(format, a...)
    return u.KvAppend(key, []byte(v))
}
func (u *Unqlite)KvFetch(key, value []byte) ([]byte, error){
    outlen := C.unqlite_int64(len(value))
    e := C.unqlite_kv_fetch(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&value[0]), &outlen)
    err := code2Error(e)
    n := int(outlen)
    if n <=0 {
        return nil, err
    }
    return value[:n], err
}
func (u *Unqlite)KvDelete(key []byte)error{
    e := C.unqlite_kv_delete(u.db, unsafe.Pointer(&key[0]), C.int(len(key)))
    return code2Error(e)
}
func (u *Unqlite)Kv_config_hash(hash func([]byte) uint) {
}

//Cursor Iterator Interfaces
type Cursor struct {
    c   *C.unqlite_kv_cursor
}
func (u *Unqlite)NewCursor() (*Cursor,error) {
    var cur Cursor
    e := C.unqlite_kv_cursor_init(u.db, &(cur.c))
    return &cur, code2Error(e)
}
func (u *Unqlite)Release(c *Cursor)error{
    e := C.unqlite_kv_cursor_release(u.db, c.c)
    c.c = nil
    return code2Error(e)
}
func (c *Cursor)Reset()error{
    e := C.unqlite_kv_cursor_reset(c.c)
    return code2Error(e)
}

// positioning database cursors

func (c *Cursor)Seek(key []byte, iPos int)error{
    e := C.unqlite_kv_cursor_seek(c.c, unsafe.Pointer(&key[0]), C.int(len(key)), C.int(iPos))
    return code2Error(e)
}
func (c *Cursor)First()error{
    e := C.unqlite_kv_cursor_first_entry(c.c)
    return code2Error(e)
}
func (c *Cursor)Last()error{
    e := C.unqlite_kv_cursor_last_entry(c.c)
    return code2Error(e)
}
func (c *Cursor)Valid()error{
    e := C.unqlite_kv_cursor_valid_entry(c.c)
    return code2Error(e)
}
func (c *Cursor)Next() error{
    e := C.unqlite_kv_cursor_next_entry(c.c)
    return code2Error(e)
}
func (c *Cursor)Prev()error{
    e := C.unqlite_kv_cursor_prev_entry(c.c)
    return code2Error(e)
}

// extracting data from database cursors

func (c *Cursor)Key(value []byte) (int,error) {
    outlen := C.int(len(value))
    e := C.unqlite_kv_cursor_key(c.c, unsafe.Pointer(&value[0]), &outlen)
    return int(outlen), code2Error(e)
}
func (c *Cursor)Data(data []byte) (int64,error) {
    outlen := C.unqlite_int64(len(data))
    e := C.unqlite_kv_cursor_data(c.c, unsafe.Pointer(&data[0]), &outlen)
    return int64(outlen), code2Error(e)
}


// deleteing records using database cursors
func (c *Cursor)Delete()error{
    e := C.unqlite_kv_cursor_delete_entry(c.c)
    return code2Error(e)
}

// Jx9 document store
type Jx9 struct {
    j *C.unqlite_vm
}
func (u *Unqlite)Compile(jx9Code string) (*Jx9, error){
    code := C.CString(jx9Code)
    defer C.free(unsafe.Pointer(code))
    var jx9 Jx9
    e := C.unqlite_compile(u.db, code, C.int(len(jx9Code)), &jx9.j)
    return &jx9, code2Error(e)
}
func (u *Unqlite)Compile_file(file string) (*Jx9, error) {
    f := C.CString(file)
    defer C.free(unsafe.Pointer(f))
    var jx9 Jx9
    e := C.unqlite_compile_file(u.db, f, &jx9.j)
    return &jx9, code2Error(e)
}
func (j *Jx9)Release()error{
    e := C.unqlite_vm_release(j.j)
    j.j = nil
    return code2Error(e)
}
func (u Unqlite)Vm_config(){
}
func (j *Jx9)Exec()error{
    e := C.unqlite_vm_exec(j.j)
    return code2Error(e)
}
func (j *Jx9)Reset()error{
    e := C.unqlite_vm_reset(j.j)
    return code2Error(e)
}
func (j *Jx9)Dump(xConsumer func(dump []byte) int ){
}

// manual transaction
func (u *Unqlite)Begin()error{
    e := C.unqlite_begin(u.db)
    return code2Error(e)
}
func (u *Unqlite)Commit()error{
    e := C.unqlite_commit(u.db)
    return code2Error(e)
}
func (u *Unqlite)Rollback()error{
    e := C.unqlite_rollback(u.db)
    return code2Error(e)
}
// utility interfaces
// unqlite_load_mmaped_file
// unqlite_util_release_mmaped_file
// unqlite_util_random_string
// unqlite_util_random_num

// create_function
// delete_function
// create_constant
// delete_constant

type Value struct {
    v *C.unqlite_value
}
func (j *Jx9)NewScalar()*Value{
    var val Value
    val.v = C.unqlite_vm_new_scalar(j.j)
    if val.v == nil {
        return nil
    }
    return &val
}
func (j *Jx9)NewArray() *Array{
    var a Array
    a.v = C.unqlite_vm_new_array(j.j)
    if a.v == nil {
        return nil
    }
    return &a
}
func (j *Jx9)ReleaseValue(v *Value)error{
    e := C.unqlite_vm_release_value(j.j, v.v)
    v.v = nil
    return code2Error(e)
}
func (j *Jx9)CreateFunc(){
}
func (j *Jx9)DeleteFunc(name string)error{
    s := C.CString(name)
    defer C.free(unsafe.Pointer(s))
    e := C.unqlite_delete_function(j.j, s)
    return code2Error(e)
}
func (j *Jx9)CreateConst(){
}
func (j *Jx9)DeleteConst(name string)error{
    s := C.CString(name)
    defer C.free(unsafe.Pointer(s))
    e := C.unqlite_delete_constant(j.j, s)
    return code2Error(e)
}
type Context struct {
    c *C.unqlite_context
}
func (c *Context)NewScalar() *Value{
    var v Value
    v.v = C.unqlite_context_new_scalar(c.c)
    if v.v == nil {
        return nil
    }
    return &v
}
func (c *Context)NewArray() *Array{
    var a Array
    a.v = C.unqlite_context_new_array(c.c)
    if a.v == nil {
        return nil
    }
    return &a
}
func (c *Context)ReleaseValue(v *Value){
    C.unqlite_context_release_value(c.c, v.v)
    v.v = nil
}
func (c *Context)PushAuxData(data interface{})error{
    e := C.unqlite_context_push_aux_data(c.c, unsafe.Pointer(&data))
    return code2Error(e)
}
func (c *Context)PeekAuxData() (interface{}, bool){
    p := C.unqlite_context_peek_aux_data(c.c)
    if p != nil {
        return *(*interface{})(unsafe.Pointer(p)), true
    }
    return nil, false
}
    

// init value of Value
func (v *Value)Int(i int)error{
    e := C.unqlite_value_int(v.v, C.int(i))
    return code2Error(e)
}
func (v *Value)Int64(i64 int64)error{
    e := C.unqlite_value_int64(v.v, C.unqlite_int64(i64))
    return code2Error(e)
}
func (v *Value)Bool(b bool)error{
    bb := C.int(0)
    if b {
        bb = C.int(1)
    }
    e := C.unqlite_value_bool(v.v, bb)
    return code2Error(e)
}
func (v *Value)Null()error{
    e := C.unqlite_value_null(v.v)
    return code2Error(e)
}
func (v *Value)Float(d float64)error{
    e := C.unqlite_value_double(v.v, C.double(d))
    return code2Error(e)
}
func (v *Value)String(str string)error{
    s := C.CString(str)
    defer C.free(unsafe.Pointer(s))
    e := C.unqlite_value_string(v.v, s, C.int(len(str)))
    return code2Error(e)
}
func (v *Value)StringFormat(format string, a... interface{})error {
    str := fmt.Sprintf(format, a...)
    s   := C.CString(str)
    defer C.free(unsafe.Pointer(s))
    e := C.unqlite_value_string(v.v, s, C.int(len(str)))
    return code2Error(e)
}
func (v *Value)ResetStringCursor()error{
    e := C.unqlite_value_reset_string_cursor(v.v)
    return code2Error(e)
}
func (v *Value)Resource(p interface{})error{
    e := C.unqlite_value_resource(v.v, unsafe.Pointer(&p))
    return code2Error(e)
}
func (v *Value)Release()error{
    e := C.unqlite_value_release(v.v)
    return code2Error(e)
}
func (v *Value)ToInt() int {
    i := C.unqlite_value_to_int(v.v)
    return int(i)
}
func (v *Value)ToBool() bool {
    i := C.unqlite_value_to_bool(v.v)
    return int(i) != 0
}
func (v *Value)ToInt64() int64{
    i := C.unqlite_value_to_int64(v.v)
    return int64(i)
}
func (v *Value)ToDouble() float64{
    f := C.unqlite_value_to_double(v.v)
    return float64(f)
}
func (v *Value)ToString() string {
    s := C.unqlite_value_to_string(v.v,nil)
    return C.GoString(s)
}
func (v *Value)ToResource() interface{}{
    p := C.unqlite_value_to_resource(v.v)
    return *(*interface{})(unsafe.Pointer(p))
}
func (v *Value)Cmp(v2 *Value, strict bool) int {
    bb := C.int(0)
    if strict {
        bb = C.int(1)
    }
    i := C.unqlite_value_compare(v.v, v2.v, bb)
    return int(i)
}
func (v *Value)IsInt() bool {
    i := C.unqlite_value_is_int(v.v)
    return int(i) != 0
}
func (v *Value)IsFloat() bool {
    i := C.unqlite_value_is_float(v.v)
    return int(i) != 0
}
func (v *Value)IsBool() bool {
    i := C.unqlite_value_is_bool(v.v)
    return int(i)!=0
}
func (v *Value)IsString() bool {
    i := C.unqlite_value_is_string(v.v)
    return int(i)!=0
}
func (v *Value)IsNull() bool {
    i := C.unqlite_value_is_null(v.v)
    return int(i)!=0
}
func (v *Value)IsNumeric() bool {
    i := C.unqlite_value_is_numeric(v.v)
    return int(i)!=0
}
func (v *Value)IsCallable() bool {
    i := C.unqlite_value_is_callable(v.v)
    return int(i)!=0
}
func (v *Value)IsScalar() bool {
    i := C.unqlite_value_is_scalar(v.v)
    return int(i)!=0
}
func (v *Value)IsJsonArray() bool {
    i := C.unqlite_value_is_json_array(v.v)
    return int(i)!=0
}
func (v *Value)IsJsonObject() bool {
    i := C.unqlite_value_is_json_object(v.v)
    return int(i)!=0
}
func (v *Value)IsResource() bool {
    i := C.unqlite_value_is_resource(v.v)
    return int(i)!=0
}
func (v *Value)IsEmpty() bool {
    i := C.unqlite_value_is_empty(v.v)
    return int(i)!=0
}
// Setting the result of a foreign function
func (c *Context)Int(i int)error{
    e := C.unqlite_result_int(c.c, C.int(i))
    return code2Error(e)
}
func (c *Context)Int64(i64 int64)error{
    e := C.unqlite_result_int64(c.c, C.unqlite_int64(i64))
    return code2Error(e)
}
func (c *Context)Bool(b bool)error{
    bb := C.int(0)
    if b {
        bb = C.int(1)
    }
    e := C.unqlite_result_bool(c.c, bb)
    return code2Error(e)
}
func (c *Context)Float(f float64)error{
    e := C.unqlite_result_double(c.c, C.double(f))
    return code2Error(e)
}
func (c *Context)Null(){
    C.unqlite_result_null(c.c)
}
func (c *Context)String(str string)error{
    s := C.CString(str)
    defer C.free(unsafe.Pointer(s))
    e := C.unqlite_result_string(c.c, s, C.int(len(str)))
    return code2Error(e)
}
func (c *Context)StringFormat(format string, a...interface{})error{
    str := fmt.Sprintf(format, a...)
    return c.String(str)
}
func (c *Context)Value(v *Value)error{
    e := C.unqlite_result_value(c.c, v.v)
    return code2Error(e)
}
func (c *Context)Resource(userData interface{})error{
    e := C.unqlite_result_resource(c.c, unsafe.Pointer(&userData))
    return code2Error(e)
}
func (c *Context)UserData() interface{}{
    p := C.unqlite_context_user_data(c.c)
    return *(*interface{})(unsafe.Pointer(p))
}
func (c *Context)FuncName() string {
    s := C.unqlite_function_name(c.c)
    return C.GoString(s)
}

// Json Array/Object Management Interfaces
type Array  Value
func (a *Array)ArrayFetch(key string) *Value{
    var value Value
    s := C.CString(key)
    defer C.free(unsafe.Pointer(s))
    value.v = C.unqlite_array_fetch(a.v, s, C.int(len(key)))
    if value.v == nil {
        return nil
    }
    return &value
}
//export goArrayWalk
func goArrayWalk(a, b *C.unqlite_value, p unsafe.Pointer) int{
    aa  := Value{v:a}
    bb  := Value{v:b}
    fn  := *(*func(l,r Value)int)(unsafe.Pointer(&p))
    return fn(aa, bb)
}
func (arr *Array)Walk(fn func(a, b Value)int)error{
    e := C.unqlite_array_walk_CB(arr.v, *(*unsafe.Pointer)(unsafe.Pointer(&fn)))
    return code2Error(e)
}
func (a *Array)Add(key, value *Value)error{
    e := C.unqlite_array_add_elem(a.v, key.v, value.v)
    return code2Error(e)
}
func (a *Array)AddStringKey(key string, value *Value)error{
    k := C.CString(key)
    defer C.free(unsafe.Pointer(k))
    e := C.unqlite_array_add_strkey_elem(a.v, k, value.v)
    return code2Error(e)
}
func (a *Array)Count() uint {
    return uint(C.unqlite_array_count(a.v))
}
// Call Context Handling Interfaces
//func (c *Context)Output(
// output_format
// throw_error
// throw_error_format
// random_num
// random_string
// user_data
// push_aux_data
// peek_aux_data
// result_buf_length
// function_name
// 
// Call Context Memory Management Interfaces
// alloc_chunk
// realloc_chunk
// free_chunk

func Config(){
}
func Init()error{
    e := C.unqlite_lib_init()
    return code2Error(e)
}
func Shutdown()error{
    e := C.unqlite_lib_shutdown()
    return code2Error(e)
}
func IsThreadSafe() bool {
    i := C.unqlite_lib_is_threadsafe()
    return int(i) != 0
}
func Version() string {
    v := C.unqlite_lib_version()
    return C.GoString(v)
}
func Signature() string {
    s := C.unqlite_lib_signature()
    return C.GoString(s)
}
func Ident() string {
    s := C.unqlite_lib_ident()
    return C.GoString(s)
}
func Copyright() string {
    s := C.unqlite_lib_copyright()
    return C.GoString(s)
}

