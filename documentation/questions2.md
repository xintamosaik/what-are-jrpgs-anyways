# Follow-up Questions
## Technical Implementation Questions

1. **Templ & Frontend Integration**: Since you mentioned templ for templating but also React with Vite, how do you envision these working together? Will templ handle server-side rendering with React taking over on the client?
- Good catch. Lets use vite for the web app and do as much as we can and then see. Wait, we can not send html over to react in an htmx way, right?

2. **Canvas References**: You mentioned adding reference material for Canvas implementation. Are there any specific existing tile editors (like Tiled, RPG Maker, etc.) whose UX you'd like to emulate or avoid?
- I will go with RPG maker for now but steal some ideas from godot. Scenes for example. But that is not that important. Most of this project is in "hello world" state or below. When we get it wired up we can move around pixels and forms. 

3. **File Structure**: What's your preferred organization for the project? Should we separate backend/frontend into distinct directories, or organize by feature?
- by feature if possible. Things will move a bit for a while anyway.

## Data & Storage Questions

4. **Initial Data Format**: For the JSON-based storage of maps and tilesets, do you have a rough schema in mind, or should we design this from scratch?
Too much will change. We will nail it down later.

5. **Asset Management**: How should we handle tileset images? Store file paths, embed base64, or something else?
Good question. simple object storage like AWS does but far less complex?

## Next Steps Questions

6. **Immediate Task**: Based on the answers about prioritization, would you like to focus next on:
> - Setting up the React+Vite frontend
I will probably, yeah
> - Implementing basic tileset upload/display
After vite
- Creating the map editor canvas
In react
- Something else?
That is it for now.

7. **Development Workflow**: Would a basic README update with setup instructions (Go, Node/npm, running the servers) be helpful as a next step?
Not yet. That is too detailed for now. There will be loads of changes.

