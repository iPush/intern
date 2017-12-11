package intern

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestIntern(t *testing.T) {
	GetInstance()
	s := "abcde"
	interned := Intern(s)

	want := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	got := (*reflect.StringHeader)(unsafe.Pointer(&interned)).Data
	if want != got {
		t.Errorf("failed to intern string")
	}

	third := fmt.Sprintf("%s%s", "abc", "de")
	interned = Intern(third)
	want = (*reflect.StringHeader)(unsafe.Pointer(&third)).Data
	got = (*reflect.StringHeader)(unsafe.Pointer(&interned)).Data
	if want == got {
		t.Errorf("failed to intern string")
	}

}

func BenchmarkIntern(b *testing.B) {
	GetInstance()

	a := "this is a test text"
	Intern(a)

	c := "this is a very long abcdefg hijklmn opqrst uvwxyz fadskjflksdajflkjsadklfjklasdjfksadjf;klsdjaflksdjalkfjsl;kajfklsdajflk"
	Intern(c)

	b.ReportAllocs()
	b.N = 500000000
	b.SetBytes(int64(len(a) + len(c)))
	b.RunParallel(func(pb *testing.PB) {
		var s string
		var t string
		for pb.Next() {
			s = Intern(fmt.Sprintf("%s%s", "this is", "a test text"))
			t = Intern(fmt.Sprintf("%s", "this is a very long abcdefg hijklmn opqrst uvwxyz fadskjflksdajflkjsadklfjklasdjfksadjf;klsdjaflksdjalkfjsl;kajfklsdajflk"))
		}
		_ = s
		_ = t
	})
}

func BenchmarkString(b *testing.B) {
	a := "this is a short text"
	c := "this is a very long abcdefg hijklmn opqrst uvwxyz fadskjflksdajflkjsadklfjklasdjfksadjf;klsdjaflksdjalkfjsl;kajfklsdajflk"
	b.ReportAllocs()
	b.N = 500000000

	b.SetBytes(int64(len(a) + len(c)))
	b.RunParallel(func(pb *testing.PB) {
		var s string
		var t string
		for pb.Next() {
			s = fmt.Sprintf("%s%s", "this is", "a test text")
			t = fmt.Sprintf("%s%s", "this is a very long abcdefg hijklmn opqrst uvwxyz fadskjflksdajflkjsadklfjklasdjfksadjf;klsdjaflksdjalkfjsl;kajfklsdajflk")
		}
		_ = s
		_ = t
	})

}
