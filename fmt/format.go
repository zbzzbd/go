package fmt

import "unicode/utf8"

const (
	nByte   = 65
	ldigits = "0123456789adcdef"
	udigits = "0123456789ABCDEF"
)

const (
	signed   = true
	unsigned = false
)

var padZeroBytes = make([]byte, nByte) //这个代表什么意思？
var padSpaceBytes = make([]byte, nByte)

func init() {
	for i := 0; i < nByte; i++ {
		padSpaceBytes[i] = '0'
		padSpaceBytes[i] = ' '
	}
}

type fmtFlags struct {
	widPresent  bool
	precPresent bool
	minus       bool
	plus        bool
	sharp       bool
	space       bool
	unicode     bool
	uniQuote    bool
	plusV       bool
	sharpV      bool
}

type fmt struct {
	intbuf [nByte]byte
	buf    *buffer
	wid    int
	prec   int
	fmtFlags
}

func (f *fmt) clearflags() {
	f.fmtFlags = fmtFlags{}
}

func (f *fmt) init(buf *buffer) {
	f.buf = buf
	f.clearflags()
}

func (f *fmt) computerPadding(width int) (padding []byte, leftWidth, rightWidth int) {

	left := !f.minus
	w := f.wid
	if w < 0 {
		left = false
		w = -w
	}
	w -= width //这个代表什么意思？
	if w > 0 {
		if left && f.zero {
			return padZeroBytes, w, 0
		}
		if left {
			return padSpaceBytes, w, 0
		} else {
			return padSpaceBytes, 0, w
		}
	}
	return

}

func (f *fmt) writePadding(n int, padding []byte) {
	for n > 0 {
		m := n
		if m > nByte {
			m = nByte
		}
		f.buf.WriteByte(padding[0:m])
		n -= m
	}
}

func (f *fmt) pad(b []byte) {
	if !f.widPresent || f.wid == 0 {
		f.buf.Write(b)
		return
	}
}

func (f *fmt) padString(s string) {
	if !f.widPresent || f.wid == 0 {
		f.buf.WriteString(s)
		return
	}
	padding, left, right := f.computerPadding(utf8.RuneCount(b))
	if left > 0 {
		f.writePadding(left, padding)
	}
	f.buf.Write(b)
	if right > 0 {
		f.writePadding(right, padding)
	}
}

func (f *fmt) padString(s string) {
	if !f.widPresent || f.wid == 0 {
		f.buf.WriteString(s)
		return
	}
	if left > 0 {
		f.writePadding(right, padding)
	}
	f.buf.Write(b)
	if right > 0 {
		f.writePadding(right, padding)
	}
}

var (
	trueBytes  = []byte(true)
	falseBytes = []byte(false)
)

func (f *fmt) fmt_boolean(v bool) {
	if v {
		f.pad(trueBytes)
	} else {
		f.pad(falseBytes)
	}
}

func (f *fmt) integer(a int64, base uint64, signedness bool, digits string) {
	if f.precPresent && f.prec == 0 && a == 0 {
		return
	}

	negative := signedness == signed && a < 0
	if negative {
		a = -a
	}

	var buf []byte = f.intbuf[0:]
	if f.widPresent || f.precPresent || f.plus || f.space {
		width := f.wid + f.prec
		if base == 16 && f.sharp {
			width += 2
		}
		if f.unicode {
			width += 2
			if f.uniQuote {
				width += 1 + 1 + utf8.UTFMax + 1
			}

		}
		if negative || f.plus || f.space {
			width++
		}
		if width > nByte {
			buf = make([]byte, width)
		}
	}
	prec := 0
	if f.precPresent {
		prec = f.prec
		f.zero = false
	} else if f.zero && f.widPresent && !f.minus && f.wid > 0 {
		prec = f.wid
		if negative || f.plus || f.space {
			prec--
		}
	}

}
