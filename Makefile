IMAGE_NAME=exchangemonitor
build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -it -p 8080:8080 $(IMAGE_NAME)

remove:
	docker stop $(IMAGE_NAME) && docker rm $(IMAGE_NAME)

clean:
	docker rmi $(IMAGE_NAME)