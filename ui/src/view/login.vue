<script setup lang="ts">
import { reactive, ref } from "vue"
import { useRouter } from "vue-router"
import { HTTPError } from "ky"

import authService from "@/service/auth"
import useAuthStore from "@/store/auth"
import useToast from "@/composables/useToast"
import Overlay from "@/components/Overlay.vue"

const router = useRouter()
const authStore = useAuthStore()

const form = reactive({ username: "", password: "" })
const loading = ref(false)

async function login() {
    loading.value = true
    try {
        const user = await authService.login({
            username: form.username,
            password: form.password,
        })
        authStore.setUser(user)
        router.push({ name: "dashboard" })
    } catch (err) {
        if (!(err instanceof HTTPError)) throw err
        const toast = useToast()
        switch (err.response.status) {
            case 401:
                toast.display({ message: "Inexistent user", variant: "error" })
                break;
            case 410:
                toast.display({ message: "Inactive user", variant: "error" })
                break;
        }
        form.username = ""
        form.password = ""
    }
    loading.value = false
}
</script>

<template>
    <div class="relative w-screen min-h-screen ">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8 text-2xl font-bold text-center">Login</h1>
            <Overlay :show="loading">
                <form @submit.prevent="login">
                    <label class="form-control w-full mb-4">
                        <div class="label">
                            <span class="label-text">Username</span>
                        </div>
                        <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                    </label>
                    <label class="form-control w-full mb-8">
                        <div class="label">
                            <span class="label-text">Password</span>
                        </div>
                        <input class="input input-bordered w-full" type="password" required v-model="form.password" />
                    </label>
                    <button class="btn btn-block btn-primary" type="submit">
                        Log In
                    </button>
                </form>
            </Overlay>
        </div>
    </div>
</template>
