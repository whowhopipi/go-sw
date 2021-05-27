package main

/* #include <pthread.h>
#include <pthread.h>
 extern void setTLS(int);
 extern int getTLS();
*/
import "C"
import "runtime"

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	C.setTLS(1234)
	println("HHH", C.getTLS())
}

//export Ttt
func Ttt(v int) {
	println("TTTTTTT", v)
}
