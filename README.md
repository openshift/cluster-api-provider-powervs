# OpenShift cluster-api-provider-powervs

This repository hosts an implementation of a provider for PowerVS for the
OpenShift [machine-api](https://github.com/openshift/cluster-api).

This provider runs as a machine-controller deployed by the
[machine-api-operator](https://github.com/openshift/machine-api-operator)

### How to build the images in the RH infrastructure
The Dockerfiles use `as builder` in the `FROM` instruction which is not currently supported
by the RH's docker fork (see [https://github.com/kubernetes-sigs/kubebuilder/issues/268](https://github.com/kubernetes-sigs/kubebuilder/issues/268)).
One needs to run the `imagebuilder` command instead of the `docker build`.

Note: this info is RH only, it needs to be backported every time the `README.md` is synced with the upstream one.

## Deploy machine API plane with kubernetes cluster

1. **Deploying the cluster**

    Use any existing mechanism for deploying the kubernetes cluster, e.g: kubeadm https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/

2. **Deploying machine API controllers**

    For development purposes the powervs machine controller itself will run out of the machine API stack.
    Otherwise, docker images needs to be built, pushed into a docker registry and deployed within the stack.

    To deploy the stack:
    ```
    kustomize build config | kubectl apply -f -
    ```

3. **Deploy secret with Power VS credentials**

   Power VS nodeupdate controller assumes existence of a secret file:

   ```yaml
   apiVersion: v1
   kind: Secret
   metadata:
     name: powervs-credentials-secret
     namespace: openshift-machine-api
   type: Opaque
   data:
     IBMCLOUD_API_KEY: FILLIN
   ```

   You can use `examples/render-powervs-secrets.sh` script to generate the secret:
   ```sh
   # Set the IBMCLOUD_API_KEY with a proper IBM Cloud API Key
   $ IBMCLOUD_API_KEY=<API_KEY> ./examples/render-powervs-secrets.sh examples/addons.yaml | kubectl apply -f -
   ```

   Go to [How to create IBM Cloud API Key](#How-to-create-IBM-Cloud-API-Key) for creating API Key

4. **Test by creating example machine**

   ```shell
   # Update the relevant fields like serviceInstanceID, imageID, subnets, keyName etc.. 
   $ kubectl create -f examples/machine-with-user-data.yaml
   $ kubectl create -f examples/userdata.yml
   ```
## Test locally built aws actuator

1. **Tear down machine-controller**

   Deployed machine API plane (`machine-api-controllers` deployment) is (among other
   controllers) running `machine-controller`. In order to run locally built one,
   simply edit `machine-api-controllers` deployment and remove `machine-controller` container from it.

1. **Build and run aws actuator outside of the cluster**

   ```sh
   $ go build -o bin/machine-controller-manager github.com/openshift/cluster-api-provider-powervs/cmd/manager
   ```

   ```sh
   $ .bin/machine-controller-manager --kubeconfig ~/.kube/config --logtostderr -v 5 -alsologtostderr
   ```
      If running in container with `podman`, or locally without `docker` installed, and encountering issues, see [hacking-guide](https://github.com/openshift/machine-api-operator/blob/master/docs/dev/hacking-guide.md#troubleshooting-make-targets).


1. **Deploy k8s apiserver through machine manifest**:

   To deploy user data secret with kubernetes apiserver initialization (under [config/master-user-data-secret.yaml](config/master-user-data-secret.yaml)):

   ```yaml
   $ kubectl apply -f config/master-user-data-secret.yaml
   ```

   To deploy kubernetes master machine (under [config/master-machine.yaml](config/master-machine.yaml)):

   ```yaml
   $ kubectl apply -f config/master-machine.yaml
   ```

1. **Pull kubeconfig from created master machine**

   The master public IP can be accessed from AWS Portal. Once done, you
   can collect the kube config by running:

   ```
   $ ssh -i SSHPMKEY ec2-user@PUBLICIP 'sudo cat /root/.kube/config' > kubeconfig
   $ kubectl --kubeconfig=kubeconfig config set-cluster kubernetes --server=https://PUBLICIP:8443
   ```

   Once done, you can access the cluster via `kubectl`. E.g.

   ```sh
   $ kubectl --kubeconfig=kubeconfig get nodes
   ```

## Deploy k8s cluster in AWS with machine API plane deployed

1. **Generate bootstrap user data**

   To generate bootstrap script for machine api plane, simply run:

   ```sh
   $ ./config/generate-bootstrap.sh
   ```

   The script requires `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables to be set.
   It generates `config/bootstrap.yaml` secret for master machine
   under `config/master-machine.yaml`.

   The generated bootstrap secret contains user data responsible for:
   - deployment of kube-apiserver
   - deployment of machine API plane with aws machine controllers
   - generating worker machine user data script secret deploying a node
   - deployment of worker machineset

1. **Deploy machine API plane through machine manifest**:

   First, deploy generated bootstrap secret:

   ```yaml
   $ kubectl apply -f config/bootstrap.yaml
   ```

   Then, deploy master machine (under [config/master-machine.yaml](config/master-machine.yaml)):

   ```yaml
   $ kubectl apply -f config/master-machine.yaml
   ```

1. **Pull kubeconfig from created master machine**

   The master public IP can be accessed from AWS Portal. Once done, you
   can collect the kube config by running:

   ```
   $ ssh -i SSHPMKEY ec2-user@PUBLICIP 'sudo cat /root/.kube/config' > kubeconfig
   $ kubectl --kubeconfig=kubeconfig config set-cluster kubernetes --server=https://PUBLICIP:8443
   ```

   Once done, you can access the cluster via `kubectl`. E.g.

   ```sh
   $ kubectl --kubeconfig=kubeconfig get nodes
   ```

## How to create IBM Cloud API Key

### Method 1: Create an API Key via IAM
https://cloud.ibm.com/docs/account?topic=account-userapikey&locale=en#create_user_key

### Method 2: Create an API Key via Service Account
<!TODO: Add information>

# Upstream Implementation
Other branches of this repository may choose to track the upstream
Kubernetes [Cluster-API AWS provider](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)

In the future, we may align the master branch with the upstream project as it
stabilizes within the community.
