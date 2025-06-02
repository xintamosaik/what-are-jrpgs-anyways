# JRPG Maker: README

## Project Description

A web-native, progressively-enhanced tool for creating tile-based maps and interactive scenes, initially focused on JRPG-style games. The tool is designed for quick, zero-install play and gradual onboarding into deeper customization and self-hosting.

**Vision:**  
Empower young or novice creators to make games directly in the browser, while enabling advanced users to self-host and extend the system.

---

## Cornerstone Concepts

- **Spatial Primacy:** All meaningful data (tiles, collisions, callbacks) is spatially indexed on a 2D grid.
- **Tiles as Entities:** Tiles can carry optional behavior, metadata, or interactivity; they act like ECS-style entities.
- **Progressive Enhancement:** Users can start creating instantly, with options for login-based persistence, self-hosting, and collaboration.
- **Dogfooding as Design:** The tool is built for creators to make games themselves first; community enablement is secondary but important.

---

## Initial Technical Stack

- **Frontend:** HTML+JS (React and Tailwind CSS planned)
- **Backend:** Go (for performance, static binaries, and easy deployment)
- **Persistence:** Postgres and/or NoSQL/JSON (decision pending)

---

## Big Architectural Questions

- Relational vs. Document Storage
- Frontend-Backend Contract (tight vs. loose coupling)
- Auth Model (simple, low-friction for self-hosters)
- Deployment Tiers (hosted, local, browser-only, installer, server)

---

## Initial User Scenario

A 14-year-old with a browser and basic digital literacy can:
- Upload tilesets
- Paint maps
- Export to a format usable in simple game engines
- Never be blocked by complex runtime, infrastructure, or technical vocabulary

---

## Design Ethos

- Be composable like UNIX tools
- Be transparent like open source
- Be responsive like a good game: feedback, agency, low-friction flow
- This is a tool, not a platform—tools shape what we can make

---

## 🧭 High-Risk Decisions (Summary)

1. **Runtime & Execution Model:** Web app, locally deployable or server-hosted, multiplayer by default.
2. **Rendering Technology:** HTML5 `<canvas>` for maps/tilesets, DOM for forms/dialogs, layering/zoom/selection desired.
3. **Data Format & Storage:** Postgres for some data, mostly NoSQL/JSON, human-editable files, versioning TBD.
4. **Modularity vs Monolith:** Editor and game engine are separate, plugin boundaries defined early.
5. **Scripting / Custom Logic:** JS scripting (possibly visual), interpreted, sandboxing is a future concern.
6. **Pathfinding & World Topology:** Layered maps, ECS-like logic, scenes as maps.
7. **Save/Load/Game State:** Prefer backend state, minimal frontend state, serialization/crash recovery TBD.
8. **Editor Architecture:** Editor and game are separate (tabs/iframes), drag-and-drop, tile painting, hot reload.
9. **Distribution & Deployment:** GitHub Releases, offline support, asset packs/mods/DLC planned.

---

## 🏁 Immediate XP Slice: What to Build First

**Goal:**  
*An empty app and an empty game that send "hello" to each other, running as two Go binaries with two UIs.*

### Minimal Starter Plan

1. **Two Go Apps**
    - `editor.go` → serves `localhost:3000` (Editor UI)
    - `game.go` → serves `localhost:3001` (Game UI)
    - Each exposes `GET /api/ping` returning `{ "message": "pong" }`

2. **Frontend**
    - Editor: HTML+JS canvas + button, sends ping to backend, displays response.
    - Game: HTML+JS canvas, shows a moving shape, same ping cycle.

3. **Shared Data**
    - Dummy shared file (e.g. `shared/map.json`) both backends can read/write.

---

## 🧩 Architectural Guidance

- **Multiplayer by Design:**  
  Design all game logic as if networked (action dispatching, not direct state mutation).
- **Plugin Boundaries:**  
  Define Go interfaces for backend plugins and a plugin registry pattern for frontend.
- **Scripting & Sandboxing:**  
  Use a stub like `evalThis(script, context)` for now; plan for sandboxing later.

---

## 🔍 Tech Spikes (Next 1–2 Weeks)

- **Go Embed & Packaging:** Use `go:embed` for static files, test multi-platform builds.
- **Canvas Layering & Input:** Prototype canvas layers, add zoom/selection.
- **Simple Save/Load:** Minimal map in editor, save/load as JSON, verify handoff between editor/game/backend.

---

## 🧩 Safe to Defer

- Quest editor
- Combat system
- NPC logic/scripting
- Visual programming interface
- Tile slicing/asset pipeline

---

## ❗ Risk Hotspots

- **Multiplayer:** Isolate game state logic for future networked sync.
- **Plugin System:** Define interfaces and boundaries early.
- **Sandboxing:** Plan for user script isolation, but don't block progress now.

---

## ✅ Next Actions

1. Scaffold two Go servers (`editor.go`, `game.go`) serving static UIs.
2. Implement `/api/ping` endpoint in both.
3. Create minimal HTML+JS UIs for each, with canvas and ping button.
4. Add a shared JSON file for future data handoff.
5. Document the round-trip "hello" flow in logs.

---

## 📁 Suggested Folder Structure

```
/server/
  editor.go
  game.go
/shared/
  map.json
/ui/
  editor/
    index.html
    editor.js
  game/
    index.html
    game.js
/README.md
```

---

## 🚦 Minimal XP Step: End-to-End "Ping"

**Goal:**  
Have a Go server serve a static HTML file with a button. When clicked, the button sends a request to `/api/ping`, and the server responds with `"pong"`. The response is shown in the browser.

**Why?**  
This proves your backend, frontend, and API contract all work—end-to-end.

---

## 💡 Want a Ready-to-Run Scaffold?

If you want, I can generate a starter repo structure with Go code, HTML, and JS for this XP slice. Just ask!

---