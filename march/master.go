package march

import (
    "crypto/md5"
    "encoding/binary"
    "sort"
)

const Size = 16
const BlockSize = 16

type StoreRef struct {
    Id uint16
    IPAddr string
}


type Master struct {
    Stores map[uint16]*StoreRef
    Keys []uint16
}

func NewMaster() *Master {
    ms := new(Master)
    ms.Stores = make(map[uint16]*StoreRef)
    return ms
}

func HashKey(key string) uint16 {
    key_hash := md5.New().Sum([]byte(key))
    return binary.LittleEndian.Uint16(key_hash)
}

func (s *Master) FindSuccessiveId(id uint16) uint16 {
    if len(s.Keys) == 0 {
        return 0
    }

    var ident uint16
    for _, ident = range s.Keys {
        if id < ident {
            return ident
        }
    }
    return s.Keys[0]
}

func (s *Master) ResolveKeyToNode(key string) *StoreRef {
    key_hash := HashKey(key)
    node, ok := s.Stores[key_hash]
    if ok {
        return node
    } else {
        id := s.FindSuccessiveId(key_hash)
        node, _ := s.Stores[id]
        return node
    }
}

type Uint16Slice []uint16
func (s Uint16Slice) Len() int           { return len(s) }
func (s Uint16Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s Uint16Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *Master) StoresCount() int {
    return len(s.Keys)
}

func (s *Master) AddStore(ip string) {
    key_hash := HashKey(ip)
    store_ref := new(StoreRef)
    store_ref.IPAddr = ip
    s.Stores[key_hash] = store_ref
    s.Keys = append(s.Keys, key_hash)
    sort.Sort(Uint16Slice(s.Keys))
}
