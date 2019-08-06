# All services of this project, including gitlab are install on K8S


Before start, assume you already have a gitlab deployment and service, if not, run these command below

```
# deployment
kubectl apply -f deploy/deploy-gitlab.yml

# svc
kubectl apply -f deploy/svc-gitlab.yml
```



## 1. Setup gitlab-ci runner

In order to setup runner, use the token provide on gitlab console(Find token at Setting > CI/CD > Runners)

`curl --request POST "<gitlab_host>/api/v4/runners" --form "token=<xxxxx>"`


You will get some response like below:


`{"id":1,"token":"oooooxxxxx"}`

## 2. Deploy configmap and the runner

```
# configmap 
kubectl apply -f deploy/configmap-runner.yml

# deployment
kubectl apply -f deploy/deploy-runner.yml
```

### 3. Start your CI/CD pipeline


