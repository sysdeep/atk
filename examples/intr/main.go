package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/sysdeep/atk/tk"
	"github.com/sysdeep/atk/tk/interp"
)

// пример вызова интерпритатора напрямую
func main() {

	runtime.LockOSThread()

	intr, err := interp.NewInterp()
	if err != nil {
		panic(err)
	}

	err = intr.InitTcl("")
	if err != nil {
		panic(err)
	}

	err = intr.InitTk("")
	if err != nil {
		panic(err)
	}

	err = intr.Eval("pack [label .l -text aaaa]")
	if err != nil {
		panic(err)
	}

	fmt.Println(intr.TclVersion())

	interp.MainLoop(func() {
		err = intr.Eval(`
			pack [frame .f] -expand 1 -fill both

			pack [label .f.la -text bbb]
			pack [ttk::progressbar .f.bar1] -padx 20 -pady 20 -expand 1
			
			pack [button .f.b -text Exit -command {destroy .}] -side bottom


			bind . <Control-q> { exit }

		`)
		if err != nil {
			panic(err)
		}

		// try async
		go func() {
			i := 0
			inc := -1
			for {
				if i > 99 || i < 1 {
					inc = -inc
				}
				i += inc
				time.Sleep(5e7)
				vv := fmt.Sprintf(".f.bar1 configure -value {%d}", i)
				tk.Async(func() {
					intr.Eval(vv)
				})
			}
		}()

	})
	fmt.Println("finish")
}
