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

- **Frontend:** HTML, CSS, JavaScript. Initially, this will be a single `index.html`, a single `style.css`, and JavaScript (either inline or a single `script.js`). We are using Fixi (vendored) for client-server communication. (React and Tailwind CSS are planned for later, as the project grows).
- **Backend:** Go. Initially, this will be a single `main.go` file. (Chosen for performance, static binaries, and easy deployment).
- **Persistence:** Postgres and/or NoSQL/JSON (decision pending).

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
*A single Go application serving a web UI. The UI will have a button that communicates with the backend via Fixi to fetch and display data or content.*

### Minimal Starter Plan

1.  **Single Go App**
    *   `main.go` → serves `localhost:3000` (Web UI)
    *   Exposes initial API endpoints, including `GET /api/ping` returning `{ "message": "pong" }`, and an endpoint to serve placeholder HTML content.

2.  **Frontend**
    *   A single `index.html` file for structure, a `style.css` for basic styling, and JavaScript (e.g., in `script.js` or inline) for interactivity.
    *   Uses Fixi (vendored) for backend communication.
    *   Features:
        *   A button that, when clicked, uses Fixi to request data/content from the Go backend.
        *   The main content area is updated with the response received from the backend (e.g., placeholder HTML or a "pong" message).

3.  **Shared Data (for near future)**
    *   A dummy shared file (e.g. `shared/data.json`) that the Go backend can read from and write to, demonstrating basic data handling.

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

1.  Finalize the single Go server (`main.go`) to robustly serve the static `index.html`, `style.css`, and any JavaScript files from a `/static` directory (or using `go:embed`).
2.  Ensure Fixi is correctly vendored and integrated for the existing button-to-backend communication that replaces main content.
3.  Implement a distinct `GET /api/ping` endpoint in `main.go` that returns `{ "message": "pong" }`. Add a new, separate button or mechanism in the frontend to specifically test this "ping" endpoint and display its response.
4.  Create an initial `shared/data.json` file. Implement basic functionality in the Go backend to read from this file and potentially an endpoint for the frontend to fetch this data.
5.  Document the current simplified architecture, the Fixi integration, and the end-to-end "ping" and content-loading flow in the logs and/or README.
6.  Begin planning the first interactive feature beyond basic content swapping, e.g., displaying content from `shared/data.json` or a very simple canvas interaction.

---

## 📁 Suggested Folder Structure

```
/
  main.go                 # Single Go backend file
/static/                  # For all static assets
  index.html              # Main HTML file
  style.css               # Main CSS file
  script.js               # Main JavaScript file (if not inline)
/vendor/
  fixi/                   # Vendored Fixi library files
    fixi.min.js           # (or however Fixi is structured)
/shared/
  data.json               # Dummy shared data file (for future use)
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