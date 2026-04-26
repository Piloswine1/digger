<script setup lang="ts">
import { ref } from 'vue'
import { useContainers } from './composables/useContainers'
import { Button } from '@/components/ui/button'
import { RefreshCw } from '@lucide/vue'
import { cn } from '@/lib/utils'

import ContainerCard from './components/ContainerCard.vue'
import LogViewer from './components/LogViewer.vue'

const { data: containers, isLoading, isError, error, refetch } = useContainers()

const selectedContainerId = ref('')
</script>

<template>
  <div class="min-h-screen p-6 max-w-7xl mx-auto">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold">Containers</h1>
      <Button variant="outline" size="sm" @click="refetch()" :disabled="isLoading">
        <RefreshCw :class="cn('h-4 w-4', isLoading && 'animate-spin')" />
        Refresh
      </Button>
    </div>

    <div v-if="isLoading" class="text-muted-foreground text-sm">Loading containers...</div>
    <div v-else-if="isError" class="text-destructive text-sm">
      Error: {{ (error as Error)?.message }}
    </div>
    <div v-else class="grid grid-cols-1 lg:grid-cols-[350px_1fr] gap-6 items-start">
      <div class="space-y-2">
        <div v-if="!containers?.length" class="text-muted-foreground text-sm">
          No active containers
        </div>
        <ContainerCard
          v-for="c in containers"
          :key="c.Id"
          :container="c"
          :selected="selectedContainerId === c.Id"
          @click="selectedContainerId = c.Id"
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
