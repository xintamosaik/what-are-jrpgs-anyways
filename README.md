# Go Web RPG App

This project is a web application built with Go as the backend and HTML for the frontend. It allows users to send a "hello" message from the app to the game.

## Project Structure

```
go-web-rpg-app
├── cmd
│   └── server
│       └── main.go        # Entry point of the application
├── internal
│   ├── app
│   │   └── app.go         # Application logic
│   └── game
│       └── game.go        # Game-related functionality
├── web
│   ├── static
│   │   └── style.css       # CSS styles for the frontend
│   └── templates
│       └── index.html      # Main HTML template for the frontend
├── go.mod                  # Module definition for the Go application
├── go.sum                  # Checksums for module dependencies
└── README.md               # Documentation for the project
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-web-rpg-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/server/main.go
   ```

4. **Access the application:**
   Open your web browser and navigate to `http://localhost:8080` to interact with the app.

## Usage Guidelines

- The application provides a simple interface to send a "hello" message to the game.
- Responses from the game will be displayed on the frontend.
- You can modify the application logic in `internal/app/app.go` and the game functionality in `internal/game/game.go` as needed.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

# JRPG Maker: Architecture & XP Plan

This document synthesizes the architectural decisions, priorities, and next actions for the JRPG Maker project, based on the evolving discussion and feedback. It is intended as a living README for the early XP (eXtreme Programming) phase.

---

## 🧭 High-Risk Decisions (Summary)

**1. Runtime & Execution Model**
- Web app, locally deployable or server-hosted.
- Both always-on server and local/in-memory play.
- Multiplayer is a goal from the start.

**2. Rendering Technology**
- Mostly HTML5 `<canvas>` for maps/tilesets.
- DOM forms for UI/dialogs.
- Layering, transitions, zoom, and selection are desired.

**3. Data Format & Storage**
- Postgres for some data, mostly NoSQL/JSON for game content.
- Human-editable, hackable files preferred.
- Versioning/migrations: open question.

**4. Modularity vs Monolith**
- Editor and game engine are separate.
- Games can be exported and self-hosted.
- Plugin boundaries are to be defined early.

**5. Scripting / Custom Logic**
- JS for scripting, possibly with a visual editor.
- Scripts interpreted, not compiled.
- Sandboxing is a future concern.

**6. Pathfinding & World Topology**
- Layered maps, ECS-like logic.
- Scenes as maps, with teleporters/edges.

**7. Save/Load/Game State**
- Prefer backend state, minimal frontend state.
- Serialization and crash recovery are open questions.

**8. Editor Architecture**
- Editor and game are separate, possibly in different tabs or iframes.
- Drag-and-drop, tile painting, and hot reload are goals.

**9. Distribution & Deployment**
- GitHub Releases for binaries.
- Offline support is a goal.
- Asset packs/mods/DLC are planned.

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

### Multiplayer by Design
- Design all game logic as if networked, even if not implemented yet.
- Use action dispatching (`game.dispatch({ type: "MOVE", ... })`) instead of direct state mutation.
- This enables logging, syncing, replay, and future multiplayer.

### Plugin Boundaries
- Define Go interfaces for backend plugins, even if only one implementation exists.
- Frontend can use a plugin registry pattern.
- Dynamic loading can come later; focus on clear seams now.

### Scripting & Sandboxing
- User JS scripts will be part of the game.
- For now, use a stub like `evalThis(script, context)` and add a TODO for sandboxing.
- In the future, consider iframes, Web Workers, or server-side isolation.

---

## 🔍 Tech Spikes (Next 1–2 Weeks)

- **Go Embed & Packaging:**  
  Use `go:embed` for static files, test multi-platform builds.

- **Canvas Layering & Input:**  
  Prototype canvas layers (background, sprite, UI), add zoom/selection.

- **Simple Save/Load:**  
  Minimal map in editor, save/load as JSON, verify handoff between editor/game/backend.

---

## 🧩 Safe to Defer

- Quest editor
- Combat system
- NPC logic/scripting
- Visual programming interface
- Tile slicing/asset pipeline

---

## ❗ Risk Hotspots

- **Multiplayer:**  
  Isolate game state logic so it can be swapped for networked sync later.

- **Plugin System:**  
  Define interfaces and boundaries early, even if not dynamic yet.

- **Sandboxing:**  
  Plan for user script isolation, but don't block progress now.

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

## 💡 Want a Ready-to-Run Scaffold?

If you want, I can generate a starter repo structure with Go code, HTML, and JS for this XP slice. Just ask!

---