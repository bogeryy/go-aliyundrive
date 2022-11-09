local_package=github.com/bogeryy/go-aliyundrive
all: lint

lint:
	@$(set_env) go fmt ./...
	@$(set_env) goimports -local $(local_package) -w .