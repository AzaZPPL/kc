# kc (Kubernetes Clusters)

## Description
When you're working with different kubernetes clusters it can become a hassle to keep changing the `KUBECONFIG` environment variable when you need to switch to a different cluster.
This tries to make life easier by providing a binary file that shows the clusters available inside of your `$HOME/.kube/kc/` folder. You can choose one of these clusters and it will copy the command to your clipboard.

## Installation
1. Download the binary file available on github
2. Put it in your `$PATH` (e.g. `$HOME/bin` folder)
3. Go to your `.kube` directory and make a new folder called `kc` (e.g. `$HOME/.kube/kc`)
4. Put your cluster config in this new `kc` folder
