.PHONY : build run-staging run-prod stop run-proxy-amazon run-prod-secrets
.PHONY : zombie-maker-debian-binary

build: zombie-maker
	cd ../..; goxc -os linux -arch amd64
	cp -f $(shell bash -xc 'find $(GOPATH)/bin -name dockerize_linux_amd64.tar.gz| head -1') .dockerize_linux_amd64.tar.gz
	docker build -t dockerized-nginx .
	rm .dockerize_linux_amd64.tar.gz


zombie-maker: zombie-maker.c
	# build a native and debian binary
	docker run --rm -v "$(PWD)":/tmp -w /tmp gcc gcc -o zombie-maker-debian-binary zombie-maker.c
	gcc -o zombie-maker zombie-maker.c


stop:
	docker rm -f dockerized-nginx 2>/dev/null || true

run-staging: stop
	docker run -d -p 80:80 -e DEPLOYMENT_ENV=staging --name dockerized-nginx dockerized-nginx \
		     -overlay /tmp/overlays/_common/html:/usr/share/nginx/ \
	         -overlay "/tmp/overlays/{{ .Env.DEPLOYMENT_ENV }}/html:/usr/share/nginx/" \
	         -- nginx
	open http://$(shell docker-machine ip $(shell docker-machine active))/robots.txt
	open http://$(shell docker-machine ip $(shell docker-machine active))

run-prod: stop
	docker run -d -p 80:80 -e DEPLOYMENT_ENV=prod --name dockerized-nginx dockerized-nginx \
		     -overlay /tmp/overlays/_common/html:/usr/share/nginx/ \
	         -overlay "/tmp/overlays/{{ .Env.DEPLOYMENT_ENV }}/html:/usr/share/nginx/" \
	         -- nginx
	open http://$(shell docker-machine ip $(shell docker-machine active))/robots.txt

run-proxy-amazon: stop
	docker run -d -p 80:80 -e PROXY_PASS_URL="https://www.amazon.com/" --name dockerized-nginx dockerized-nginx
	open http://$(shell docker-machine ip $(shell docker-machine active))

run-prod-secrets: stop
	docker run -d -p 80:80 -e DEPLOYMENT_ENV=prod --name dockerized-nginx \
				-v $(PWD):/secrets \
				dockerized-nginx \
					-overlay /tmp/overlays/_common/html:/usr/share/nginx/ \
					-overlay "/tmp/overlays/{{ .Env.DEPLOYMENT_ENV }}/html:/usr/share/nginx/" \
					-template "/secrets/secrets.html.tmpl:/usr/share/nginx/html/secrets.html" \
			        -secrets /secrets/secrets.env -- nginx 
	open http://$(shell docker-machine ip $(shell docker-machine active))/secrets.html

run-zombie-test: stop
	docker run -d --name dockerized-nginx dockerized-nginx -reap -- zombie-maker 10 200 
	@echo "Tailing logs -- look for zombies"
	docker logs -f dockerized-nginx


