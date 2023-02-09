<script setup lang="ts">
import { useEmails } from '@/composables/useEmails';
import type { Email } from '@/interfaces/email';
import { computed } from 'vue';

const { isLoadingEmails, emails, onSelectEmail } = useEmails();

const onEmailClick = (email: Email) => {
    onSelectEmail(email)
}

</script>

<template>

    <div v-if="isLoadingEmails" class="loading-container">

        <h5>Looking for the term...</h5>
        <p>Wait pleas</p>
    </div>

    <ul v-else-if="emails.length > 0">
        <li v-for="email in emails" :key="email.message_id" @click="onEmailClick(email)">
            <h4>{{ email.subject }}</h4>
            <span>👤 {{ email.from.split('@')[0] }} - 📅 {{ new Date(email.date).toLocaleDateString() }}</span> 
        </li>
    </ul>
</template>

<style scoped>
.loading-container {
    text-align: center;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
}

h5 {
    font-size: 20px;
    color: #81bcf7;
    font-family: "Open Sans", sans-serif;
    margin-top: 50px;
}

li {
    background-color: #2C3E50;
    margin-bottom: 10px;
    padding: 20px;
    border-radius: 10px;
    cursor: pointer;
}

h4 {
    color: #c8faf7;
    font-size: 22px;
    margin-bottom: 10px;
    font-family: "Open Sans", sans-serif;
}

span {
    font-size: 16px;
    color: #666666;
    margin-bottom: 10px;
    font-family: "Open Sans", sans-serif;
}
</style>