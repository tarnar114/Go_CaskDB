package main
import(
   "os"
    "sync"
   _"time"
)
//KV pair struct used to represent propertiees
// of the KV's in the data files
type KVPair struct {
    timestamp int 
    Key_Size int 
    Value_Size int 
    Key string 
    Val string
};
//struct to represent metadata in hashmap 
type Hash struct{
    timestamp uint64 
    value_pos uint64 
    value_sz uint32 
    file_id uint32
}
//struct to represent hashmap and mutext for concurrent reads
type HashMap struct{
    data map[int]Hash
    mu sync.RWMutex
}
//struct to represent Data file used to store KV's
type File struct {
    file *os.File
    offset uint64
    mu sync.Mutex
}


