module github.com/tnngo/lad/benchmarks

go 1.21

replace github.com/tnngo/lad => ../

require (
	github.com/apex/log v1.9.0
<<<<<<< HEAD
	github.com/go-kit/log v0.2.0
	github.com/rs/zerolog v1.26.0
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/multierr v1.7.0
	github.com/tnngo/lad v1.19.1
	gopkg.in/inconshreveable/log15.v2 v2.0.0-20200109203555-b30bc20e4fd1
=======
	github.com/go-kit/log v0.2.1
	github.com/rs/zerolog v1.30.0
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.23.0
	gopkg.in/inconshreveable/log15.v2 v2.16.0
>>>>>>> 87577d85d58b6d92d0158967df29303d04d30e36
)

require (
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/term v0.12.0 // indirect
)
