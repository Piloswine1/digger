<script setup lang="ts">
import { cn } from '@/lib/utils'

const props = defineProps<{
  type?: string
  placeholder?: string
  min?: string | number
  max?: string | number
  step?: string | number
}>()

const model = defineModel<string | number>()
</script>

<template>
  <input
    :type="type ?? 'text'"
    :placeholder="placeholder"
    :min="min"
    :max="max"
    :step="step"
    :class="
      cn(
        'flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50',
        $attrs.class,
      )
    "
    v-model="model"
  />
</template>
