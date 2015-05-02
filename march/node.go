
package march

type Node struct {
    masterRef *Master
    storeRef *MarchStore
}

func NewNode(mode string) *Node {
    node := new(Node)
    if mode == "master" {
        node.masterRef = NewMaster()
    } else {
        node.storeRef = NewMarchStore()
    }
    return node
}

func (node *Node) GetKey(key string) string {
    if node.masterRef != nil {
        return node.masterRef.GetKey(key)
    } else {
        return node.storeRef.GetKey(key)
    }
}

func (node *Node) PutKey(key string, value string) bool {
    if node.masterRef != nil {
        return node.masterRef.PutKey(key, value)
    } else {
        return node.storeRef.PutKey(key, value)
    }
}

func (node *Node) DeleteKey(key string) bool {
    if node.masterRef != nil {
        return node.masterRef.DeleteKey(key)
    } else {
        return node.storeRef.DeleteKey(key)
    }
}
