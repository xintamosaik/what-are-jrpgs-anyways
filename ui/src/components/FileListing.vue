<script setup lang="ts">
import { onMounted, ref, computed } from 'vue';
import FolderNode from './FolderNode.vue';

interface UploadResponse {
  folders: string[];
  files: string[];
}

type ExtensionEmojis = Record<string, string>;

interface ExtensionCallbacks {
  [key: string]: (file: string) => void;
}

const props = defineProps<{
  folder?: string;
  basePath?: string;  // Add this prop
}>();

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

const currentPath = computed(() => {
  if (!props.folder) return '';
  return props.basePath ? `${props.basePath}/${props.folder}` : props.folder;
});

async function refreshUploads(): Promise<void> {
  try {
    const path = props.folder ? currentPath.value : '';
    const url = path ? `/list_uploads?folder=${encodeURIComponent(path)}` : "/list_uploads";
    const response = await fetch(url);
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
  <div class="file-listing">
    <FolderNode 
      v-for="folder in uploads.folders"
      :key="folder"
      :folder="folder"
      :base-path="currentPath"
    />
    
    <button
      v-for="file in uploads.files"
      :key="file"
      class="node"
      @click="handleFileClick(file)"
    >
      {{ getEmoji(getFileExtension(file)) }} {{ file }}
    </button>
  </div>
</template>

<style scoped>
.file-listing {
  display: flex;
  flex-direction: column;
}

.node {
  text-align: left;
}
</style>
