<script setup lang="ts">
import { } from "vue"
</script>

<template>
    <div class="relative min-h-screen w-full">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8">Not Found</h1>
            <router-link class="btn" :to="{ name: 'dashboard' }">To Dashboard</router-link>
        </div>
    </div>
</template>
