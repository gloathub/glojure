package lang

type Volatile struct {
	val interface{}
}

var (
	_ IDeref = (*Volatile)(nil)
)

func NewVolatile(val interface{}) *Volatile {
	return &Volatile{
		val: val,
	}
}

func (v *Volatile) Deref() interface{} {
	return v.val
}

func (v *Volatile) Reset(val interface{}) interface{} {
	v.val = val
	return val
}
