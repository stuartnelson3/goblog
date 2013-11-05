package main

import (
    "os/exec"
    "bytes"
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/hex"
    "time"
    "bufio"
    "os"
)

func main() {
    var app string
    var out bytes.Buffer

    envVariables :=
        map[string]string{"BLOGTOKEN":"", "HASHED_USERNAME":"", "HASHED_PASSWORD":""}

    GetAppName(&app, out)

    envVariables = GetEnvVariables(envVariables)

    for envVar, value := range envVariables {
        h := sha512.New()
        io.WriteString(h, value)
        hashedValue := hex.EncodeToString(h.Sum(nil))
        SetEnvVar(envVar, hashedValue, app, out)
    }

    fmt.Println("Finished setting up environment variables.")
}

func GetEnvVariables(envVariables map[string]string) map[string]string {
    scanner := bufio.NewScanner(os.Stdin)
    for envVar, _ := range envVariables {
        fmt.Printf("Enter the value to be hashed for %s\n", envVar)
        scanner.Scan()
        envVariables[envVar] = scanner.Text()
    }
    return envVariables
}

func GetAppName(app *string, out bytes.Buffer) {
    fmt.Printf("\nPick your app from the list of available apps:\n")
    time.Sleep(time.Millisecond * 500)

    cmd := exec.Command("heroku", "apps")
    cmd.Stdout = &out
    cmd.Stderr = &out
    cmd.Run()
    fmt.Println(out.String())

    fmt.Printf("What is the name of your app?\n")
    fmt.Scanln(app)
}

func SetEnvVar(envVar string, hashedValue string, app string, out bytes.Buffer) {
    cmd := exec.Command("heroku", "config:set", envVar+"="+hashedValue, "--app", app)
    cmd.Stdout = &out
    cmd.Stderr = &out
    cmd.Run()
    fmt.Println(out.String())
}
