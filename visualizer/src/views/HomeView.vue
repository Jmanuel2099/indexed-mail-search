<script setup lang="ts">
import EmailContent from '@/components/EmailContent.vue';
import SearchBar from '@/components/SearchBar.vue';


</script>

<template >
    <div class="home_container">
        <div class="col-1">
            <SearchBar />
        </div>
        <div class="col-2">
            <EmailContent />
        </div>
    </div>
</template>

<style scoped>
.home_container {
    /* background-color: #020202; */
    display: flex;
    height: 100vh;
}

.col-1 {
    width: 35%;
    background-color: #F1F1F1;
    overflow: auto;
  }

  .col-2 {
    width: 65%;
    /* background-color: #E0E0E0; */
  }
</style>
