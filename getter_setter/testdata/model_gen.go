package testdata

import "github.com/guregu/null/v5"

func (m *Model1) GetName() string {
	return m.name
}

func (m *Model1) SetName(v string) {
	m.name = v
}

func (m *Model1) GetAge() int {
	return m.age
}

func (m *Model1) SetAge(v int) {
	m.age = v
}

func (m *Model1) GetNullable() null.Bool {
	return m.nullable
}

func (m *Model1) SetNullable(v null.Bool) {
	m.nullable = v
}

func (m *Model1) GetReadonly() []byte {
	return m.readonly
}

func (m *Model1) SetPointer(v *byte) {
	m.pointer = v
}

func (m *Model1) GetM() map[string]int {
	return m.m
}

func (m *Model1) SetC(v chan struct{}) {
	m.c = v
}

func (m *Model1) SetS(v b) {
	m.s = v
}

func (m *Model1) __Model1Iface() {}

type Model1Iface interface {
	Create(a []string)

	GetName() string
	SetName(v string)
	GetAge() int
	SetAge(v int)
	GetNullable() null.Bool
	SetNullable(v null.Bool)
	GetReadonly() []byte
	SetPointer(v *byte)
	GetM() map[string]int
	SetC(v chan struct{})
	SetS(v b)

	__Model1Iface()
}
