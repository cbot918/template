
IMAGE=nodejs-web
CONTAINER=nodejs-web

drun:
	docker build -t nodejs-web .
	docker run -it --name nodejs-web -p 3000:3000 nodejs-web

clean:
	docker stop $(CONTAINER)
	docker container rm $(CONTAINER)
	docker image rm $(IMAGE)

.PHONY: drun clean