package unqlite
//#cgo CFLAGS: -I.
//#include <unqlite.h>
//#include <string.h>
//int unqlite_kv_fetch_callback_CB(unqlite *pDb, const void *key, int nKeyLen, void *fn);
//int unqlite_kv_cursor_key_callback_CB(unqlite_kv_cursor *c, void *fn);
//int unqlite_kv_cursor_data_callback_CB(unqlite_kv_cursor *c, void *fn);
import "C"
import (
    "unsafe"
)

// use closure instead of another parameter for func
func (u *Unqlite)KvFetchCallback(key []byte, xConsumer func(data[]byte)) error{
    i := C.unqlite_kv_fetch_callback_CB(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&xConsumer))
    return code2Error(i)
}

//export goKvFetchCallback
func goKvFetchCallback(a unsafe.Pointer, b uint, fn unsafe.Pointer) int {
    if b==0 {
        return 0
    }
    data := make([]byte, b)
    C.memcpy(unsafe.Pointer(&data[0]), a, C.size_t(b))
    Fn := *(*func([]byte)int)(fn)
    return Fn(data)
}
func (c *Cursor)KeyCallback(xConsumer func(value []byte)int) error{
    i := C.unqlite_kv_cursor_key_callback_CB(c.c, unsafe.Pointer(&xConsumer))
    return code2Error(i)
}
func (c *Cursor)DataCallback(xConsumer func(value []byte)int)error {
    i := C.unqlite_kv_cursor_data_callback_CB(c.c, unsafe.Pointer(&xConsumer))
    return code2Error(i)
}
//export goKvCursorCallback
func goKvCursorCallback(data unsafe.Pointer, b uint, fn unsafe.Pointer) int{
    if b==0 {
        return 0
    }
    pdata := make([]byte, b)
    C.memcpy(unsafe.Pointer(&pdata[0]), data, C.size_t(b))
    f := *(*func([]byte)int)(fn)
    return f(pdata)
}
