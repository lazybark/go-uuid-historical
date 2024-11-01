package examples

import (
	"fmt"
	"time"

	"github.com/lazybark/go-uuid-historical/uuidv7"
)

func Usage() {
	timestamp := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	uuid, err := uuidv7.GenerateUUIDv7(timestamp)
	if err != nil {
		fmt.Println("Error generating UUID:", err)

		return
	}

	fmt.Println("Generated UUID v7:", uuid)

	layout := "2006-01-02 15:04:05.000 -0700"

	t, err := time.Parse(layout, "2024-10-25 08:58:09.662 +0000")
	if err != nil {
		fmt.Println(err)
	}

	uuid, err = uuidv7.GenerateUUIDv7(t)
	if err != nil {
		fmt.Println("Error generating UUID:", err)

		return
	}

	fmt.Println("UUID v7 (custom time):", uuid)
}
