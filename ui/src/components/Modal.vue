<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref, watch } from "vue"
import { XMarkIcon } from "@heroicons/vue/24/solid"

const props = defineProps<{ show: boolean; title: string }>()
const emit = defineEmits<{ closing: [] }>()

const dialogEl = ref<HTMLDialogElement>()
watch(
    () => props.show,
    (v) => {
        if (v) {
            dialogEl.value?.showModal()
        } else {
            dialogEl.value?.close()
        }
    }
)
function emitClosing() {
    emit("closing")
}
onMounted(() => {
    if (!dialogEl.value) return
    dialogEl.value.addEventListener("close", emitClosing)
})
onBeforeUnmount(() => {
    if (!dialogEl.value) return
    dialogEl.value.removeEventListener("close", emitClosing)
})
</script>

<template>
    <Teleport to="#modal">
        <dialog class="modal" ref="dialogEl">
            <div class="modal-box">
                <form method="dialog">
                    <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2" @click="dialogEl?.close()">
                        <XMarkIcon class="w-6" />
                    </button>
                </form>
                <h3 class="font-bold text-lg mb-4">{{ title }}</h3>
                <slot />
            </div>
        </dialog>
    </Teleport>
</template>
