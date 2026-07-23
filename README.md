## MyKorr

A self-hosted information relay for portable esp32 communicators.

> .[!Warning]
> This repository has been archived. Please view the newer version at https://codeberg.org/zirquitbyte/MyKorr  

---

### Quick start

1. Clone
```bash
git clone https://github.com/zirquitbyte/MyKorr.git
cd MyKorr
```

2. Install dependencies
```bash
go mod tidy
```

3. Run
```bash
go run ./cmd/server/main.go
```

Server starts on:
    http://localhost:5656

---

### API Endpoints

#### GET /messages

Retrieve messages. 

Example:

```bash
curl "http://localhost:5656/messages?conversation_id=1"
```

#### POST /messages

Send messages. 

Example:

```bash
❯ curl -X POST http://localhost:5656/messages \          
  -H "Content-Type: application/json" \
  -d '{"conversation_id":"1","sender":"me","body":"hello world"}'
```

---

### License

This project is licensed under Apache-2.0 license.

---

### Contributing

No external contributions accepted at this time. Working on our contributions policy.

---

### MVP Roadmap

• Basic firmware
• First messaging connectors


## V1

• Multiple connectors
• Encrypted storage
• Overlay VPN pairing
• RSS/article system
• Ebook/music support


## Longterm goals

• Applet ecosystem
• Third-party applications
• Audio
• GPS
• Mesh networking

---

### Contact

contact@zirquitbyte.dev

