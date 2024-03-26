# Lad


<div align="center">

Add context logger for zap.

![Zap logo](assets/logo.png)


</div>

## Installation

`go get -u github.com/tnngo/lad`

Note that zap only supports the two most recent minor versions of Go.

## Change

Support for logging contextual metadata, 

### Context

```go
defineContext := Context(func(ctx context.Context) []Field {
  var fields []Field

  if dc, ok := ctx.Value(requestID).(string); ok {
  	fields = append(fields, String(string(requestID), dc))
  }

  return fields
})

logger, _ := lad.NewDevelopment(defineContext)

ctx := context.TODO()
ctx = context.WithValue(ctx, requestID, "123456789")

logMessage := "tnngo"

logger.WithContext(ctx).Info(logMessage)
logger.WithContext(ctx).Debug("1")
logger.Sugar().WithContext(ctx).Debug(2)
```
