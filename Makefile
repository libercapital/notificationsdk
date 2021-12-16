unit-test:
	gotestsum -- ./... -failfast -race -coverprofile ./coverage.out
