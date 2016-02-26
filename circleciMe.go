package main

import (
  "fmt"
  "encoding/json"
  "gopkg.in/h2non/gentleman.v0"
)

type Output struct {
    In_beta_program bool
    Selected_email string
    Avatar_url string
    Name string
    Login string
    Github_id int32
}

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define base URL
  cli.URL("https://circleci.com/api/v1/me?circle-token=" + authToken)

  // Create a new request based on the current client
  req := cli.Request()
  req.SetHeader("Accept", "application/json")

  // Perform the request
  res, err := req.Send()
  if err != nil {
    fmt.Printf("Request error: %s\n", err)
    return
  }
  if !res.Ok {
    fmt.Printf("Invalid server response: %d\n", res.StatusCode)
    return
  }

  // Reads the whole body and returns it as string
//  fmt.Printf("%s\n", res.String())
  fmt.Printf("%s\n", res.Bytes())

    output := Output{}
    err = json.Unmarshal(res.Bytes(), &output)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("%#v\n", output)
}
