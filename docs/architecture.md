## MyKorr

A self-hosted information relay for portable esp32 communicators.

**Status:** Prototype / WIP / Proof-of-concept

---

### System overview

+------------------------------------------------------------------+
| External Platforms                                               |
| Telegram | Signal | Discord | WhatsApp | RSS / Email             |
+------------------------------------------------------------------+
                               |
                               v

+------------------------------------------------------------------+
| Connector Runtime                                                |
| Headless clients                                                 |
+------------------------------------------------------------------+
                               |
                               v

+------------------------------------------------------------------+
| Message Normalization                                            |
| - Unified schema                                                 |
| - Necessary metadata                                             |
| - Deduplication                                                  |
+------------------------------------------------------------------+
                               |
                               v

+------------------------------------------------------------------+
| Storage Layer                                                    |
| - Encrypted message DB                                           |
| - Search index                                                   |
| - RSS cache                                                      |
| - Media storage                                                  |
+------------------------------------------------------------------+
            |                            |                     |
            v                            v                     v

+-------------------+  +--------------------+  +-------------------+
| Query Engine      |  | Messaging API      |  | File Sync         |
| - RSS fetch       |  | - Sync             |  | - Fetch mp3       |
| - Article fetch   |  | - Messaging        |  | - Fetch ebook     |
| - Adblock         |  | - Auth             |  | - From storage    |
+-------------------+  +--------------------+  +-------------------+
            \                            |                     /
             \                           |                    /
              \                          |                   /
               \                         |                  /
                \                        |                 /
                 v                       v                v

+------------------------------------------------------------------+
| Secure Overlay Network                                           |
| - WireGuard                                                      |
| - Pairing                                                        |
| - Relay server for CGNAT                                         |
+------------------------------------------------------------------+
                               |
                               v

+------------------------------------------------------------------+
| ESP32 Communicator                                               |
| - E-ink UI                                                       |
| - Chorded input                                                  |
| - Navigation                                                     |
| - Offline cache                                                  |
| - Media playback                                                 |
+------------------------------------------------------------------+


---

### Backend Architecture Diagram

flowchart TB

+--------------------------------------------------------------+
| API Gateway                                                  |
| HTTP | Auth | Session | Device routing                       |
+--------------------------------------------------------------+
             |                    |                    |
             v                    v                    v

+----------------------+  +----------------------+  +----------------------+
| Message API          |  | Query API            |  | Device API           |
| - Inbox              |  | - RSS                |  | - Sync               |
| - Send               |  | - Article fetch      |  | - Presence           |
| - Conversations      |  |   by query           |  |                      |
+----------------------+  +----------------------+  +----------------------+
             \                    |                    /
              \                   |                   /
               \                  |                  /
                \                 |                 /
                 v                v                v

+--------------------------------------------------------------+
| Internal Event Bus                                           |
| - Message events                                             |
| - Connector events                                           |
| - Sync events                                                |
| - Queue / retry                                              |
+--------------------------------------------------------------+
             |                    |                    |
             v                    v                    v

+----------------------+  +----------------------+  +----------------------+
| Connector Runtime    |  | Message Store        |  | Query Engine         |
| - Telegram           |  | - SQLite             |  | - RSS parser         |
| - Signal             |  | - Encryption         |  | - Article fetch      |
| - Discord            |  | - Metadata index     |  | - Adblock            |
| - WhatsApp           |  | - Attachments        |  |                      |
+----------------------+  +----------------------+  +----------------------+
             |
             v

+--------------------------------------------------------------+
| Connector Isolation Layer                                    |
| - Per-connector state                                        |
| - Health checks                                              |
+--------------------------------------------------------------+

---

### ESP32 Frontend Architecture Diagram

flowchart TB

+--------------------------------------------------------------+
| User Interaction                                             |
| Chorded keyboard | Rotary encoder | Buttons                  |
+--------------------------------------------------------------+
                              |
                              v

+--------------------------------------------------------------+
| Input Processing                                              |
| - Key mapping                                                 |
| - Navigation state                                            |
| - Shortcuts                                                   |
+--------------------------------------------------------------+
                              |
                              v

+--------------------------------------------------------------+
| UI State Engine                                               |
| - Conversations                                               |
| - Message rendering                                           |
| - Articles                                                    |
| - Media                                                       |
| - Notifications                                               |
+--------------------------------------------------------------+
                              |
                              v

+--------------------------------------------------------------+
| Network Sync Client                                           |
| - API requests                                                |
| - Auth                                                        |
| - Cache sync                                                  |
| - Overlay VPN tunnel                                          |
+--------------------------------------------------------------+
                              |
                              v

+--------------------------------------------------------------+
| Local Device Cache                                            |
| - Messages                                                    |
| - UI state                                                    |
| - Offline queue                                               |
| - Cached articles                                             |
+--------------------------------------------------------------+

---

### Trust Boundary / Encryption Diagram

flowchart TB

+--------------------------------------------------------------+
| External Platforms                                           |
+--------------------------------------------------------------+
                              |
                              v

+--------------------------------------------------------------+
| Connector Layer                                              |
+--------------------------------------------------------------+
                              |
                              v

+--------------------------------------------------------------+
| Normalization / Processing                                   |
| temporary plaintext                                          |
+--------------------------------------------------------------+
                              |
                              v
                       (Encryption Boundary)

+--------------------------------------------------------------+
| Encrypted Storage                                            |
| - Messages                                                   |
| - Attachments                                                |
| - User content                                               |
+--------------------------------------------------------------+
                              |
                              v

+--------------------------------------------------------------+
| Messaging API                                                |
+--------------------------------------------------------------+
                              |
                              v

+--------------------------------------------------------------+
| ESP32 Device                                                 |
| Trusted endpoint | Final rendering                           |
+--------------------------------------------------------------+

---

### Goals

## MVP (Current Goal)

• Go server daemon
• Fake connector
• Telegram connector
• Message normalization
• SQLite DB
• API
• Basic ESP32 text interface


## V1

• Multiple connectors
• Encrypted storage
• Overlay VPN pairing
• RSS/article system
• Ebook/music support



## Longterm goals

• Applet ecosystem
• Voice note transcription
