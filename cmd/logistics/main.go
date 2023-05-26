package main

func main() {
    app, cleanup, err := newWire()
    if err != nil {
        panic(err)
    }
    defer cleanup()

    // start and wait for stop signal
    if e := app.Run(); e != nil {
        panic(e)
    }
}
