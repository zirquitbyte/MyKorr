Date: 16.05.2026

# Title
Use message store layer between connectors and database

## Status

accepted

## Context

Unreliable nature of connector options necessitates modularity.

## Decision

Instead of hardcoding connectors, we're going to have a person-in-the-middle type of layer everything is pointing to.


## Consequences

Swapping connectors will be easier. 
