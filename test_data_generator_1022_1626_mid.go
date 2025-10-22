// 代码生成时间: 2025-10-22 16:26:22
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// UserData represents the user data structure
type UserData struct {
    ID       int
    Name     string
    Email    string
    Password string
}

// NewUserData generates a new UserData instance with random data
func NewUserData() UserData {
    rand.Seed(time.Now().UnixNano())
    return UserData{
        ID:       rand.Intn(10000),
        Name:     fmt.Sprintf("User%d", rand.Intn(100)),
        Email:    fmt.Sprintf("user%d@example.com", rand.Intn(100)),
        Password: fmt.Sprintf("password%d", rand.Intn(100)),
    }
}

// GenerateTestData creates a slice of UserData instances
func GenerateTestData(count int) ([]UserData, error) {
    var users []UserData
    for i := 0; i < count; i++ {
        user := NewUserData()
        users = append(users, user)
    }
    return users, nil
}

func main() {
    count := 10 // Number of test data to generate
    testData, err := GenerateTestData(count)
    if err != nil {
        fmt.Println("Error generating test data: ", err)
        return
    }

    fmt.Println("Generated Test Data:")
    for _, user := range testData {
        fmt.Printf("ID: %d, Name: %s, Email: %s, Password: %s
", user.ID, user.Name, user.Email, user.Password)
    }
}