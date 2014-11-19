package unqlite
//#include <unqlite.h>
//
import "C"
import (
    "errors"
)
var (
    err_mem error = errors.New("Err: No mem")
    err_unknown  error = errors.New("Err: Unknown")
    err_busy    error = errors.New("Err: Busy")
    err_io  error = errors.New("Err: I/O")
    err_abort error = errors.New("Err: abort")
    err_read_only error = errors.New("Err: Read only")
    err_not_impl error = errors.New("Err: Not implemented")
    err_perm error = errors.New("Err: Perm")
    err_limit error = errors.New("Err: Limit")
    err_not_found error = errors.New("Err: Not found")
    err_compile error = errors.New("Err: Compile")
    err_vm error = errors.New("Err: Vm")
    err_corrupt error = errors.New("Err: Corrupt")
    err_eof error = errors.New("Err: Eof")
)
func code2Error(code C.int) error {
    if code == C.UNQLITE_OK {
        return nil
    }
    switch code {
    case C.UNQLITE_NOMEM:
        return err_mem
    case C.UNQLITE_ABORT:
        return err_abort
    case C.UNQLITE_IOERR:
        return err_io
    case C.UNQLITE_CORRUPT:
        return err_corrupt
    case C.UNQLITE_BUSY:
        return err_busy
    case C.UNQLITE_PERM:
        return err_perm
    case C.UNQLITE_NOTIMPLEMENTED:
        return err_not_impl
    case C.UNQLITE_NOTFOUND:
        return err_not_found
    case C.UNQLITE_EOF:
        return err_eof
    case C.UNQLITE_LIMIT:
        return err_limit
    case C.UNQLITE_COMPILE_ERR:
        return err_compile
    case C.UNQLITE_VM_ERR:
        return err_vm
    case C.UNQLITE_READ_ONLY:
        return err_read_only
    default:
        return err_unknown
    }
}
