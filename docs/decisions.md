# Decisions


## Go as primary backend language

Golang is our choice for server runtime.
Faster to dev in compared to other options considered (e.g. Rust), good for concurrency.


## API layer

Communication between end device and backend to be implemented using an API layer.
Helps avoid tight coupling, keeps possibility for multiple client types open for the future, simplifies sync.


## Message store

Communication between messaging platforms and our system to happen through a message store layer.
Reasoning: allows for modular approach while avoiding tight coupling. Messaging connectors become easier to replace (and we expect to need that, considering how frequently that landscape changes).


## Connectors

We reuse existing connectors where possible, otherwise relying on headless clients.


## Message normalization

All connector output is normalized into an internal message format to keep storage and UI platform-agnostic.
Will make querying and sync easier.


## SQLite as message database

SQLite to be used as database. Lightweight, reliable, suitable for SBCs.


## WireGuard tunneling

Connection of device to backend through WireGuard.
No-brainer.


## Networking relay

In case of CGNAT on user's side, a relay VPS will be used.


## Encryption boundary at storage layer

Messages temporarily exist in RAM as plaintext during connector processing but are encrypted right away before going to persistent storage.
Due to the nature of connecting external messaging platforms, we can't avoid decrypting messages for processing. But we can design our side with encryption in mind.


## Modular subsystem architecture

Core systems are kept loosely coupled (e.g. message store layer between connectors and database) to prevent chaos as features expand and when necessary connector maintenance (swapping one implementation for another) comes. Simplifies iteration.


## Self-hosted by design

The system is designed primarily around self-hosted SBC servers (with possibility of SaaS model in mind).
Self-hosted SBC plug-and-play approach both aligns well with our philosophy, while also having the added benefit of keeping the short window messages exist as plaintext during processing on home turf.
This comes with a constraint of limited computing power SBCs can provide, which does not constitute a particular issue for us because we aim for low power, low bandwidth kind of tech.
