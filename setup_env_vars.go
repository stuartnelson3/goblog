package main

import (
    "os/exec"
    "bytes"
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/hex"
    "time"
)

func main() {
    var envVar, app, input string
    var out bytes.Buffer

    GetAppName(&app, &out)
    GetVariable(&envVar)
    GetValueToHash(&input, envVar)

    h := sha512.New()
    io.WriteString(h, input)
    hashedValue := hex.EncodeToString(h.Sum(nil))

    SetEnvVar(envVar, hashedValue, app, &out)
}

func GetValueToHash(input *string, envVar string) {
    fmt.Printf("Enter the value to be hashed for %s\n", envVar)
    fmt.Scanln(input)
}

func GetVariable(env *string) {
    fmt.Println("Enter the environment variable to set:")
    fmt.Println("(options: HASHED_PASSWORD, HASHED_USERNAME, BLOGTOKEN")
    fmt.Scanln(env)
}

func GetAppName(app *string, out *bytes.Buffer) {
    fmt.Printf("\nPick your app from the list of available apps:\n")
    time.Sleep(time.Millisecond * 500)

    cmd := exec.Command("heroku", "apps")
    cmd.Stdout = out
    cmd.Stderr = out
    cmd.Run()
    fmt.Println(out.String())
    out.Reset()

    fmt.Printf("What is the name of your app?\n")
    fmt.Scanln(app)
}

func SetEnvVar(envVar string, hashedValue string, app string, out *bytes.Buffer) {
    cmd := exec.Command("heroku", "config:set", envVar+"="+hashedValue, "--app", app)
    cmd.Stdout = out
    cmd.Stderr = out
    cmd.Run()
    fmt.Println(out.String())
}
