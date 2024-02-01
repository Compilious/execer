package main

func main() {
cmd := exec.Command("go","run","child.go")
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr 
err := cmd.Start()
if err != nil {
_, _ = fmt.Fprint(os.Stderr, "start: " + err.Error())
return
}

time.Sleep(time.Millisecond * 500)
err = cmd.Process.Signal(os.Interrupt)
if err != nil {
_, _ = fmt.Fprintf(os.Stderr , "Signal: " + err.Error())
return 
}

err = cmd.Wait()
if err != nil {
_, _ = fmt.Fprintf(os.Stderr, "Wait: " + err.Error())
}
return
}
