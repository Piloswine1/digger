<script setup lang="ts">
import { ref } from 'vue'
import { useContainers } from './composables/useContainers'
import { Button } from '@/components/ui/button'
import ModeToggle from '@/components/ModeToggle.vue'
import { RefreshCw } from '@lucide/vue'

import ContainerCard from './components/ContainerCard.vue'
import LogViewer from './components/LogViewer.vue'
import ContainersSelector from './components/ContainersSelector.vue'
import type { ContainersMode } from './api'

const containersMode = ref<ContainersMode>('all')
const { data: containers, isLoading, isError, error, refetch } = useContainers(containersMode);

const selectedContainerId = ref('')
</script>

<template>
  <div class="min-h-screen pr6 max-w-[140rem] mx-auto">
    <div class="flex items-crnter gap-2 mb-6 pt-4">
      <ContainersSelector v-model="containersMode" class="flex-1" />
      <Button variant="outline" size="sm" @click="refetch()" :disabled="isLoading">
        <RefreshCw :class="isLoading && 'animate-spin'" />
        Refresh
      </Button>
      <ModeToggle />
    </div>

    <div v-if="isLoading" class="text-muted-foreground text-sm">Loading containers...</div>
    <div v-else-if="isError" class="text-destructive text-sm">
      Error: {{ (error as Error)?.message }}
    </div>
    <div v-else
      :class="[
        'grid grid-cols-1 lg:grid-cols-[450px_1fr] gap-6 items-start',
      ]">
      <div class="space-y-2 max-h-[calc(100vh-6rem)] overflow-y-auto py-2">
        <div v-if="!containers?.length" class="text-muted-foreground text-sm">
          No active containers
        </div>
        <ContainerCard
          v-for="c in containers"
          :key="c.Id"
          :container="c"
          v-model="selectedContainerId"
        />
      </div>

      <div
        v-if="!selectedContainerId"
        class="flex items-center justify-center h-64 text-muted-foreground border rounded-xl border-dashed"
      >
        Select a container to view logs
      </div>
      <LogViewer v-else :container-id="selectedContainerId" />
    </div>
  </div>
</template>
