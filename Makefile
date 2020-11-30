.PHONY: build-update-lambda build-local setup-local clean

DEST := bin

build-update-lambda:
	scripts/build-local.sh
	scripts/update-lambda-func-code.sh

build-local:
	scripts/build-local.sh

setup-local:
	scripts/init-custom-localstack.sh

clean:
	rm -f ${DEST}/*
	docker-compose down -v
