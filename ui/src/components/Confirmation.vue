<script setup lang="ts">
import { } from "vue"

import useConfirmation from "../composables/useConfirmation"

const { active, acceptTrigger, rejectTrigger, modalContent } = useConfirmation()
const emit = defineEmits(["closing"])

function handleAccept() {
    acceptTrigger.value = Symbol()
}
function handleReject() {
    rejectTrigger.value = Symbol()
}
</script>

<template>
    <Teleport to="#confirmation">
        <input type="checkbox" id="confirmation-modal" class="modal-toggle" v-model="active" />
        <div class="modal">
            <div class="modal-box relative">
                <div>
                    <h3 class="text-lg text-center font-bold">
                        {{ modalContent.title }}
                    </h3>
                    <p class="py-4 text-center">
                        {{ modalContent.message }}
                    </p>
                    <div class="h-12"></div>
                    <div class="grid grid-cols-2 gap-4">
                        <button class="btn btn-success" @click="handleAccept">
                            Accept
                        </button>
                        <button class="btn btn-error" @click="handleReject">
                            Reject
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </Teleport>
</template>
