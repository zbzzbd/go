package fmt
import (
	"errors"
	"io"
	"math"
	"os"
	"reflect"
	"strconv"
	"sync"
	"unicode/utf8"
	"debug/elf"
	"fmt"
)

type runeUnreader interface {
	UnreadRune() error
}
type ScanState interface {
	ReadRune() (r rune, size int ,err error)
	UnreadRune() error
	SkipSpace()

	Token(skipSpace bool, f func(rune) bool) (token []byte ,err error)
	Width() (wid int ,err error)
	Read(buf []byte) (n int ,er error)
}

type Scanner interface {
	Scan(state ScanState,verb rune) error
}

func Scan(a ...interface{}) (n int,err error) {
	return Fscan(os.Stdin,a...)
}

func Scanf(format string , a ...interface{}) (n int ,err error) {
	return Fscanf(os.Stdin,format,a...)
}
func Scanln(a ...interface{}) (n int, err error){
	return Fscanln(os.Stdin, a...)
}
type  stringReader string

func (r *stringReader) Read(b []byte) (n int ,err error)  {
	n = copy(b,*r)
	*r = (*r)[n:]
	if n ==0 {
		err = io.EOF
	}
	return
}

func Sscanf(str string,format string , a ...interface{}) (n int ,err error)  {
	return Fscanf((*stringReader(&str)),format,a...)
}
func Fscan(r io.Reader, a ...interface{}) (n int ,err error)  {
	s,old := newScanState(r,true,false)
	n, err = s.doScan(a)
	s.free(old)
	return
}
func Fscanf(r io.Reader, formate string ,a ...interface{}) (n int ,err error) {
	s,old :=newScanState(r,flase ,flase)
	n, err = s.doScanf(format,a )
	s.free(old)
	return
}

type scanError struct {
	err error
}
const  eof  = -1
type ss struct {
	rr io.RuneReader
	buf buffer
	peekRune rune
	prevRune rune
	count int
	atEOF bool
	ssave
}

type  ssave struct {
	validSave bool
	nlIsEnd bool
	nlIsSpace bool
	argLimit int
	limit int
	maxWid int
}
func (s *ss) Read(buf []byte) (n int ,err error) {
	return 0,errors.New("ScanState's Read should not be called .Use ReadRune")
}
func (s *ss) ReadRune() (r rune ,size int ,err error) {
	if s.peekRune>=0 {
		s.count++
		r = s.peekRune
		size = utf8.RuneLen(r)
		s.prevRune = r
		s.peekRune = -1
		return
	}
	if s.atEOF || s.nlIsEnd && s.prevRune == '\n' || s.count>= s.argLimit {
		err = io.EOF
		return
	}
	r,size,err=s.rr.ReadRune()
	if err == nil {
		s.count++
		s.prevRune =r
	}else if err = io.EOF {
		s.atEOF = true
	}
	return

}

func (s *ss)Width() (wid int ,ok bool) {
	if s.maxWid == bugWid{
		return 0,false
	}
	return s.maxWid,true
}

func (s *ss) getRune() (r rune)  {
	if u,ok:=s.rr.(runeUnreader);ok {
		u.UnreadRune()
	}else {
		s.peekRune = s.prevRune
	}
	s.prevRune = -1
	s.count--
	return nil
}
func (s *ss) mustReadRune() (r rune)  {
	r = s.getRune()
	if r ==eof {
		s.error(io.ErrUnexpectedEOF)
	}
	return
}
func (s *ss) UnreadRune() error  {
	if u,ok :=s.rr.(runeUnreader);ok {
		u.UnreadRune()
	}else {
		s.peekRune = s.prevRune
	}
	s.prevRune =-1
	s.count--
	return nil
}

func (s *ss) error() (err error)  {
	panic(scanError{errors.New(err)})
}

func (s *ss) errorString() (err string) {
	panic(scanError{errors.New(err)})
}
func (s *ss) Token(skipSpace bool, f func(rune) bool) (tok []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			if se, ok := e.(scanError); ok {
				err = se.err
			} else {
				panic(e)
			}
		}
	}()
	if f == nil {
		f = notSpace
	}
	s.buf = s.buf[:0]
	tok = s.Token(skipSpace,f)
	return
}
var space = [][2]uint16 {
	{0x0009,0x000d},
	{0x0020,0x0020},
	{0x0085,0x0085},
	{0x00a0,0x00a0},
	{0x1680,0x1680},
	{0x2000,0x200a},
	{0x2028,0x2029},
	{0x202f,0x202f},
	{0x205f,0x205f},
	{0x3000,0x3000},

}

func isSpace (r rune) bool {
	if r >=1 <<16{
		return false
	}
	rx := uint16(r)
	for _,rng:=range space {
		if rx <rng[0] {
			return false
		}
		if rx <= rng[1] {
			return true
		}
	}
	return false
}
type   readRune struct {
  	reader io.Reader
	buf	[utf8.UTFMax]byte
	pending int
	pendBuf [utf8.UTFMax]byte
}


func (r *readRune) readByte() (b byte, err error) {
	if r.pending >0 {
		b = r.pendBuf[0]
		copy(r.pendBuf[0:],r.pendBuf[1:])
		r.pending --
		return
	}
	n ,err := io.ReadFull(r.reader, r.pendBuf[0:1])
    if n != 1 {
		return  0, err
	}
	return  r.pendBuf[0],err
}
