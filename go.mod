module github.com/jasondavindev/hacktoberfest-2020

go 1.15

replace (
<<<<<<< HEAD
	github.com/jasondavindev/hacktoberfest-2020/config => ./config
	github.com/jasondavindev/hacktoberfest-2020/domain => ./domain
=======
  github.com/jasondavindev/hacktoberfest-2020/command => ./command
  github.com/jasondavindev/hacktoberfest-2020/config => ./config
  github.com/jasondavindev/hacktoberfest-2020/listener => ./listener
>>>>>>> 63170ad0a90bc21fe56ec45e9d75533ca43f68f3
)

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/joho/godotenv v1.3.0
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	go.uber.org/config v1.4.0
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20201001193750-eb9a90e9f9cb // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200930145003-4acb6c075d10 // indirect
	golang.org/x/sync v0.0.0-20200930132711-30421366ff76 // indirect
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20201001230009-b5b87423c93b // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/fsnotify.v1 v1.4.7
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
