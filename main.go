/*
Serve is a very simple static file server in go
Usage:
        -p="443": port to serve on
        -d=".":    the directory of static files to host
        -c="./example.com.crt" certFile
        -k="./example.com.key" keyFile
Navigating to https://example.com will display the index.html or directory
listing file.
*/
package main

import (
        "flag"
        "log"
        "net/http"
)

func main() {
        port := flag.String("p", "8080", "port to serve on")
        directory := flag.String("d", ".", "the directory of static file to host")
        certFile := flag.String("c", ".", "certFile")
        keyFile := flag.String("k", ".", "keyFile")
        flag.Parse()

        http.Handle("/", http.FileServer(http.Dir(*directory)))

        log.Printf("Serving %s on HTTP port: %s\ncertFile: %s, keyFile: %s\n", *directory, *port, *certFile, *keyFile)
        log.Fatal(http.ListenAndServeTLS(":"+*port, certFile, keyFile, nil))
}
