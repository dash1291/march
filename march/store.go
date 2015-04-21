package march

import "sync"

type MarchStore struct {
    Store map[string]string
    sync.Mutex
}

func NewMarchStore() *MarchStore {
    ms := new(MarchStore)
    ms.Store = make(map[string]string)
    return ms
}

func (s *MarchStore) GetKey(key string) string {
    s.Lock()
    val, ok := s.Store[key]
    s.Unlock()

    if !ok {
        return ""
    } else {
        return val
    }
}

func (s *MarchStore) DeleteKey(key string) bool {
    s.Lock()
    _, ok := s.Store[key]

    if ok {
        delete(s.Store, key)
        s.Unlock()
        return true
    } else {
        s.Unlock()
        return false
    }
}

func (s *MarchStore) PutKey(key string, val string) bool {
    s.Lock()
    s.Store[key] = val
    s.Unlock()
    return true
}
