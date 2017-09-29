package server

import (
    "encoding/hex"
    "fmt"
    "google.golang.org/appengine"
    "google.golang.org/appengine/log"
    "google.golang.org/appengine/memcache"
    "net/http"
    "time"
    "context"
    "google.golang.org/appengine/datastore"
    "encoding/json"
    "strconv"
)

var local *time.Location

func init() {
    http.HandleFunc("/temp", handler)
    http.HandleFunc("/push", handlerPush)
    http.HandleFunc("/data", handlerData)
    local, _ = time.LoadLocation("Europe/Prague")
}

// [GET] https://kachnicka-170919.appspot.com/push?x={data}&Time={Time}&Snr={Snr}&Station={Station}&Lat={Lat}&Lng={Lng}&Device={Device}

type Metadata struct {
    Time     int32
    Snr      string
    Station  string
    Device   string
    Position appengine.GeoPoint
}

func handlerData(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    q := datastore.NewQuery("Metadata")

    var meta []Metadata
    keys, err := q.GetAll(ctx, &meta);
    if err != nil {
        fmt.Fprintf(w, "Datastore error: %v", err)
    }

    if len(keys) > 0 {
        datastore.Get(ctx, keys[0], &meta)
        for i, d := range meta {
            b, _ := json.Marshal(d)
            fmt.Fprintf(w, "\n%d -> %s", i, b)
        }
    }
}

func handlerPush(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    ctx := appengine.NewContext(r)

    tempS := r.URL.Query().Get("x")
    if tempS == "" {
        fmt.Fprint(w, "Missing data")
        w.WriteHeader(400)
        return
    }

    if r.URL.Query().Get("device") == "" {
        fmt.Fprint(w, "Device name required")
        w.WriteHeader(400)
        return
    }

    decoded, err := hex.DecodeString(tempS)

    if err != nil {
        fmt.Fprintf(w, "Unable to convert: %s", tempS)
        w.WriteHeader(400)
        return
    }
    temp := fmt.Sprintf("%s", decoded)
    setKeyToCache(ctx, "TEMP", temp, r)

    tempTime := time.Now().In(local).String()
    setKeyToCache(ctx, "TEMP-TIME", tempTime, r)

    timeMs, err := strconv.ParseInt(r.URL.Query().Get("time"), 10, 32);
    if err != nil {
        log.Errorf(ctx, "Error converting time %s", r.URL.Query().Get("time"))
    }

    lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 32);
    if err != nil {
        log.Errorf(ctx, "Error converting lat %s", r.URL.Query().Get("lat"))
    }
    lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 32);
    if err != nil {
        log.Errorf(ctx, "Error converting lng %s", r.URL.Query().Get("lng"))
    }

    gp := appengine.GeoPoint{Lat: lat, Lng: lng}
    meta := Metadata{
        Time:     int32(timeMs),
        Snr:      r.URL.Query().Get("snr"),
        Station:  r.URL.Query().Get("station"),
        Device:   r.URL.Query().Get("device"),
        Position: gp,
    }
    log.Infof(ctx, "Struct %v", meta)
    key := datastore.NewIncompleteKey(ctx, "Metadata", nil)
    if _, err := datastore.Put(ctx, key, &meta); err != nil {
        log.Errorf(ctx, "Unable to store data %v", err)
    }
    dur := time.Since(start)
    fmt.Fprintf(w, "Done, Time %f ms", float32(dur.Seconds()*1000))
}

func handler(w http.ResponseWriter, r *http.Request) {

    temp, err1 := getKeyFromCache("TEMP", r)
    time, err2 := getKeyFromCache("TEMP-TIME", r)

    if err1 != nil || err2 != nil {
        fmt.Fprintf(w, "Unable to get from cache: %v %v", err1, err2)
        w.WriteHeader(500)
        return
    }

    //d := '\u2103'
    fmt.Fprintf(w, "Teplota bazen -> %s čas: %s", temp, time)
}

func setKeyToCache(ctx context.Context, key string, value string, r *http.Request) {

    item := &memcache.Item{
        Key:   key,
        Value: []byte(value),
    }
    if err := memcache.Set(ctx, item); err == memcache.ErrNotStored {
        log.Infof(ctx, "item with key %q already exists", item.Key)
    } else if err != nil {
        log.Errorf(ctx, "error adding item: %v", err)
    }
}

func getKeyFromCache(key string, r *http.Request) (string, error) {

    ctx := appengine.NewContext(r)
    if item, err := memcache.Get(ctx, key); err == memcache.ErrCacheMiss {
        log.Infof(ctx, "item not in the cache")
        return "---", err
    } else if err != nil {
        log.Infof(ctx, "Error %v", err)
        return "---", err
    } else {
        return string(item.Value), nil
    }
}
