package march

type MarchStore struct {
    Store map[string]string
}

func NewMarchStore() *MarchStore {
    ms := new(MarchStore)
    ms.Store = make(map[string]string)
    return ms
}

func (s *MarchStore) GetKey(key string) string {
    val, ok := s.Store[key]

    if !ok {
        return ""
    } else {
        return val
    }
}

func (s *MarchStore) DeleteKey(key string) bool {
    _, ok := s.Store[key]

    if ok {
        delete(s.Store, key)
        return true
    } else {
        return false
    }
}

func (s *MarchStore) PutKey(key string, val string) bool {
    s.Store[key] = val
    return true
}
