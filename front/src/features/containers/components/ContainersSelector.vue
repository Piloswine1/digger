<script setup lang="ts">
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'

import type { ContainersMode } from '../api';

const mode = defineModel<ContainersMode>({default: 'all'});
</script>

<template>
  <DropdownMenu>
    <div :class="$attrs.class">
      <DropdownMenuTrigger as-child>
          <Button variant="outline" class="font-bold">
            <span v-if="mode === 'all'">
              Containers
            </span>
            <span v-if="mode === 'active'">
              Containers (active)
            </span>
          </Button>
      </DropdownMenuTrigger>
    </div>
    <DropdownMenuContent align="end">
      <DropdownMenuItem @click="mode = 'all'">All</DropdownMenuItem>
      <DropdownMenuItem @click="mode = 'active'">Active</DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>
