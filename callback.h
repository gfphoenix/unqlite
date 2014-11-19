#ifndef CALLBACKS_H
#define CALLBACKS_H

UNQLITE_APIEXPORT int unqlite_kv_fetch_callback(unqlite *pDb,const void *pKey,
	                    int nKeyLen,int (*xConsumer)(const void *,unsigned int,void *),void *pUserData);
int unqlite_array_walk_CB(unqlite_value *arr, void *fn);
int cArrayWalk(unqlite_value *a, unqlite_value *b, void *fn);

int unqlite_kv_fetch_callback_CB(unqlite *pDb, const void *key, int nKeyLen, void *fn);
int cKvFetchCallback(const void *, unsigned int, void *);
//
//
#endif
