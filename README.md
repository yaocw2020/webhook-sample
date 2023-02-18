# webhook sample
The sample implements a simple validator and mutator for sample.webhook.io/foo.
It also implements a CRD conversion webhook to convert sample.webhook.io/foo between v1 to v2.

### How to deploy
```shell
export KUBECONFIG=<your kubeconfig>
make apply
```