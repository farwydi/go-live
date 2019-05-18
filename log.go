package main

import (
    "bytes"
    "os"
    "sync"
)

var logFile *os.File

var queueLog = make(chan string, 1)
var bufferLog = make([]string, bufferSize)
var workLog sync.WaitGroup
var counterLog = 0

const bufferSize = 10000

func init() {
    var err error

    os.Remove(*LogFile)
    logFile, err = os.OpenFile(*LogFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)

    if err != nil {
        panic(err)
    }
}

var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func doneLog() {
    b := bufferPool.Get().(*bytes.Buffer)
    for i, s := range bufferLog {
        if i >= counterLog {
            break
        }
        b.WriteString(s)
    }

    logFile.Write(b.Bytes())
}

func saverLog() {
    for t := range queueLog {
        if counterLog >= bufferSize {
            counterLog = 0

            b := bufferPool.Get().(*bytes.Buffer)
            for _, s := range bufferLog {
                b.WriteString(s)
            }

            logFile.Write(b.Bytes())
            b.Reset()
            bufferPool.Put(b)
        }

        bufferLog[counterLog] = t
        counterLog++
        workLog.Done()
    }
}

func log(msg string) {
    if *PrintActionPtr {
        if epoch % *GenomeSkip == 0 {
            workLog.Add(1)
            queueLog <- msg
        }
    }
}
