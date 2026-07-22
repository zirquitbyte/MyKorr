## RhyCell Device Firmware Architecture

Firmware architecture for portable esp32 communicators.

---

### System Overview

``` mermaid
flowchart TD
  A["User Input </br> (chorded keyboard, rotary encoder, buttons, etc.)"] --> B["Input Manager </br> decode key chords, normalize output"]
  B --> C["UI Manager </br> map input to UI actions (left, right, enter, back, etc.)"]
  C --> D["Application Runtime </br> active app routing, lifecycle (start, suspend, kill etc.)"]
  D --> E["Application </br> chat, settings, weather, ebook reader, etc."]
  E --> F["Application SDK </br> API for apps, messaging, storage, notifications, UI, settings, etc."]
  F --> G["Services </br> Messaging, Storage, Notifications, Settings,Scheduler, Sync, Search/Request"]
  G --> H["Platform Layer </br> Network, Display, Audio, Power, Crypto, Time, Filesystem"]
  H --> I["Hardware Abstraction Layer (HAL) </br> call display/audio/etc. without pin-specific code Display driver, Keyboard driver, Battery driver, SPI, I2C, GPIO"]
  I --> J["ESP-IDF / FreeRTOS"]
  J --> K["ESP32 Hardware"]

```

---


### Data flow

``` mermaid
flowchart TD
  A["Application (chat)"] --> B["Application SDK"]
  B --> C["Messaging Service"]
  C --> D["Sync Service"]
  D --> E["Protocol Layer"]
  E --> F["Network Service"]
  F --> G["Wireguard Tunnel"]
  G --> H["Backend API"]
  H --> I["Backend Services"]

```



### Event bus

Services to be communicating with each other through event bus
