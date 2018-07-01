client-build:
	docker build -t retro-memo-client .

client-run:
	docker run -it -v $(shell pwd)/src:/usr/src/app/src -v $(shell pwd)/public:/usr/src/app/public -p 8080:8080 retro-memo-client

client-run-osx:
	docker run -it -v $(shell pwd)/src:/usr/src/app/src:cached -v $(shell pwd)/public:/usr/src/app/public:cached -p 8080:8080 retro-memo-client

up:	client-build client-run

up-osx:	client-build client-run-osx
