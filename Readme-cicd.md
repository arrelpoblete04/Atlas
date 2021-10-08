

tkn pipeline delete golang-pipeline --force
kubectl apply -f golang-pipeline.yaml

tkn taskrun delete golang-taskrun --force
kubectl apply -f golang-tekton-taskrun.yaml

tkn task delete golang-task-one-go-build --force
kubectl apply -f task-01-go-build.yaml

tkn task delete golang-task-two-gosec --force
kubectl apply -f task-02-go-security.yaml

tkn task delete golang-task-three-go-build --force
kubectl apply -f task-03-go-build.yaml

tkn taskrun delete golang-taskrun-one-go-build --force
kubectl apply -f taskrun-01-go-build.yaml
tkn taskrun logs golang-taskrun-go-build -f

tkn task delete golang-task-two-docker-build-push --force
kubectl apply -f task-02-docker.yaml
tkn taskrun delete taskrun-build-and-push --force
kubectl apply -f taskrun-02-docker.yaml
tkn taskrun logs -f taskrun-build-and-push


-------------------------------------------------------

kubectl apply -f golang-secret.yaml
kubectl apply -f golang-service-account.yaml
kubectl apply -f golang-tekton-pipeline-resource.yaml

tkn task delete golang-task-one-go-build --force
kubectl apply -f task-01-go-build.yaml

tkn task delete golang-task-two-docker-build-push --force
kubectl apply -f task-02-docker.yaml

tkn resource delete golang-pipeline-resource --force
kubectl apply -f golang-tekton-pipeline-resource.yaml

tkn pipeline delete golang-pipeline --force
kubectl apply -f golang-pipeline.yaml

tkn pipelinerun delete golang-build-pipeline --force
kubectl apply -f golang-pipelinerun.yaml
tkn pipelinerun logs -f golang-build-pipeline



------

kubectl apply -f triggers-01-rbac-role.yaml
kubectl apply -f triggers-02-Eventlistener.yaml
kubectl apply -f triggers-03-secret.yaml
kubectl apply -f triggers-04-triggerbinding.yaml
kubectl apply -f triggers-05-triggertemplate.yaml
kubectl apply -f triggers-06-ingress.yaml




