Serverless-Operator
===================

Bringing [GitOps](https://www.weave.works/blog/what-is-gitops-really) to Serverless.

## Description

The purpose of the Serverless-Operator project is to enable GitOps for serverless applications. A Git repo should represent the desired state of any number of serverless solutions and this state will be reconciled with the current state within a cloud provider. 

The project has been started due to a gap in the serverless CICD experience, especially for scenarios where:
- multiple environments are needed
- environment promotion is needed
- large and complex serverless solutions

Initially we will be targetting the [Serverless Framework](https://serverless.com/) and its envisioned that the project will be self hosted. However, in the future its envisioned that there may be a SaaS offering offering a free and paid for tiers for people that don't want to run their own infrastructure.

Its influenced by [OpenFaas Cloud](https://github.com/openfaas/openfaas-cloud).

## Contributing

As this is a new project we are looking for people to help out with the project. This can be in the form of development, feature requests, testing and in numerous other ways.

## Expected Features

- [ ] Kubernetes Operator that will understand custom resources of kind 'ServerlessApplicationRelease'.  The operator will deploy, update, remove a Serverless Framework solution based on this. [Flux](https://github.com/weaveworks/flux) from Weaveworks can be used to monitor a Git repo for these resources.
- [ ] Serverless Framework packager that will monitor a source repo and automatically package a serverless application when there is a change. This will integrate with the GitHub Status API.
- [ ] Self-hosted solution that can be deployed to Kubernetes. With scripts for EKS that uses `eksctl`
- [ ] SaaS Edition that people can use if they don't want to run Kubernetes.
- [ ] A large sample serverless framework application that cen be used for demos.
- [ ] Support secrets via Bitnami SealedSecrets
