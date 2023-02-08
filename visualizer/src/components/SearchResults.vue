<script setup lang="ts">
import { useEmails } from '@/composables/useEmails';
import { computed } from 'vue';

const { isLoadingEmails, emails } = useEmails();

interface formattedEmail {
    message_id: string;
    subject: string;
    date: Date;
    from: string;
}

const formattedEmails = computed((): formattedEmail[] => {
    let arr: formattedEmail[] = [];
    emails.value.forEach(email => {
        let emailFormate: formattedEmail = {
            message_id: email.message_id,
            subject: email.subject,
            from: email.from.split('@')[0],
            date: new Date(email.date)
        };
        arr.push(emailFormate);
    });
    return arr;
})


</script>

<template>
    <div v-if="isLoadingEmails">
        <h5>Looking for emails...</h5>
        <span>Wait please !!</span>
    </div>

    <ul v-else-if="emails.length > 0">
        <li v-for="email in formattedEmails" :key="email.message_id">
            <h4>{{ email.subject }}</h4>
            <p>From: {{ email.from }}</p>
            <span> Date: {{ email.date.toLocaleDateString() }}</span>
        </li>
    </ul>
</template>

<style scoped>

</style>