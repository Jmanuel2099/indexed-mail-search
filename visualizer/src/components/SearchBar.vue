<script setup lang="ts">
import { ref, computed } from 'vue';
import { useEmails } from '@/composables/useEmails';
import SearchResults from './SearchResults.vue';

const debouncedValue = ref('')
const debounceTimeout = ref()
const { searchEmailsByTerm } = useEmails();

const searchTemr = computed({
    get() {
        return debouncedValue.value
    },
    set(newVal: string) {
        if (debounceTimeout.value) clearTimeout(debounceTimeout.value)

        debounceTimeout.value = setTimeout(() => {
            debouncedValue.value = newVal;
            searchEmailsByTerm(newVal);
        }, 2000);
    },
});

</script>

<template>
    <div >
        <input v-model="searchTemr" type="text" placeholder="Search a word..." />
        <SearchResults />
    </div>
</template>

<style scoped>

</style>