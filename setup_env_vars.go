package main

import (
    "os/exec"
    "strings"
    "bytes"
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/hex"
    "time"
    "bufio"
    "os"
    "code.google.com/p/gopass"
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
    var password string
    for envVar, _ := range envVariables {
        prompt := "Enter the value to be hashed for " + envVar + "\n"
        if envVar == "HASHED_PASSWORD" {
            password, _ = gopass.GetPass(prompt)
            scanner = bufio.NewScanner(strings.NewReader(password))
        } else {
            fmt.Printf(prompt)
        }
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
