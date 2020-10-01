module github.com/jasondavindev/hacktoberfest-2020

go 1.15

replace (
  github.com/jasondavindev/hacktoberfest-2020/domain => ./domain
)

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/joho/godotenv v1.3.0
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
)
