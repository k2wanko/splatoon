# IkaLink Client for Splatoon

## Usage
```go
import(
  "os"

  "github.com/k2wanko/splatoon"
)

func main() {
  c := splatoon.NewClient(nil)
  c.Auth(os.Getenv("N_USERNAME"), os.Getenv("N_PASSWORD"))
  ss, _ := c.Schedules()
}

```
