[![Build Status](https://travis-ci.org/moshloop/configadm.svg?branch=master)](https://travis-ci.org/moshloop/configadm)
[![codecov](https://codecov.io/gh/moshloop/configadm/branch/master/graph/badge.svg)](https://codecov.io/gh/moshloop/configadm)
[![Go Report Card](https://goreportcard.com/badge/github.com/moshloop/configadm)](https://goreportcard.com/report/github.com/moshloop/configadm)

configadm is a node instance configuration tool focused on bootstrapping nodes for container based environments


# Design
`configadm`
![](./docs/configadm.png)

## Phases

`configadm` using a chain of phases. phases earlier in the chain can update items later in the chain. e.g. The `CRI` and `Kubernetes` phases can add packages to be installed in the `packages` phase.

## Virtual Filesystem

Phases do not write directly to the filesystem, only to the virtual filesystem. Phases can also read and update files written by earlier phases.

## Commands

Commands are executed at 3 specific points:

**Pre Commands**
Pre-commands are used to prepare the environment for execution, OS detection and setting of runtime flags is done in this phase so that they can be used in all other phases. e.g. set an environment variable based on the output of a command.

**Commands**
Phases can only append to this Commands list.

**Post Commands**
Post commands run after all the phases have completed and can be used for cleanup functions are for handing off to other systems.

## Apply

One all phases have run the virtual filesystem and command list is used to either create a cloud-init file or ISO. Or applied directly to the system using a shell.

# Design Principles

## Cloud Native

`configadm` has native support for containerized environments and *nothing* else - It is designed to bootstrap immutable container hosts and then handoff to the container(s) and/or orchestrators for everything else.

* systemd
* containers
* CRI
* Kubernetes

## Dependency Free

**Runtime Dependencies**
> `configadm` is built in pure Go and distributed as a statically linked binary.

Ansible is a good example of how bad dynamically linked tools can get - The dependencies for core Ansible is relatively easy to solve using a Python virtualenv, the 1000's of modules that makeup the extended ansible platform depend on hundreds of other packages - These dependencies are at best visible in documentation and error messages, there is no way of knowing what dependencies (nevermind the version) that are required for a given playbook.

**Explicit Execution Dependencies**
> `configadm` does not support any control flows (besides for runtime tags), the ordering of actions is well defined and cannot be changed.

Ansible provides explicit ordering using control statements such as loops and conditionals,  this explicit ordering is coupled with implicit variable management with a very complex precedence hierarchy and ruleset. Creating a mental model of what a given variable will be is almost impossible. Unit testing this state is impossible.

**Implicit Execution Dependencies**
Unlike Ansible, CFEngine and Terraform do not have explicit ordering, Some might argue that they don't have ordering altogether - However the use of input variables and classes create implicit ordering with a complex runtime state machine - It is almost impossible to create and execute this state machine mentally making it very difficult to troubleshoot and test.

## Stateless
> `configadm` uses a virtual filesystem and command set for all higher order functions - this makes it trivial to compose and unit functionality.

While the execution model of Ansible does not have persistent state, it is state driven. Facts discovered at runtime can alter behavior, when used across a cluster of machine the impact of state becomes even more important with intermittent connection issues to cluster members potentially creating conflicting state between runs.
The use of change tracking also makes it impossible to ascertain whether a given step will execute.

Ansible, CFEngine and Terraform all execute on the underlying components directly, making only integration testing possible (and difficult)



