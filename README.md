# kc (Kubernetes Clusters)

## Description
When you're using a bunch of clusters it can become a hassle to choose the correct and setting it to your kubeconfig.
This tries to make life easier by providing a binary file you can run and it will show the clusters available inside of your `$HOME/.kube/kc/` folder.

## Usage
The variable

## Installation
You can install it in two ways right now.

1. Download the binary file available on github 
2. Download / clone this repo and building it with `go build .`

After that you can use the `kc` executable. I recommend putting it in your `$PATH` (e.g. `$HOME/bin` foler) so you can run `kc` from everywhere.