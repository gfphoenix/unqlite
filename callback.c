#include <_cgo_export.h>
#include "callback.h"

int cArrayWalk(unqlite_value *a, unqlite_value *b, void *fn)
{
    return goArrayWalk(a,b,fn);
}
int cKvFetchCallback(const void *a, unsigned int b, void *fn)
{
    void *aa = (void *)a;
    return goKvFetchCallback(aa, b, fn);
}
int cKvCursorCallback(const void *key, unsigned int len, void *fn)
{
    void *kk = (void *)key;
    return goKvCursorCallback(kk, len, fn);
}


int unqlite_array_walk_CB(unqlite_value *arr, void *fn)
{
    return unqlite_array_walk(arr, cArrayWalk, fn);
}
int unqlite_kv_fetch_callback_CB(unqlite *pDb, const void *key, int nKeyLen, void *fn)
{
    return unqlite_kv_fetch_callback(pDb, key, nKeyLen, cKvFetchCallback, fn);
}
int unqlite_kv_cursor_key_callback_CB(unqlite_kv_cursor *c, void *fn)
{
    return unqlite_kv_cursor_key_callback(c, cKvCursorCallback, fn);
}
int unqlite_kv_cursor_data_callback_CB(unqlite_kv_cursor *c, void *fn)
{
    return unqlite_kv_cursor_data_callback(c, cKvCursorCallback, fn);
}
