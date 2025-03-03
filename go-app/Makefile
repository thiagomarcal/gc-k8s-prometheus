build:
	eval $(minikube docker-env)
	docker build -t go-app:latest .
	minikube image load go-app:latest

deploy:
	kubectl delete -f deployment.yaml
	kubectl apply -f deployment.yaml
	kubectl rollout restart deployment go-app
