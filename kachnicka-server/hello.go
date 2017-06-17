package hello

import (
	"encoding/hex"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/push", handlerPush)
}

func handlerPush(w http.ResponseWriter, r *http.Request) {
	temp := r.URL.Query().Get("x")
	decoded, err := hex.DecodeString(temp)
  ctx := appengine.NewContext(r)
	if err != nil {
		log.Errorf(ctx, "Unabel to convert: %s", temp)
	} else {
		v := fmt.Sprintf("%s", decoded)
    
		item := &memcache.Item{
			Key:   "TEMP",
			Value: []byte(v),
		}
		if err := memcache.Set(ctx, item); err == memcache.ErrNotStored {
			log.Infof(ctx, "item with key %q already exists", item.Key)
		} else if err != nil {
			log.Errorf(ctx, "error adding item: %v", err)
		}
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	if item, err := memcache.Get(ctx, "TEMP"); err == memcache.ErrCacheMiss {
		log.Infof(ctx, "item not in the cache")
		fmt.Fprint(w, "Empty cache")
	} else if err != nil {
		// log.Errorf(ctx, "error getting item: %v", err)
		fmt.Fprint(w, "Error %v", err)
	} else {
		//          log.Infof(ctx, "the lyric is %q", item.Value)
		s := string(item.Value)
		fmt.Fprint(w, "Kachnicka temp -> ", s)
	}
}

