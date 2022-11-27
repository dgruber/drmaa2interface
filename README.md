# drmaa2interface

DRMAA2 compatible native Go interfaces and structs for building DRMAA2 compatible middleware

## Why using drmaa2interface?

This repository simplifies the process to create Go DRMAA2 wrappers for job schedulers,
process managers, resource management systems etc. (like for starting up OS processes,
workflows, containers, pods...).

## What is DRMAA2?

DRMAA2 is an acronym for [Distributed Resource Management Application API version 2](http://www.ogf.org/documents/GFD.194.pdf) which
is an open and freely usable standard defined by the [Open Grid Forum](http://www.ogf.org).

Unlike other standards it is a common subset of functions available in all major DRMs
(like Univa Grid Engine, LSF, SLURM, Condor, PBS).

## More information

More information can be found at [the DRMAA website](http://www.drmaa.org).

Please feel free to create issues on github.

## Projects

The DRMAA2 interface is used in following projects:

- [wfl - a simple workflow language for Go](https://github.com/dgruber/wfl)
- [drmaa2os - An implementation of the DRMAA2 interface for OS processes, container runtimes, and more](https://github.com/dgruber/drmaa2os)
