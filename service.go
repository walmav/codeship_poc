package microservice

import "strings"

type Node interface {
	GetId(string) (string, error)
}

type NodeImpl struct {
}

func (NodeImpl) GetId(s string) (string, error) {
	return strings.ToLower("blah://blahhost/blahspace/blahprocess?blahattr="), nil

}
