build-consumer:
	docker build --build-arg app=consumer -t go-with-nats:consumer .
	
build-producer:
	docker build --build-arg app=producer -t go-with-nats:producer .