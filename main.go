package main

// #cgo pkg-config: python3-embed ncurses
// #include <Python.h>
import "C"

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unsafe"
)

func main() {

	defer C.Py_Finalize()
	C.Py_Initialize()

	s := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for s.Scan() {
		l := s.Text()

		if strings.HasPrefix(l, ":l ") {
			l = strings.TrimPrefix(l, ":l ")

			cl := C.CString(l)
			mode := C.CString("r")
			file := C.fopen(cl, mode)
			C.free(unsafe.Pointer(mode))

			if file == nil {
				fmt.Println("! file could not be opened")
				fmt.Print("> ")
				continue
			}

			C.PyRun_SimpleFileExFlags(file, cl, 1, nil)
			C.free(unsafe.Pointer(cl))
		} else {

			cl := C.CString(l)
			C.PyRun_SimpleStringFlags(cl, nil)
			C.free(unsafe.Pointer(cl))
		}

		fmt.Print("> ")
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
