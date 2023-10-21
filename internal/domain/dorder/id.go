package dorder

func NewID(value int64) (ID, error) {
	if value <= 0 {
		return IDEmptyValue(), ErrInvalidID
	}

	return ID{value: value}, nil
}

type ID struct {
	value int64
}

func (id ID) Value() int64 {
	return id.value
}

func (id ID) IsEmpty() bool {
	return id.Equals(IDEmptyValue())
}

func (id ID) Equals(other ID) bool {
	return id.value == other.value
}

func IDEmptyValue() ID {
	return ID{}
}
