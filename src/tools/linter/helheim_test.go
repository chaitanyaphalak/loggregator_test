// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package linter_test

import (
	"os"
	"time"
)

type mockFileInfo struct {
	NameCalled chan bool
	NameOutput struct {
		Ret0 chan string
	}
	SizeCalled chan bool
	SizeOutput struct {
		Ret0 chan int64
	}
	ModeCalled chan bool
	ModeOutput struct {
		Ret0 chan os.FileMode
	}
	ModTimeCalled chan bool
	ModTimeOutput struct {
		Ret0 chan time.Time
	}
	IsDirCalled chan bool
	IsDirOutput struct {
		Ret0 chan bool
	}
	SysCalled chan bool
	SysOutput struct {
		Ret0 chan interface{}
	}
}

func newMockFileInfo() *mockFileInfo {
	m := &mockFileInfo{}
	m.NameCalled = make(chan bool, 100)
	m.NameOutput.Ret0 = make(chan string, 100)
	m.SizeCalled = make(chan bool, 100)
	m.SizeOutput.Ret0 = make(chan int64, 100)
	m.ModeCalled = make(chan bool, 100)
	m.ModeOutput.Ret0 = make(chan os.FileMode, 100)
	m.ModTimeCalled = make(chan bool, 100)
	m.ModTimeOutput.Ret0 = make(chan time.Time, 100)
	m.IsDirCalled = make(chan bool, 100)
	m.IsDirOutput.Ret0 = make(chan bool, 100)
	m.SysCalled = make(chan bool, 100)
	m.SysOutput.Ret0 = make(chan interface{}, 100)
	return m
}
func (m *mockFileInfo) Name() string {
	m.NameCalled <- true
	return <-m.NameOutput.Ret0
}
func (m *mockFileInfo) Size() int64 {
	m.SizeCalled <- true
	return <-m.SizeOutput.Ret0
}
func (m *mockFileInfo) Mode() os.FileMode {
	m.ModeCalled <- true
	return <-m.ModeOutput.Ret0
}
func (m *mockFileInfo) ModTime() time.Time {
	m.ModTimeCalled <- true
	return <-m.ModTimeOutput.Ret0
}
func (m *mockFileInfo) IsDir() bool {
	m.IsDirCalled <- true
	return <-m.IsDirOutput.Ret0
}
func (m *mockFileInfo) Sys() interface{} {
	m.SysCalled <- true
	return <-m.SysOutput.Ret0
}
