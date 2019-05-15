package main

import "os"

var logFile *os.File

func init() {
    var err error

    os.Remove(*LogFile)
    logFile, err = os.OpenFile(*LogFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)

    if err != nil {
        panic(err)
    }
}

func log(msg string) {
    if *PrintActionPtr {
        logFile.Write([]byte(msg))
    }
}
