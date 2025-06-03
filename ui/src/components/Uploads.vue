<script setup lang="ts">
import { onMounted } from 'vue';

interface UploadResponse {
  folders: string[];
  files: string[];
}

interface ExtensionEmojis {
  [key: string]: string;
}

interface ExtensionCallbacks {
  [key: string]: (file: string) => void;
}

const extensionEmojis: ExtensionEmojis = {
  "png": "🖼️",
  "jpg": "🖼️",
  "jpeg": "🖼️",
  "gif": "🎨",
  "mp3": "🎵",
  "wav": "🎵",
  "ogg": "🎵",
  "txt": "📄",
  "folder": "📁"
};

const test = (): void => {
  console.log("Test function called");
  alert("This is a test function for PNG files.");
}

const extensionCallbacks: ExtensionCallbacks = {
  "txt": test
};

async function askForFilesAndFoldersInUploads(): Promise<UploadResponse | undefined> {
  try {
    const response = await fetch("/list_uploads");
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error("Error fetching files and folders:", error);
    return undefined;
  }
}

const refreshUploads = async (event: Event): Promise<void> => {
  event.preventDefault();
  const data = await askForFilesAndFoldersInUploads();
  if (!data) return;

  const uploadsDiv = document.getElementById("uploadList");
  if (!uploadsDiv) {
    console.error("uploadList element not found");
    return;
  }
  
  uploadsDiv.innerHTML = "";

  data.folders?.forEach(folder => {
    const folderElement = document.createElement("button");
    folderElement.style.textAlign = "left";
    folderElement.textContent = `${extensionEmojis["folder"]} ${folder}`;
    uploadsDiv.appendChild(folderElement);
  });

  data.files?.forEach(file => {
    const fileElement = document.createElement("button");
    fileElement.style.textAlign = "left";
    const extension = file.split('.').pop()?.toLowerCase() || "";
    
    if (extensionEmojis[extension]) {
      fileElement.textContent = `${extensionEmojis[extension]} ${file}`;
    }

    if (extensionCallbacks[extension]) {
      fileElement.onclick = () => extensionCallbacks[extension](file);
    }
    uploadsDiv.appendChild(fileElement);
  });
}

onMounted(() => {
  refreshUploads(new Event('load'));
});
</script>

<template>
  <div id="uploads">
    <h1>Uploads</h1>
    <div id="uploadList" style="display:flex; flex-direction: column;"></div>
  </div>
</template>

<style scoped></style>
