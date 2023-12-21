pullredisimage:
	docker pull redis:latest

createdb:
	docker run -p 6379:6379 redis:latest
