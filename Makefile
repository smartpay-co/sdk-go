.PHONY: integration-test GUARD

integration-test: guard-SMARTPAY_SECRET_KEY guard-SMARTPAY_PUBLIC_KEY guard-API_BASE
	go test tests/integration/* -v


guard-%: GUARD
	@ if [ -z '${${*}}' ]; then echo 'Environment variable $* not set.' && exit 1; fi

GUARD:
