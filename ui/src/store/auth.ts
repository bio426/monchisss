import { ref, computed } from "vue"
import { defineStore } from "pinia"

import * as tAuth from "@/type/auth"

const STORAGE_USER_KEY = "monchisss-user"

export default defineStore("auth", () => {
    const user = ref<tAuth.User>()
    const authenticated = computed(() => user.value != undefined)

    function setUserFromStorage() {
        const strUser = localStorage.getItem(STORAGE_USER_KEY)
        if (strUser == null) return
        const eUser = JSON.parse(strUser) as tAuth.ExpirableUser
        const expiryDate = new Date(eUser.expiryDate)
        if (new Date() > expiryDate) {
            localStorage.removeItem(STORAGE_USER_KEY)
            return
        }
        user.value = eUser
    }
    setUserFromStorage()

    function setUser(data: tAuth.ExpirableUser) {
        user.value = data
        localStorage.setItem(STORAGE_USER_KEY, JSON.stringify(data))
    }

    function clearUser() {
        user.value = undefined
        localStorage.removeItem(STORAGE_USER_KEY)
    }

    return { user, authenticated, setUser, clearUser }
})
