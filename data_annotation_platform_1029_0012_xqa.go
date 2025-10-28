// 代码生成时间: 2025-10-29 00:12:51
package main

import (
    "github.com/revel/revel"
    "{{ .AppImportPath }}/app" // The path to your application's main package
)

// Main function is the entry point of the application
func main() {
    // Initialize the Revel application
    revel.Init(
        revel.NewFileSystemAppConfig(),
    )

    // Run the application
    revel.Run(
        // Register your application's controllers
        app.NewApp(),
    )
}
