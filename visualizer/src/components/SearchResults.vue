<script setup lang="ts">
import { useEmails } from '@/composables/useEmails';
import type { Email } from '@/interfaces/email';
import EmptyTray from './messages/EmptyTray.vue';
import Loading from './messages/Loading.vue'

const { isLoadingEmails, emails, onSelectEmail } = useEmails();

const onEmailClick = (email: Email) => {
    onSelectEmail(email)
}
</script>

<template>

    <div v-if="!isLoadingEmails && emails.length === 0">
        <EmptyTray />
    </div>

    <div v-else-if="isLoadingEmails && emails.length === 0">
        <Loading />
    </div>
    <ul v-else-if="emails.length > 0">
        <li v-for="email in emails" :key="email.message_id" @click="onEmailClick(email)">
            <h4>{{ email.subject }}</h4>
            <span>👤 {{ email.from.split('@')[0] }} - 📅 {{ new Date(email.date).toLocaleDateString() }}</span>
        </li>
    </ul>
</template>

<style scoped>
h5 {
    font-size: 20px;
    color: #0078bd;
    font-family: "Open Sans", sans-serif;
    margin-top: 50px;
}

li {
    background-color: #bebebe;
    margin-bottom: 10px;
    padding: 10px;
    border-radius: 5px;
    cursor: pointer;
}

h4 {
    color: #0078bd;
    font-size: 22px;
    margin-bottom: 10px;
    font-family: "Open Sans", sans-serif;
}

span {
    font-size: 16px;
    color: #2c2c2c;
    margin-bottom: 10px;
    font-family: "Open Sans", sans-serif;
}
</style>