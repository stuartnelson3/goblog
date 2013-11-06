package main

import (
    "os/exec"
    "strings"
    "bytes"
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/hex"
    "bufio"
    "os"
    "code.google.com/p/gopass"
)

func main() {
    var app string

    envVariables :=
        map[string]string{"BLOGTOKEN":"", "HASHED_USERNAME":"", "HASHED_PASSWORD":""}

    GetAppName(&app)

    envVariables = GetEnvVariables(envVariables)

    for envVar, value := range envVariables {
        h := sha512.New()
        io.WriteString(h, value)
        hashedValue := hex.EncodeToString(h.Sum(nil))
        SetEnvVar(envVar, hashedValue, app)
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

func GetAppName(app *string) {
    var stdout, stderr bytes.Buffer
    fmt.Printf("\nPick your app from the list of available apps:\n")

    cmd := exec.Command("heroku", "apps")
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    cmd.Run()
    if err := stderr.String(); len(err) > 0 {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(stdout.String())

    fmt.Printf("What is the name of your app?\n")
    fmt.Scanln(app)
}

func SetEnvVar(envVar string, hashedValue string, app string) {
    var stdout bytes.Buffer
    cmd := exec.Command("echo", "heroku", "config:set", envVar+"="+hashedValue, "--app", app)
    cmd.Stderr = &stdout
    cmd.Run()
    if err := stdout.String(); len(err) > 0 {
        fmt.Println(err)
        return
    }
    fmt.Printf("Successfully set %s for app %s\n", envVar, app)
}
