package Honey

import (
	"io"
	"time"
)

const (
	reset   = "\033[0m"  // Reset to default color
	red     = "\033[31m" // Red
	green   = "\033[32m" // Green
	yellow  = "\033[33m" // Yellow
	blue    = "\033[34m" // Blue
	magenta = "\033[35m" // Magenta
	cyan    = "\033[36m" // Cyan
	white   = "\033[37m" // White
)

type Loger_Writer struct {
	w *io.Writer
}
type Loger_Writer_Two struct {
	First  *io.Writer
	Second *io.Writer
}
type LOGGER interface {
	Error(err error)
	Info(message string)
	Panic_Err(err error)
}

func New_Two_Writer(fisrt io.Writer, second io.Writer) LOGGER {
	return &Loger_Writer_Two{
		First:  &fisrt,
		Second: &second,
	}
}
func New_OneLogger(w io.Writer) Loger_Writer {
	return Loger_Writer{w: &w}
}
func (l Loger_Writer) Error(err error) {
	io.WriteString(*l.w, cyan+time.Now().String()+reset+red+"  ->> ERROR : \n\n "+err.Error()+reset)
}
func (l Loger_Writer) Info(message string) {
	io.WriteString(*l.w, cyan+time.Now().String()+reset+"   ->> MESSAGE : \n\n"+message)
}
func (l Loger_Writer) Panic_Err(err error) {
	io.WriteString(*l.w, cyan+time.Now().GoString()+red+"   ->> PANIC : \n\n"+err.Error()+reset)
}

func (l *Loger_Writer_Two) Error(err error) {
	go Write_Error(l.First, err)
	go Write_Error(l.Second, err)
}

func (l *Loger_Writer_Two) Info(message string) {
	go Write_Info(l.First, message)
	go Write_Info(l.Second, message)
}

func (l *Loger_Writer_Two) Panic_Err(err error) {
	go Write_Panic(l.First, err)
	go Write_Panic(l.Second, err)
}
