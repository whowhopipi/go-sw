// +build sw64

package main

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"sync/atomic"
	"syscall"
	_ "unsafe"
)

//implement in ./asm.s
func asmfn()
func issue53()
func issue55()
func issue5502()
func issue22()

func timediv(v int64, div int32, rem *int32) int32 {
	res := int32(0)
	for bit := 30; bit >= 0; bit-- {
		if v >= int64(div)<<uint(bit) {
			v = v - (int64(div) << uint(bit))
			res += 1 << uint(bit)
		}
	}
	if v >= int64(div) {
		if rem != nil {
			*rem = 0
		}
		return 0x7fffffff
	}
	if rem != nil {
		*rem = int32(v)
	}
	return res
}

var e int

func NeqPtr(a *int) bool {
	return a != &e
}

func bigLocals() byte {
	var tmp [81920]byte
	return tmp[0]
}

//go:noinline
func issue70(x_73 int) {
	if x_73/31 != 2 {
		println("73/31 != 2 got", x_73/31)
	}
}

//go:noinline
func issue77(x_1 int32) int32 {
	return x_1 << 31 >> 31 // Create 0x0000 or 0xFFFF.
}

//go:noinline
func add(a int, b int) int {
	return a + b
}

//go:noinline
func sub(a int, b int) int {
	return a - b
}

//go:noinline
func divide(a, b int) int {
	return a / b
}

//go:noinline
func mod(a, b int) int {
	return a % b
}

//go:noinline
func div(a float64, b float64) float64 {
	return a / b
}

//go:noinline
func mul(a float64, b float64) float64 {
	return a * b
}

//go:noinline
func fmas(a float32, b float32, c float32) float32 {
	return a*b + c
}

//go:noinline
func fmad(a float64, b float64, c float64) float64 {
	return a*b + c
}

//go:noinline
func fmss(a float32, b float32, c float32) float32 {
	return a*b - c
}

//go:noinline
func fmsd(a float64, b float64, c float64) float64 {
	return a*b - c
}

//go:noinline
func checkbool(v bool) {
	if v {
		println(true, "=true")
	} else {
		println(false, "=false")
	}
}

func rEq(a, b int) bool {
	return a == b
}
func fEq(a, b float64) bool {
	return a == b
}

func c()  { c1(); println("pass function call chain") }
func c1() { c2(); println("C1 done") }
func c2() { c3(); println("C2 done") }
func c3() { println("C3 done") }

func f1() {
	println("pass static call")
}

func invokeFn(a func()) {
	a()
}

func checkZero() {
	var buf [100]byte
	for i := 0; i < len(buf); i++ {
		if buf[i] != 0 {
			println("zeroing failed")
			return
		}
	}
	var buf2 [100]int
	for i := 0; i < len(buf2); i++ {
		if buf2[i] != 0 {
			println("zeroing failed")
			return
		}
	}
	var buf3 [100]float64
	for i := 0; i < len(buf3); i++ {
		if buf3[i] != 0 {
			println("zeroing failed")
			return
		}
	}
}

func checkCopy() {
	var buf [100]byte
	var target [100]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}
	copy(target[:], buf[:])
	for i := 0; i < len(buf); i++ {
		if target[i] != buf[i] {
			println("copy failed")
			return
		}
	}
}

func checkMove() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	number1 := numbers[2:10]
	for idx, _ := range number1 {
		if number1[idx] != numbers[idx+2] {
			println("move failed")
			break
		}
	}
}

func round(n, a uint) uint {
	return (n + a - 1) &^ (a - 1)
}

func gcd(a, b uint32) uint32 {
	for b != 0 {
		c := a % b
		a = b
		b = c
	}
	return a
}

func throw(err string) {
	println(err)
	syscall.Exit(-1)
}

func fa() {
	var _p = struct {
		localSize uintptr
	}{33}
	pold := &_p
	p := pold
	if atomic.LoadUintptr(&p.localSize) != _p.localSize {
		throw("LoadUintpr return a wrong value!")
	}
	if p != pold {
		throw("LoadUintpr shouldn't change p!")
	}
}

//go:noinline
func bitclear(a, b uint64) uint64 {
	return a &^ b
}

func f2() {
	if bitclear(12, 4) != 8 {
		throw("12 &^ 4 != 8")
	}
	bs := make([]byte, 33)
	bs[3] = 'b'

	kv := map[string]string{
		"OS":   "linux",
		"ARCH": "sw64",
	}
	if kv["ARCH"] != "sw64" {
		throw("map implement failed")
	}
	if !strings.HasPrefix("abcd123", "abcd") {
		throw("HasPrefix failed..")
	}
	ps := strings.Split("a/b/c", "/")
	if len(ps) != 3 {
		throw("split length wrong")
	}
	if ps[0] != "a" || ps[1] != "b" || ps[2] != "c" {
		throw("split value wrong")
	}
}

func main() {
	reflectmethod4()
	testReocver()
	asmfn()
	issue53()
	issue55()
	issue57()
	issue5502()
	issue22()
	issue70(73)
	println("issue77 == -1 ", issue77(3))
	fa()
	f2()
	c()
	invokeFn(f1)
	invokeFn(func() {
		println("pass lambda call")
	})

	println("gcd(33333, 9) ==", gcd(33333, 9))

	closure := "closure"
	invokeFn(func() {
		println("pass ", closure, " call")
	})

	checkZero()
	checkCopy()
	checkMove()

	if round(8192, 8192) != 8192 {
		println("round failed")
	}

	println("3.0==4.0 -->", fEq(3.0, 4.0))
	println("3.0!=4.0 -->", !fEq(3.0, 4.0))

	println("3==4 -->", rEq(3, 4))
	println("3!=4 -->", !rEq(3, 4))

	println("3+5==8 ==>", add(3, 5) == 8)

	println("13/2==6 ==>", divide(13, 2) == 6)
	println("33333333/11==3030303 ==>", divide(33333333, 11) == 3030303)
	println("7%5==2 ==>", mod(7, 5) == 2)

	println("8888-888==8000 ==>", sub(8888, 888) == 8000)

	checkbool(true)
	checkbool(false)

	var a int
	println("&a != &b ==>", NeqPtr(&a))

	var e int32
	if timediv(12345*1000000000+54321, 1000000000, &e) != 12345 || e != 54321 {
		println("bad timediv")
	}

	println("9/3==3 ==>", div(9, 3) == 3)

	println("3*3==9 ==>", mul(3, 3) == 9)

	println("fmas:-257.77*256.88+255.99 == -6.595996e+004 ==>", fmas(-257.77, 256.88, 255.99) == (float32(-257.77)*float32(256.88)+float32(255.99)))
	println("fmss:-257.77*256.88-(-255.99) == -6.595996e+004 ==>", fmss(-257.77, 256.88, -255.99) == (float32(-257.77)*float32(256.88)-float32(-255.99)))
	println("fmad:-257.77*256.88+255.99 == -6.595997e+004 ==>", fmad(-257.77, 256.88, 255.99) == (float64(-257.77)*float64(256.88)+float64(255.99)))
	println("fmsd:-257.77*256.88-(-255.99) == -6.595997e+004 ==>", fmsd(-257.77, 256.88, -255.99) == (float64(-257.77)*float64(256.88)-float64(-255.99)))

	TestSyscall()
	syscall.Exit(0)
}

func TestSyscall() {
	testGettimeofday()

	syscall.Getenv("SHELL")
	testMmap()
	testSetGetenv("TESTENV", "AVALUE")
	testSetGetenv("TESTENV", "")
}

func testMmap() {
	b, err := syscall.Mmap(-1, 0, syscall.Getpagesize(), syscall.PROT_NONE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		println("Mmap: ", err)
		syscall.Exit(-1)

	}
	if err := syscall.Munmap(b); err != nil {
		println("Munmap: ", err)
		syscall.Exit(-1)
	}
}

func testReocver() {
	testSIGSEGV()

	defer func() {
		recover()
	}()
	panic("shouldn't panic")
}
func testSIGSEGV() {
	defer func() {
		recover()
	}()
	atomic.AddInt32(nil, 0)
}

func testSetGetenv(key, value string) {
	err := syscall.Setenv(key, value)
	if err != nil {
		println("Setenv failed to set ", value, ": ", err)
		syscall.Exit(-1)
	}
	newvalue, found := syscall.Getenv(key)
	if !found {
		println("Getenv failed to find ", key, " variable (want value ", value, ")")
		syscall.Exit(-1)
	}
	if newvalue != value {
		println("Getenv(", key, ") = \"", newvalue, "\"; want \"", value, "\"")
		syscall.Exit(-1)
	}
}

func testGettimeofday() {
	tv := &syscall.Timeval{}
	if err := syscall.Gettimeofday(tv); err != nil {
		println(err)
		syscall.Exit(-1)
	}
	if tv.Sec == 0 && tv.Usec == 0 {
		println("Sec and Usec both zero")
		syscall.Exit(-1)
	}
}

func issue57() {
	// Operate on numbers of different precision.
	var x, y, z big.Float
	x.SetInt64(1000)          // x is automatically set to 64bit precision
	y.SetFloat64(2.718281828) // y is automatically set to 53bit precision
	z.SetPrec(32)
	z.Add(&x, &y)
	fmt.Printf("x = %.10g (%s, prec = %d, acc = %s)\n", &x, x.Text('p', 0), x.Prec(), x.Acc())
	fmt.Printf("y = %.10g (%s, prec = %d, acc = %s)\n", &y, y.Text('p', 0), y.Prec(), y.Acc())
	fmt.Printf("z = %.10g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
}

type M int

var called = false

func (m M) UniqueMethodName() {
	called = true
}

var v M

func reflectmethod4() {

	fn := reflect.ValueOf(v).Method(0).Interface().(func())
	fn()
	if !called {
		panic("UniqueMethodName not called")
	}

}
