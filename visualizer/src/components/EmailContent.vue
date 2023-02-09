<script setup lang="ts">
import { useEmails } from '@/composables/useEmails';
import { computed } from 'vue';

const { selectedEmail, selectedEmailReady } = useEmails();

const conditionsToShowEmail = computed((): boolean => !selectedEmailReady.value)
</script>

<template>
    <div v-if="conditionsToShowEmail">
        seleccione un email!!!
    </div>
    <div v-else>
        <h3>{{ selectedEmail?.subject }}</h3>
    </div>
</template>

<style scoped>

</style>