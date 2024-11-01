# UUIDv7 Generator for Golang

**Key idea behind this implementation** is that this lib can be used to generate UUIDs from specific timestamps, which is useful for generating UUIDs for historical data. Actually, it's the only reason why I created this lib.

Ever needed to convert a dozen thosuand records from a database into UUIDs v7 to meet some new weird criteria? This lib is for you. Never? Lucky you.

It also has a CLI tool that can be used to generate UUIDs from a list of timestamps.

**Features**

* UUID v7 Compatibility: Implements the draft specification for UUID version 7, creating time-ordered, unique 128-bit identifiers.
* Timestamp-Based Generation: Uses current or specified timestamps, ensuring uniqueness at up to 10,000 UUIDs per second.
* Simple API: Flexible generation options, supporting standard timestamp-based UUIDs and custom timestamps.



**Limitations**

This initial implementation is optimized for generating **up to 10,000 UUIDs per second**. For higher throughput, a sequence marker or additional unique node ID would be necessary and this is not currently supported.

**Installation**

`go get github.com/lazybark/go-uuid-historical`

Usage

Generate a UUID using the current timestamp:

```go
package main

import (
    "fmt"
    "github.com/lazybark/go-uuid-historical/uuidv7"
)

func main() {
    uuid, err := uuidv7.GenerateUUIDv7()
    if err != nil {
        fmt.Println("Error generating UUID:", err)
        return
    }
    fmt.Println("Generated UUID v7:", uuid)
}
```

Generate a UUID using a custom timestamp:

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazybark/go-uuid-historical/uuidv7"
)

func main() {
   timestamp := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
   
   uuid, err := uuidv7.GenerateUUIDv7WithTimestamp(timestamp)
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
```

**CLI Tool**

You can also use the provided CLI tool to generate UUIDs from timestamps. The tool reads timestamps from standard input and outputs the corresponding UUIDs.

Use `-layouts` flag to specify the timestamp layouts to parse. The tool will attempt to parse each timestamp using the layouts in order, stopping at the first successful parse. If no layouts are provided, the tool will use the default layout "2006-01-02 15:04:05.000 -0700" and several other common layouts.

**Usage:**

`./uuidv7-cli -layouts "2006-01-02 15:04:05.000 -0700,2006-01-02 15:04:05 -0700,2006-01-02T15:04:05.000Z,2006-01-02T15:04:05Z,2006-01-02 15:04:05,2006-01-02"`