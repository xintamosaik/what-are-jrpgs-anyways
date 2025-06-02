## Architecture Questions
1. **Editor-Game Relationship**: How tightly coupled should the editor and game be? Should they share code or be completely separate applications?
- The server needs to serve both for development. For running the game it needs to be possible to only run the game. 

2. **API Design**: Are there specific API design patterns you want to follow for endpoints beyond `/api/ping`?
- api/ping is just a test. I did not decide about the patterns. 

3. **State Management**: How should state be shared between the editor and game? Is the shared JSON file the only mechanism?
- Since the game needs to run without the editor the editor needs to persist everything in files and databases. For now we will use JSON or something simple like it. If the need arises we will do sqlc and postgres other nosql ideas.

## Technical Questions
4. **Go Implementation**: Are there specific Go packages or patterns preferred for the HTTP servers?
- If possible I want to use vanilla, standard library or however to call it. We will probably use templ for web templating.

5. **Frontend Framework**: The README mentions "React and Tailwind CSS planned" - is vanilla JS preferred for now, or should I start with React?
- Good question. This is nice as a SPA I guess. Lets do that. Vite, then?

6. **Canvas Implementation**: What specific Canvas features are most important to implement first for the map editor?
- I guess we need to have something to show tilesets and something to show the map that uses the tiles. My guess is as good as yours for now. I think we need to upload images at some point. I will add some reference material next?

7. **Testing Strategy**: What's the approach for testing both frontend and backend components?
We follow Casey Muratori, Primeagen, Trashdev and Tee for now. Acceptance tests, end2end tests are primary. We will not use TDD. We will have unit tests for things that are important and complicated. We will test from the outside not implementations. 

## Development Process Questions
8. **Development Environment**: Are there specific tools or setup instructions for local development?
Go needs to be installed. We might need to add vite I guess so prepare node with something like nvm (your choice) or use bun. I think both work. For now lets use node.

9. **Contribution Workflow**: Is there a preferred git workflow (feature branches, PR templates, etc.)?
trunk for now. If we go full open source we need PRs and such I guess.

10. **Acceptance Criteria**: Beyond the technical requirements, what makes a feature "complete"?
We will be cautious with features. We will dogfeed the editor to us and make games. A feature is complete when users (us for now) say so. 


## Roadmap Questions
11. **Prioritization**: After the initial ping implementation, which features have highest priority?
The first thing will be the path to using tilesets for tilebased maps.

12. **User Testing**: How will we validate that the implementation meets the needs of the target audience (14-year-olds)?
For now we need to mentalize it ourselves. Things will probalby not go 100% there. We need exposure but for that we first need traction. For that we need MVP.

13. **Plugin System**: When should I start thinking about the plugin architecture mentioned in the README?
Now if possible. I do not yet know what is out there.
