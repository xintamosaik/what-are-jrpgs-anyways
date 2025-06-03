<script setup lang="ts">
import { onMounted, ref } from 'vue';

interface UploadResponse {
  folders: string[];
  files: string[];
}

type ExtensionEmojis = Record<string, string>;

interface ExtensionCallbacks {
  [key: string]: (file: string) => void;
}

const uploads = ref<UploadResponse>({ folders: [], files: [] });

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

const getFileExtension = (file: string): string => {
  return file.split('.').pop()?.toLowerCase() || "";
};

const getEmoji = (extension: string): string => {
  return extensionEmojis[extension] || "";
};

async function refreshUploads(): Promise<void> {
  try {
    const response = await fetch("/list_uploads");
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    uploads.value = await response.json();
  } catch (error) {
    console.error("Error fetching files and folders:", error);
  }
}

const handleFileClick = (file: string) => {
  const extension = getFileExtension(file);
  if (extensionCallbacks[extension]) {
    extensionCallbacks[extension](file);
  }
};

onMounted(() => {
  refreshUploads();
});
</script>

<template>
  <div id="uploads">
    <h1>Uploads</h1>
    <div class="upload-list">
      <button
        v-for="folder in uploads.folders"
        :key="folder"
        class="upload-item"
      >
        {{ extensionEmojis.folder }} {{ folder }}
      </button>
      
      <button
        v-for="file in uploads.files"
        :key="file"
        class="upload-item"
        @click="handleFileClick(file)"
      >
        {{ getEmoji(getFileExtension(file)) }} {{ file }}
      </button>
    </div>
  </div>
</template>

<style scoped>
.upload-list {
  display: flex;
  flex-direction: column;
}

.upload-item {
  text-align: left;
}
</style>
