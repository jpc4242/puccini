package js

import (
	"github.com/tliron/puccini/ard"
)

//
// List
//

type List []Coercible

func (self *CloutContext) NewList(list ard.List, entryConstraints Constraints, site interface{}, source interface{}, target interface{}) (List, error) {
	list_ := make([]Coercible, len(list))

	for index, data := range list {
		if entry, err := self.NewCoercible(data, site, source, target); err == nil {
			entry.SetConstraints(entryConstraints)
			list_[index] = entry
		} else {
			return nil, err
		}
	}

	return list_, nil
}

func (self List) Coerce() (interface{}, error) {
	value := make(ard.List, len(self))

	for index, coercible := range self {
		var err error
		if value[index], err = coercible.Coerce(); err != nil {
			return nil, err
		}
	}

	return value, nil
}
