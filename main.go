/*
Basically an 10 min extension on top of https://github.com/golang/go/wiki/WindowsDLLs

Copyright (c) 2017  by Dusan B. Jovanovic dbj@dbj.org

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package main

// "unsafe"
import (
	"fmt"
	"win32"
)

func main() {
	beep := win32.FactoryMethod("kernel32.dll","Beep")
	
	frequency := uintptr(1000)
	duration := uintptr(3000)

	beep.Call(frequency, duration) 

	var messageBox = win32.FactoriseMessageBox()
	var ret = messageBox.Show("This test is Done.", "DBJ*GOWIN", win32.MB_OK|win32.MB_ICONINFORMATION)

	fmt.Printf("Message Box Returned: %d\n", ret)

}
