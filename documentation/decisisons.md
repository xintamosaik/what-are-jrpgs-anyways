Absolutely. Here’s a curated list of **high-risk decisions** that experienced architects typically address early in any **somewhat complex app**, especially one with UI + logic + state like your RPG Maker:

---

## 🔴 HIGH-RISK DECISIONS TO TACKLE EARLY

### 1. **Runtime & Execution Model**

> “How and where does the app run?”

* Will it run as a **web app**, **desktop app**, or **hybrid**?
    It will be a web app.
* Will the user need a runtime (Node, Python) or is it a **self-contained binary** (e.g. Go)?
    Not to run the app. I am not sure how they would deploy their game though.
* Is there an **always-on server**, or is everything local?
    Both. It will be locally deployable or run on my server or run on their server. I will give limited space for logged users, non-logged can play in-memory. 
* Do you want **multiplayer**, or is it single-player only?
    I want multiplayer. I am not sure how though. I think that is the challenge for the game engine backend that gets deployed.


### 2. **Rendering Technology**

> “How is the visual game/editor content displayed?”

* HTML grid vs. `<canvas>` vs. WebGL vs. native UI (e.g. SDL, Wails)
    Most stuff will be canvas since we will have tilesets and maps that implement tilesets. It will be vanilla canvas with a small custom framework I guess.
* Will tile rendering be DOM-based or drawn on a canvas buffer?
    Mostly buffer. And I might be able to have multimonitor support. But not sure Would be great.
* Does your rendering system support layering, transitions, zoom, selection?
    I hope so.

Side-note: The app will have forms in Dom. Dialog or popover API.


### 3. **Data Format & Storage Model**

> “How is data (maps, quests, characters) stored and loaded?”

* JSON, YAML, SQLite, embedded Go structs?
    Postgres for some stuff. Most will be nosql. JSON probably. I do not like YAML. 
* Are files human-editable? Moddable? Compressed?
    As far as possible. I want it to be hackable.
* Can you support versioning, migrations, and schema changes?
    How would I?
* Flat vs. nested vs. relational structures?
    Depends on the data. Entities will sometimes be more document oriented with slight nestedness (1 level max)

### 4. **Modularity vs Monolith**

> “How separated are logic, UI, data, runtime?”

* Is the editor engine decoupled from the game engine?
    Good question. I think so. It will probably be a browser window or tab that autoupdates after changes.
* Can you export and run a game separately from the editor?
    That is a good point. It will "make" web browser based games, but these would depend on a backend. The game engine and the data modified. These could be self-hosted. Would be nice for users to package them. 
* Are rendering, logic, and persistence modular?
    What does that mean?
* Will you support plugins?
    Frontend. Probably. Backend. I would like to but I do not know how yet.


### 5. **Scripting / Custom Logic**

> “How will users define game logic?”

* Will you implement a real scripting language (Lua, JS)?
    JS. Possibly some visual editor for some stuff so not everyone has to code. But it will generate JS that is hackable.
* Or a visual system (like Blockly or node editors)?
    Yeah. I might
* Will you interpret it or compile it?
    JS needs to be interpreted.
* How do you sandbox it?
    Sandbox?


### 6. **Pathfinding and World Topology**

> “How does movement work, and what data structures underpin the map?”

* Will you support variable movement cost? (tiles with weights?)
    There will be layers like in RPG maker and some of them will be binary maps like "is this a tile you can even walk on?"
    I think under the hood there will be a logic very close to ECS so new layers with logic can be added.
* Will NPCs use A\*? Are graphs generated at runtime or preprocessed?
    Sorry? I don't get this question.
* Are there multiple map layers or zones? Wrap-around maps?
    I hope to be able to make maps as scenes. Every scene can be as big as possible and I will try to enable teleporters and teleport "edges".

### 7. **Save / Load / Game State Architecture**

> “How do you persist and restore the entire state of the game world?”

* Will you use snapshots? Events? State diffs?
    I will try to avoid state as much as possible on the frontend and try to store most stuff in the backend I guess. Not sure. This will be a tough nut to crack or rather it might need to be explored.
* Do you serialize internal structs or a clean exported schema?
    I do not worry too much about the game maker here, but the engine might need to talk fast and binary.
* What happens if the game crashes?
    Good question. It would be great to auto dump everything.


### 8. **Editor Architecture**

> “How do users build games inside your app?”

* Will the editor use a separate mode or share the runtime?
    I think it should be separeated, so in a way the game is always a thing (CI CD locally)
* Do you need drag-and-drop? Tile painting? Hot reload?
    Drag and drop. Yes. Tile painting. Future. Hot reaload - what?
* Can users preview logic or only test via play mode?
    I think that will be my task to enable preview. Tough one. 
* Will maps/quests be visual or just data forms?
    Maps will be visible. Quests? Not sure how I design them. 

### 9. **Distribution & Deployment**

> “How do users get the app and run it?”

* Will you ship binaries via GitHub Releases? Or via a custom installer?
    I think I will use github for realeases.
* Does it run offline fully?
    I hope to enable it.
* Will you allow asset packs / mods / DLC loading?
    Yeah I want to allow the tile editor to cut into different sized tilesets, They mostly align on a power of 2 basis and that should enable something. 
    Mods? I hope so. Since backend and frontend introduce complexity it looks like a bit of a challenge. DLC? No. Yes. Devs need to work on their game and after distributing the base game they need to mark something new. Not sure. This is a tough nut too. I will try to build something myself with this in mind and see what it does. 
* Is there a plugin system?
    I really hope so.


## Want to Prioritise?

If you'd like, I can help prioritise these *for your specific project* — for example:

* What to prototype first?
    Yeah I want to make it as in XP. from 0 to 1. The whole thing always runs. First just an empty app and an empty game that sends hello to each other. 
* What needs a tech spike vs what can wait?
    This is a mess potentially. I need to be careful to get a high degree of coherence and a low degree of knowledge coupling. 
* Which ones to settle on this week vs later?
    This week will be the XP thing. 

