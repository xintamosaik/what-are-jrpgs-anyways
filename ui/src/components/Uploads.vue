<script setup lang="ts">

</script>
<template>
  <div id="uploads">
    <h1>Uploads</h1>

    <div id="uploadList" style="display:flex; flex-direction: column;"></div>
  </div>
</template>

<script lang="js">
async function askForFilesAndFoldersInUploads() {
  try {
    const response = await fetch("/list_uploads:3000");
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error("Error fetching files and folders:", error);
  }
}

const extensionEmojis = {
  "png": "🖼️",
  "jpg": "🖼️",
  "jpeg": "🖼️",
  "gif": "🎨",
  "mp3": "🎵",
  "wav": "🎵",
  "ogg": "🎵",
  "txt": "📄",

};

function test(event) {
  console.log("Test function called");
  alert("This is a test function for PNG files.");
}
const extensionCallbacks = {
  "txt": test
};

document.addEventListener("DOMContentLoaded", () => {

  refreshUploads(new Event('load'));
});
function refreshUploads(event) {
  event.preventDefault();
  askForFilesAndFoldersInUploads().then(data => {
    const uploadsDiv = document.getElementById("uploadList");
    uploadsDiv.innerHTML = ""; // Clear previous content

    const folders = data.folders || [];
    const files = data.files || [];


    folders.forEach(folder => {
      const folderElement = document.createElement("button");
      folderElement.style.textAlign = "left"; // Align text to the left
      folderElement.textContent = `${extensionEmojis["folder"]} ${folder}`;

      uploadsDiv.appendChild(folderElement);
    });

    files.forEach(file => {
      const fileElement = document.createElement("button");
      fileElement.style.textAlign = "left"; // Align text to the left
      const extension = file.split('.').pop().toLowerCase();
      if (extensionEmojis[extension]) {
        fileElement.textContent = `${extensionEmojis[extension]} ${file}`;
      } // we only allow the extensions we have emojis for

      if (extensionCallbacks[extension]) {
        fileElement.onclick = () => extensionCallbacks[extension](file);
      }
      uploadsDiv.appendChild(fileElement);
    });
  });
}
</script>
<style scoped></style>
