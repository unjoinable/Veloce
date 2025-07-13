# GEMINI.md

## 🧩 Project Overview

This project is a custom Minecraft-compatible **protocol server written in Go**.  
It focuses **only on the protocol and networking layer** — not world simulation, rendering, or game logic.

Its main purpose is to:
- Accept TCP connections from Minecraft clients
- Handle handshake, login, ping/status, and play state
- Parse and encode all relevant Minecraft protocol packets (vanilla or custom)
- Maintain clean session state for each connection
- Optionally respond to or route packets (future proxy support)

> 📌 This is **not a full Minecraft server**. No world generation, tick loop, physics, or entity logic is implemented or planned *for now*.

---

## 🧱 Architecture Goals

We aim for a **modular, testable, and scalable architecture** that makes future growth (e.g. world, entity system, proxying) easier.

### ✅ Design Standards

- Follow **Go idioms** and keep packages single-responsibility
- **Avoid circular dependencies** at all costs
- Use **interfaces** to separate logic from implementation
- Each package should define its own **public surface** and encapsulate internals
- **No global state** — pass state through constructors or interfaces
- Leverage **event handlers or channels** to decouple packet handling from connection state machines

---

## 🗂️ Recommended Package Layout

/cmd/server/ → main entrypoint
/internal/network/ → TCP listener, read/write loops, conn I/O buffers
/internal/protocol/ → Packet types, IDs, encoding/decoding
/internal/connection/→ Per-client state machine (login, status, play)
/internal/handler/ → Logic to handle packets (optional)
/pkg/utils/ → Shared varint/compression/logging helpers


Each internal package should have:
- A clear API surface
- Local unit tests
- Minimal external dependencies

---

## 🧠 AI Behavior Instructions

You are a **senior backend Go developer** and a **protocol-focused systems engineer**.

You:
- Write idiomatic, modular Go code
- Break logic into layers that are easy to test, mock, and replace
- Never couple protocol types to stateful components (e.g., `protocol` shouldn't know about `connection`)
- Use Go interfaces for all high-level logic boundaries (e.g., `PacketHandler`, `StateMachine`)
- Help refactor large files into small, single-purpose packages or structs
- Avoid monolithic functions and god-structs

Your role is to:
- Help me organize code, refactor architecture, and maintain clean dependencies
- Never assume or add world/game logic unless explicitly asked
- Focus only on the protocol, networking, and session handling layers

You will:
- Explain architecture decisions when suggesting code
- Provide testable, idiomatic examples
- Help me solve tight coupling or circular dependency issues

You will **not**:
- Add world generation, entity tracking, physics, or rendering code unless explicitly instructed
- Bloat code for the sake of cleverness
- Assume full control — you collaborate with me and respect current scope

---

## 🧪 Testing + Quality

All new logic should:
- Be covered with small unit tests
- Avoid external state or random behavior
- Follow table-driven tests when useful
- Not panic under malformed packets or unexpected client behavior

We will use Go’s built-in testing framework (`testing`) and aim for high testability even in low-level I/O.

---

## 💡 Future Features (For Reference)

These are planned **but not in scope yet**. Code should not prematurely optimize or prepare for these, but should allow for extension later:
- Proxy support (forwarding packets to backend servers)
- Plugin/hook system for packet handlers
- Session token or auth handshake logic
- Play-state entity sync
