# Lad


<div align="center">

Add context logger for zap.

![Zap logo](assets/logo.png)

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

</div>

## Installation

`go get -u github.com/tnngo/lad`

Note that zap only supports the two most recent minor versions of Go.

## Change

Support for logging contextual metadata, 

## Quick Start


```go
logger, _ := lad.NewProduction()
defer logger.Sync() // flushes buffer, if any
sugar := logger.Sugar()
sugar.Infow("failed to fetch URL",
  // Structured context as loosely typed key-value pairs.
  "url", url,
  "attempt", 3,
  "backoff", time.Second,
)
sugar.Infof("Failed to fetch URL: %s", url)
```

When performance and type safety are critical, use the `Logger`. It's even
faster than the `SugaredLogger` and allocates far less, but it only supports
structured logging.

```go
logger, _ := lad.NewProduction()
defer logger.Sync()
logger.Info("failed to fetch URL",
  // Structured context as strongly typed Field values.
  lad.String("url", url),
  lad.Int("attempt", 3),
  lad.Duration("backoff", time.Second),
)
```

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
