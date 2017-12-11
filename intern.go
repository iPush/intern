package intern

import "sync"

type intern struct {
	lut  *sync.Map
	lock sync.RWMutex
}

var instance *intern
var once sync.Once

func NewIntern() *intern {
	return &intern{
		lut: &sync.Map{},
	}
}

// GetInstance singleton intern instance
func GetInstance() *intern {
	once.Do(func() {
		instance = NewIntern()
	})
	return instance
}

// Intern a string
func Intern(src string) string {
	str, ok := instance.lut.Load(src)
	if ok {
		return str.(string)
	} else {
		instance.lut.Store(src, src)
		return src
	}
}
