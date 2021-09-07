.PHONY: app
app:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor github.com/RedAFD/mega/cmd/app

.PHONY: i18n
i18n:
	# Note: gotext has a problem that the -dir parameter does not take effect, 
	# please correct the code and recompile.
	cd ./cmd/app && gotext -dir=../../internal/utils/i18n/locales -srclang=zh-CN update -out=../../internal/utils/i18n/catalog/catalog.go -lang=zh-CN,en

.PHONY: swagger
swagger:
	# go get -u github.com/swaggo/swag/cmd/swag
	swag init -d ./cmd/app -o ./third_party/swagger -parseDependency
