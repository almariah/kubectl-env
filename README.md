# kubectl-env

The `kubectl-env` kubectl plugin will allow templating kustomize manifests using environment variables.

## Usage

To generate kubernetes manifest from kustomize manifests using `kubectl-env`, create you own kustomize manifests with the proper templating variables. The default delimiter for template variables is `{{ .SOME_VAR }}`. Then run `kubectl` with `env` instead of `kustomize` and the needed environment variables:

```bash
SOME_VAR=some_value kubectl env /some/kustomize/manifests/
```
or applying using:

```bash
SOME_VAR=some_value kubectl env /some/kustomize/manifests/ | kubectl apply -f-
```

You can override the default delimiter using `--left-delimiter` and `--right-delimiter`. For example if the template variable is `{{{ .SOME_VAR }}}`, then you can use `kubectl-env` as follows:

```bash
SOME_VAR=some_value kubectl env /some/kustomize/manifests/ --left-delimiter="{{{" --right-delimiter="}}}"
```

## Installation

To install `kubectl-env`:

```bash
kubectl krew install env
```

## kubeclt version:

Versions of `kubectl-env` and compatibility with `kubectl`:

| kubectl-env | kubectl |
|---|---|
| v1.0.0 | v0.19.2 |
