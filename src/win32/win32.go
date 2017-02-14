/*Package win32dllProc Copyright (c) 2017  by Dusan B. Jovanovic dbj@dbj.org

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package win32

import (
	"sync"
	"syscall"
	"unsafe"
)

func main() {}

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300
)

//DllProc is defined by its name and the dll in which it resides
type DllProc interface {
      Proc () * syscall.LazyProc
}

/* dllProc
WIN32 DLL and DLL proc_ encapsulation
NOTE: types names with small first letter are hidden ie not visible outside of package
*/
type dllProc struct {
	dllName  string
	procName string
	dll_   *syscall.LazyDLL
	proc_     *syscall.LazyProc
	mux      sync.Mutex
}

/*
return the instance of the proc from a given dll dynamicaly loaded
this is DLLPorc interface implementation
*/
func (dllInstance dllProc) Proc() *syscall.LazyProc {

	if dllInstance.dllName == "" {
		panic("dllProc has not been given dllName!")
	}
	
	if dllInstance.procName == "" {
		panic("dllProc has not been given procName!")
	}

	if dllInstance.dll_ == nil {
		dllInstance.dll_ = syscall.NewLazyDLL(dllInstance.dllName)
	}

	if dllInstance.proc_ == nil {
		dllInstance.proc_ = dllInstance.dll_.NewProc(dllInstance.procName)
	}
	return dllInstance.proc_
}


//FactoryMethod returns uninitialized isntance of dllProc type
func FactoryMethod(dllNameArg string, funameArg string) DllProc {

	if dllNameArg == "" {
		panic("DLL name must be given")
	}
	if funameArg == "" {
		panic("Function name must be given")
	}
	return dllProc{dllName: dllNameArg, procName: funameArg}
}

//MessageBox shows the message using message box of the underlying OS
//in this case only WIN32
type MessageBox interface {
	Show(message string, title string, decoration uint) uintptr
}

// MessageBox implementor
type mBoxImp struct {
	theProc dllProc
}

//implementation of the Show method of the MessageBoxW interface
func (box mBoxImp) Show(message string, title string, decoration uint) uintptr {
	
	box.theProc.mux.Lock()
	defer box.theProc.mux.Unlock()

	ret, _, _ := box.theProc.Proc().Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(message))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(decoration))

	return ret
}

//FactoriseMessageBox creates WIN32 MessageBoxW
func FactoriseMessageBox() MessageBox {
	p, ok := FactoryMethod("user32", "MessageBoxW").(dllProc)

	if false == ok  {
		panic("Zero value dllProc returned from FactoryMethod('user32', 'MessageBoxW') ")
	}
	return mBoxImp{ theProc: p  }
}
