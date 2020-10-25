#!/bin/sh
kubectl create secret generic aws-credentials --from-file=credentials=$HOME/.aws/credentials --from-file=config=$HOME/.aws/config
