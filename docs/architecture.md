## MyKorr

A self-hosted information relay for portable esp32 communicators.

---

### System overview
``` mermaid
flowchart TB

A[External Platforms<br/>Telegram, Signal, Discord, WhatsApp, RSS/Email] --> B[Connector Runtime<br/>Headless clients]

B --> C[Message Normalization<br/>Unified schema, Necessary metadata, Deduplication]

C --> D[Storage Layer<br/>Encrypted message DB, Search index, RSS cache, Media storage]

D --> E1[Query Engine<br/>RSS fetch, article fetch, adblock]
D --> E2[Messaging API<br/>Sync, Messaging, Auth]
D --> E3[File Sync<br/> fetch mp3, ebook from storage]

E1 --> F[Secure Overlay Network<br/>WireGuard, Pairing, relay server for CGNAT]
E2 --> F
E3 --> F

F --> G[ESP32 Communicator<br/>E-ink UI, Chorded input, Navigation, Offline cache, Media playback]
```

---

### Backend Architecture Diagram
``` mermaid
flowchart TB

A[API Gateway<br/>HTTP,  Auth, Session, Device routing]

A --> B1[Message API<br/>Inbox, Send, Conversations]
A --> B2[Query API<br/>RSS, Article fetching by query]
A --> B3[Device API<br/>Sync, Presence]

B1 --> C[Internal Event Bus<br/>Message, Connector, Sync events, Queue/retry]
B2 --> C
B3 --> C

C --> D1[Connector Runtime<br/>Telegram, Signal, Discord, WhatsApp]
C --> D2[Message Store<br/>SQLite, Encryption, Metadata index, Attachments]
C --> D3[Query Engine<br/>RSS parser, Article fetch, Adblock]

D1 --> E[Connector Isolation Layer<br/>Per-connector state, Health checks]
```
---

### ESP32 Frontend Architecture Diagram
``` mermaid
flowchart TB

A[User Interaction<br/>Chorded keyboard, Rotary encoder, Buttons]

A --> B[Input Processing<br/>Key mapping, Navigation state, Shortcuts]

B --> C[UI State Engine<br/>Conversations, Message rendering, Articles, Media, Notifications]

C --> D[Network Sync Client<br/>API requests, Auth, Cache sync, Overlay VPN tunnel]

D --> E[Local Device Cache<br/>Messages, UI state, Offline queue, Cached articles]
```
---

### Trust Boundary / Encryption Diagram
``` mermaid
flowchart TB

A[External Platforms] --> B[Connector Layer]

B --> C[Normalization / Processing<br/>temporary plaintext]

C -->|Encryption Boundary| D[Encrypted Storage<br/>Messages, Attachments, User content]

D --> E[Messaging API]

E --> F[ESP32 Device<br/>Trusted endpoint, Final rendering]
```
---
