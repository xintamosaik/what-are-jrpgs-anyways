<script setup lang="ts">
import { ref } from 'vue';
import FileListing from './FileListing.vue';

interface FolderProps {
  folder: string;
}

const opened = ref(false);
const props = defineProps<FolderProps>();

const handleFolderClick = () => {
  opened.value = !opened.value;
};
</script>

<template>
  <div class="folder-node">
    <button 
      class="node"
      :class="{ 'node--open': opened }"
      @click="handleFolderClick"
    >
      {{ opened ? '📂' : '📁' }} {{ props.folder }}
    </button>
    
    <div 
      v-if="opened" 
      class="folder-content"
    >
      <FileListing :folder="props.folder" />
    </div>
  </div>
</template>

<style scoped>
.folder-node {
  display: flex;
  flex-direction: column;
}

.node {
  text-align: left;
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px 8px;
}

.node:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.node--open {
  font-weight: bold;
}

.folder-content {
  margin-left: 20px;
}
</style>
