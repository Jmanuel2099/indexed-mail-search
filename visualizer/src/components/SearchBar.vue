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
    <div class="search_container">
        <input v-model="searchTemr" type="text" placeholder="Search a term..." />
        <SearchResults />
    </div>
</template>

<style scoped>
.search_container {
    /* background-color: #212122; */
    width: 33%;
    height: 95%;
    padding: 10px;
    border-radius: 8px;
    position: fixed;
    top: 20px;
    left: 20px;
    overflow: auto;
    /* z-index: 9999; */
}

input[type="text"] {
    box-shadow: 10px 10px 10px rgba(142, 142, 143, 0.5);
    width: 100%;
    height: 45px;
    font-size: 15px;
    padding: 20px;
    border-radius: 8px;
    border: none;
    background-color: #bebebe;
    color: #5e5e5e;
    /* float: left; */
    margin-bottom: 20px;
}
</style>