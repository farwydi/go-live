package main

import "os"

var logFile *os.File

func init() {
    var err error

    logFile, err = os.OpenFile("./sim.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)

    if err != nil {
        panic(err)
    }
}

func log(msg string) {
    if PrintLog {
        logFile.Write([]byte(msg))
    }
}
