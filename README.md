# log
### Usage

```go
import . "github.com/ivierlin/log"

func init() {
  Debug("init")
 }
  
func main() {
  Log(Prefix("SYMPATHY", Magenta), "Is this the real life?")
  i := 42
  Info("var i =",i)
  Debug("exit")
 }
```
