package main

import (
    "vigovia/router"
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080")
}
