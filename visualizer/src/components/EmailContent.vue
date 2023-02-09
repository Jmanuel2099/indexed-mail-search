<script setup lang="ts">
import { useEmails } from '@/composables/useEmails';
import { computed } from 'vue';
import Welcome from './messages/Welcome.vue'

const { selectedEmail, selectedEmailReady } = useEmails();

const conditionsToShowEmail = computed((): boolean => !selectedEmailReady.value);
</script>

<template>
    <div v-if="conditionsToShowEmail">
        <Welcome />
    </div>
    <div v-else>
        <section class="emial_section">
            <h1 style="color: #0078bd">{{ selectedEmail?.subject }}</h1>
            <h3> 👤From:{{ selectedEmail?.from }} ( 📅 {{ selectedEmail?.date }} )</h3>
            <h3>To: {{ selectedEmail?.to }}</h3>
            <p>{{ selectedEmail?.content }}</p>
        </section>

    </div>
</template>

<style scoped>
.emial_section {
    margin: 30px 50px 50px 50px;
}

@media (min-width: 1024px) {

    .emial_section h1,
    .emial_section h3,
    .emial_section p {
        padding: 10px;
        color: #575656;
        font-family: "Open Sans", sans-serif;
        text-align: left;
    }
}
</style>