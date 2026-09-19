# Rent Manager — Kubernetes Command Reference

## Tool versions
```powershell
docker --version
kubectl version --client
kind version
```
Checks the installed Docker, kubectl, and kind versions.

## Create the local cluster
```powershell
kind create cluster --name rent-manager
```
Creates a local Kubernetes cluster with kind.

```powershell
kubectl config current-context
```
Shows the active Kubernetes context.

```powershell
kubectl get nodes
```
Shows the cluster nodes and their status.

## Namespace
```powershell
kubectl apply -f platform/kubernetes/namespace.yaml
```
Creates or updates the application namespace.

```powershell
kubectl config set-context --current --namespace=rent-manager
```
Sets `rent-manager` as the default namespace.

## Docker Compose validation
```powershell
docker compose -f platform/docker/docker-compose.yml config
```
Validates and renders the Docker Compose configuration.

## MySQL
```powershell
kubectl get storageclass
```
Lists the available storage classes.

```powershell
kubectl apply -f platform/kubernetes/mysql/secret.yaml
kubectl apply -f platform/kubernetes/mysql/pvc.yaml
kubectl apply -f platform/kubernetes/mysql/statefulset.yaml
kubectl apply -f platform/kubernetes/mysql/service.yaml
```
Creates the MySQL configuration, storage, workload, and Service.

```powershell
kubectl get pvc
```
Shows persistent volume claims and their status.

```powershell
kubectl exec -it mysql-0 -- mysql -u rent_manager -prent_manager rent_manager
```
Opens a MySQL client inside the MySQL Pod.

```powershell
kubectl exec -it mysql-0 -- mysql -u root -proot
```
Opens MySQL as root inside the Pod.

```sql
DROP DATABASE IF EXISTS rent_manager;
CREATE DATABASE rent_manager CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
GRANT ALL PRIVILEGES ON rent_manager.* TO 'rent_manager'@'%';
FLUSH PRIVILEGES;
EXIT;
```
Recreates the local database and restores user permissions.

```powershell
cmd /c "kubectl exec -i mysql-0 -- mysql -u rent_manager -prent_manager rent_manager < backup.sql"
```

Creates backup the local database.

```powershell
cmd /c "kubectl exec mysql-0 -- mysqldump -u rent_manager -prent_manager rent_manager > backup.sql"
```

Imports the SQL backup into Kubernetes MySQL.

```sql
SHOW TABLES;
```
Lists tables in the selected database.

## Backend
```powershell
kubectl apply -f platform/kubernetes/backend/configmap.yaml
kubectl apply -f platform/kubernetes/backend/secret.yaml
kubectl apply -f platform/kubernetes/backend/deployment.yaml
kubectl apply -f platform/kubernetes/backend/service.yaml
```
Creates the backend configuration, workload, and Service.

```powershell
kubectl logs deployment/backend
```
Shows backend application logs.

```powershell
kubectl port-forward service/backend 8082:8080
```
Temporarily exposes the backend on localhost port 8082.

```powershell
curl.exe http://localhost:8082
```
Tests connectivity to the backend.

## Frontend
```powershell
kubectl apply -f platform/kubernetes/frontend/configmap.yaml
kubectl apply -f platform/kubernetes/frontend/deployment.yaml
kubectl apply -f platform/kubernetes/frontend/service.yaml
```
Creates the frontend configuration, workload, and Service.

```powershell
kubectl rollout restart deployment/frontend
kubectl rollout status deployment/frontend
```
Restarts the frontend and waits for the rollout to finish.

```powershell
kubectl port-forward service/frontend 4400:80
```
Temporarily exposes the frontend on localhost port 4400.

## Inspect resources
```powershell
kubectl get pods
kubectl get services
kubectl get deployments
kubectl get replicasets
```
Shows the main Kubernetes resources.

```powershell
kubectl get pods -w
```
Watches Pod changes in real time.

```powershell
kubectl get endpoints backend
kubectl get endpoints mysql
```
Shows the Pods currently reached by each Service.

## Ingress Controller
```powershell
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```
Installs the NGINX Ingress Controller for kind.

```powershell
kubectl get pods -n ingress-nginx
kubectl get services -n ingress-nginx
```
Checks the Ingress Controller resources.

## Application Ingress
```powershell
kubectl apply -f platform/kubernetes/ingress.yaml
```
Creates the application routing rules.

```powershell
kubectl get ingress
kubectl describe ingress rent-manager
```
Shows the Ingress routes and their backends.

```powershell
kubectl port-forward -n ingress-nginx service/ingress-nginx-controller 8085:80
```
Exposes the whole application through Ingress at localhost:8085.

## Self-healing test
```powershell
kubectl get pods -w
```
Watches Kubernetes reconciliation in real time.

```powershell
kubectl delete pod <backend-pod-name>
```
Deletes a backend Pod to test automatic recreation.

## Local access

Direct port forwarding:
```text
Frontend: http://localhost:4400
Backend:  http://localhost:8082
```

Ingress:
```text
Application: http://localhost:8085
Frontend route: /
Backend API route: /v1
```

With Ingress, normal local use requires only the Ingress Controller port-forward.
```cmd
kubectl port-forward -n ingress-nginx service/ingress-nginx-controller 8085:80
```