module test-pos

go 1.17

replace github.com/devcode-pos/common => ./common

replace github.com/devcode-pos/databases => ./libs/databases

replace github.com/devcode-pos/models => ./models

replace github.com/devcode-pos/repositories => ./repositories

replace github.com/devcode-pos/controllers => ./controllers

replace github.com/devcode-pos/utils => ./utils

require (
	github.com/devcode-pos/common v0.0.0-00010101000000-000000000000
	github.com/devcode-pos/controllers v0.0.0-00010101000000-000000000000
	github.com/devcode-pos/databases v0.0.0-00010101000000-000000000000
	github.com/devcode-pos/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/devcode-pos/repositories v0.0.0-00010101000000-000000000000 // indirect
	github.com/devcode-pos/utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
