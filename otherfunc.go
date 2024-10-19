package Honey

import (
	"io"
	"time"
)

func Write_Error(w *io.Writer, err error) {
	io.WriteString(*w, cyan+time.Now().String()+reset+red+"  ->> ERROR : \n\n "+err.Error()+reset + "]]]")
}
func Write_Panic(w *io.Writer, err error) {
	io.WriteString(*w, cyan+time.Now().GoString()+red+"   ->> PANIC : \n\n"+err.Error()+reset + "]]]")
}
func Write_Info(w *io.Writer, message string) {
	io.WriteString(*w, cyan + time.Now().String() + reset + "   ->> MESSAGE : \n\n" +message + "]]]")
}
