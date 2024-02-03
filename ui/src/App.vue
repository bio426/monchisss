<script setup lang="ts">
import { } from "vue"

import Confirmation from "./components/Confirmation.vue";
import Toast from "./components/Toast.vue";
</script>

<template>
    <div>
        <router-view />
        <Confirmation />
        <Toast />
    </div>
</template>

