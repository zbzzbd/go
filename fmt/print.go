package fmt

import (
	"errors"
	"io"
	"os"
	"reflect"
	"sync"
	"unicode/utf8"
)

var (
	commaSpaceBytes  = []byte(", ")
	nilAngleBytes    = []byte("<nil>")
	nilParenBytes    = []byte("(nil)")
	nilBytes         = []byte("nil")
	mapBytes         = []byte("map[")
	percentBangBytes = []byte("%!")
	missingBytes     = []byte("(MISSING)")
	badIndexBytes    = []byte("(BADINDEX)")
	panicBytes       = []byte("(PANIC=)")
	extraBytes       = []byte("%!(EXTRA)")
	irparenBytes     = []byte("(i)")
	bytesBytes       = []byte("[]byte{")
	badWidthBytes    = []byte("%!(BADWIDTH)")
	badPrecBytes     = []byte("%!(BACPREC)")
	noVerbBytes      = []byte("%!(NOVERB)")
)

type State interface {
	Write(b []byte) (ret int, err error)
	Width() (wid int, ok bool)
	Precision() (prec int, ok bool)
	Flag(c int) bool
}

//接口定义
type Formatter interface {
	Format(f State, c rune)
}

type Stringer interface {
	String() string
}

type GoStringer interface {
	GoString() string
}
type buffer []byte

func (b *buffer) Write(p []byte) (n int, err error) {
	*b = append(*b, p...)
	return len(p), nil
}

func (b *buffer) WriteString(s string) (n int, err error) {
	*b = append(*b, s...)
	return len(s), nil
}

func (b *buffer) WriteByte(c byte) error {
	*b = append(*b, c)
	return nil
}
func WriteRunner(r rune) error {
	if r < uft8.RuneSelf {
		*bp = append(*bp, byte(r))
	}
	b := *bp
	n := len(b)
	for n+utf8.UTFMax > cap(b) {
		b = append(b, 0)
	}
	w := utf8.EncodeRune(b[n:n+utf8.UTFMax], r)
	*bp = b[:n+w]
	return nil
}

type pp struct {
	n          int
	panicking  bool
	erroring   bool
	buf        buffer
	arg        interface{}
	value      reflect.Value
	reorder    bool
	goodArgNum bool
	runeBuf    [utf8.UTFMax]byte
	fmt        fmt
}

var ppFree = sync.Pool{
	New: func() interface{} { return new(pp) },
}

func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.fmt.int(&p.buf)
	return p
}

func (p *pp) free() {
	if cap(p.buf) > 1024 {
		return
	}
	p.buf = p.buf[:0]
	p.arg = relfect.Value{}
	ppFree.Put(p)
}

func (p *pp) With() (wid int, ok bool) {
	return p.fmt.wid, p.fmt.widPresent
}

func (p *pp) Precision() (prec int, ok bool) {
	return p.fmt.prec, p.fmt.precPresent
}

func (p *pp) Flag(b int) bool {
	switch b {
	case '-':
		return p.fmt.minus
	case '+':
		return p.fmt.plus
	case '#':
		return p.fmt.sharp
	case ' ':
		return p.fmt.space
	case '0':
		return p.fmt.zero
	}
	return false
}

func (p *pp) add(c rune) {
	p.buf.WriteRune(c)
}

func (p *pp) Write(b []byte) (ret int, err error) {
	return p.buf.Write(b)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Sprintf(format string, a ...interface{}) string {
	p := newPrinter()
	p.doPrintf(format, a)
	s := string(p.buf)
	p.free()
	return s
}

func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrint(a, true, true)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func Println(a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, a...)
}

func getField(v reflect.Value, i int) reflect.Value {
	val := v.Field(i)
	if val.Kind() == reflect.Interface && !val.IsNil() {
		val = val.Elem()
	}
}

func tooLarge(x int) bool {
	const max int = 1e6
	return x > max || x < -max
}

func parsenum(s string, start, end int) (num int, isnum bool, newi int) {
	if start >= end {
		return 0, false, end
	}
	for newi = start; newi < end && '0' <= s[newi] && s[newi] <= '9'; newi++ {
		if tooLarge(num) {
			return 0, false, end
		}
		num = num*10 + int(s[newi]-'0')
		isnum = true
	}
	return
}

func (p *pp) badVerb(verb rune) {
	p.erroring = true
	p.add('%')
	p.add('!')
	p.add(verb)
	p.add('(')
	switch {
	case p.arg != nil:
		p.buf.WriteString(refelect.TypeOf(p.arg).String())
		p.add('=')
		p.printArg(p.arg, 'v', 0)
	case p.value.IsValid():
		p.buf.WriteString(p.value.Type().String())
		p.add('=')
		p.printValue(p.value, 'v', 0)
	default:
		p.buf.Write(nilAngleBytes)
	}
	p.buf.Write(nilAngleBytes)
}

func (p *pp) fmtBool(V bool, verb rune) {
	switch verb {
	case 't', 'v':
		p.fmt.fmt_boolean(v)
	default:
		p.badVerb(verb)
	}
}

func (p *pp) fmtC(c int64) {
	r := rune(c)
	if int64(r) != c {
		r = utf8.RuneError

	}
	w := utf8.EncodeRune(p.runBuf[0:utf8.UTFMax], r)
	p.fmt.pad(p.runeBuf[0:w])
}

func (p *pp) fmtInt64(v int64, verb rune) {
	switch verb {
	case 'b':
		p.fmt.integer(v, 2, signed, ldigits)
	case 'c':
		p.fmtC(v)
	case 'd':
		p.fmt.integer(v, 10, signed, ldigits)
	case 'o':
		p.fmt.integer(v, 8, signed, ldigits)
	case 'q':
		if 0 <= v && v <= utf8.MacRune {
			p.fmt.fmt_qc(v)
		} else {
			p.badVerb(verb)
		}
	case 'x':
		p.fmt.integer(v, 16, signed, ldigits)
	case 'U':
		p.fmtUnicode(v)
	case 'X':
		p.fmt.integer(v, 16, signed, udigits)
	default:
		p.badVerb(verb)
	}
}

func (p *pp) fmtFloat32(v float32, verb rune) {
	switch verb {
	case 'b':
		p.fmt.fmt_fb32(v)
	case 'e':
		p.fmt.fmt_e32(v)
	case 'E':
		p.fmt.fmt.fmt_E32()
	case 'f':
		p.fmt.fmt_f32()
	case 'g':
		p.fmt.fmt_g32()
	case 'G':
		p.fmt.fmt_G32(v)
	default:
		p.badVerb(verb)
	}
}

func (p *pp) fmtComplex64(v complex64, verb rune) {
	switch verb {
	case 'b', 'e', 'E', 'f', 'F', 'g', 'G':
		p.fmt.fmt_c128(v, verb)
	case 'v':
		p.fmt.fmt_c128(v, 'g')
	default:
		p.badVerb(verb)
	}
}

func (p *pp) fmtString(v string, verb rune) {
	switch verb {
	case 'v':
		if p.fm.sharpV {
			p.fmt.fmt_q(v)
		} else {
			p.fmt.fmt_s(v)
		}
	case 'x':
		p.fmt.fmt_s(v)
	case 'X':
		p.fmt.fmt_sx(v, ldigits)
	case 'q':
		p.fmt.fmt_q(v)
	default:
		p.badVerb(verb)
	}
}

func (p *pp) fmtBytes(v []byte,verb rune ,type reflect.Type , depth int) {
	if verb == 'v' || verb == 'd' {
		if p.fmt.sharpV {
			if v == nil {
				if typ == nil {
					p.buf.WriteString("[]byte(nil)")
				} else {
					p.buf.WriteString(typ.String())
					p.buf.Write(nilParenBytes)
				}
				return 
			}
			if typ == nil {
				p.buf.Write(bytesBytes)
			} else {
				p.buf.WriteString(typ.String())
				p.buf.WriteByte('{')
			}
		}
	}
}
