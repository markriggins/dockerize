.PHONY : build run-staging run-prod

build: 
	cd ../..; goxc -os linux -arch amd64
	cp -f $(shell bash -xc 'find $(GOPATH)/bin -name dockerize_linux_amd64.tar.gz| head -1') .dockerize_linux_amd64.tar.gz
	docker build -t dockerized-nginx .
	rm .dockerize_linux_amd64.tar.gz


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

run-prod-secrets:
	docker rm -f dockerized-nginx 2>/dev/null || true
	docker run -d -p 80:80 -e DEPLOYMENT_ENV=prod --name dockerized-nginx \
				-v $(PWD):/secrets \
				dockerized-nginx \
					-overlay /tmp/overlays/_common/html:/usr/share/nginx/ \
					-overlay "/tmp/overlays/{{ .Env.DEPLOYMENT_ENV }}/html:/usr/share/nginx/" \
					-template "/secrets/secrets.html.tmpl:/usr/share/nginx/html/secrets.html" \
			        -secrets /secrets/secrets.env -- nginx 
	open http://$(shell docker-machine ip $(shell docker-machine active))/secrets.html
