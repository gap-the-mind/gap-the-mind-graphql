module github.com/gap-the-mind/gap-the-mind-graphql

go 1.14

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/gap-the-mind/gap-the-mind-storage v0.0.0-20200717070746-d94f4a6bcbbf
	github.com/vektah/gqlparser/v2 v2.0.1
	go.uber.org/zap v1.15.0
	github.com/gap-the-mind/gap-the-mind-storage v1.0.0
)

replace github.com/go-git/go-billy/v5 => ../go-billy

replace github.com/go-git/go-git/v5 => ../go-git

replace github.com/gap-the-mind/gap-the-mind-storage => ../gap-the-mind-storage
