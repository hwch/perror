package perror

import (
        "fmt"
        "os"
        "path"
        "runtime"
        "strings"
)
// no dir, package information
func Perror(err error) {
        pc, file, line, ok := runtime.Caller(1)
        if !ok {
                fmt.Fprintf(os.Stderr, "It was not possible to recover the information.\n")
        }
        if f := runtime.FuncForPC(pc); f == nil {
                fmt.Fprintf(os.Stderr, "Can not get the function name.\n")
        } else {
                fname := f.Name()
                base := path.Base(file)
                pos := strings.LastIndex(fname, ".")
                fmt.Fprintf(os.Stderr,
                        "filename[%s],line[%d],function[%s]:%v\n",
                        base, line, fname[pos+1:], err)
        }
}
// more detailed information
func Verror(err error) {
        pc, file, line, ok := runtime.Caller(1)
        if !ok {
                fmt.Fprintf(os.Stderr, "It was not possible to recover the information.\n")
        }
        if f := runtime.FuncForPC(pc); f == nil {
                fmt.Fprintf(os.Stderr, "Can not get the function name.\n")
        } else {
                fname := f.Name()
                pos := strings.LastIndex(fname, ".")
                base := path.Base(file)
                dir := path.Dir(file)
                fmt.Fprintf(os.Stderr,
                        "directory[%s],filename[%s],line[%d],package[%s],function[%s]:%v\n",
                        dir, base, line, fname[:pos], fname[pos+1:], err)
        }
}
