<script setup lang="ts">
import { computed, nextTick, reactive, ref, watch } from "vue"
import { useRoute } from "vue-router";
import { useBreakpoints, breakpointsTailwind } from "@vueuse/core"
import { ShoppingCartIcon, ArrowUturnLeftIcon } from "@heroicons/vue/24/solid";

import * as tOrder from "@/type/order"
import orderService from "@/service/order"
import useToast from "@/composables/useToast";
import util from "@/util/common"
import Overlay from "@/components/Overlay.vue";
import ItemCard from "@/components/order/ItemCard.vue";
import CartItem from "@/components/order/CartItem.vue";
import ItemVariantModal from "@/components/order/ItemVariantModal.vue";
import ItemCompoundModal from "@/components/order/ItemCompoundModal.vue";
import { HTTPError } from "ky";

const route = useRoute()

const fixedTop = ref<HTMLDivElement>()
const fixedBottom = ref<HTMLDivElement>()
const spacerHeight = reactive({ top: "", bottom: "" })

// Llamar esta funcion cada vez que cambie el contenido de la lista de categorias
async function calculateSpacerHeight() {
    await nextTick()
    if (fixedTop.value == undefined) return
    const topHeight = fixedTop.value.clientHeight
    spacerHeight.top = (topHeight) + "px"

    if (fixedBottom.value == undefined) return
    const bottomHeight = fixedBottom.value.clientHeight
    spacerHeight.bottom = (bottomHeight) + "px"
}

// scrollspy de categorias
const categorieDivs = ref<HTMLDivElement[]>([])
function scrollToCategory(category: tOrder.Category) {
    const idx = categories.value.findIndex(c => c.name == category.name)
    if (idx == -1) return
    const el = categorieDivs.value[idx]
    const elPos = el.getBoundingClientRect().top
    const top = elPos + window.scrollY - 150
    console.log(top)
    window.scrollTo({ behavior: "smooth", top })
}
// ~

const showMobileResume = ref(false)
watch(showMobileResume, () => calculateSpacerHeight())

const { lg } = useBreakpoints(breakpointsTailwind)

// desktop fixed order
const desktopResumeHeight = computed(() => {
    const resHeight = window.innerHeight - parseFloat(spacerHeight.top)
    return resHeight + "px"
})

const loading = ref(true)

const storeInfo = reactive({ name: "", image: "" })
const categories = ref<tOrder.Category[]>([])
async function getMenu() {
    loading.value = true
    const orderId = route.params.id as string
    try {
        const res = await orderService.getMenu(orderId)
        storeInfo.name = res.storeName
        storeInfo.image = res.storeImage
        categories.value = res.categories
        loading.value = false
    } catch (err) {
        if (!(err instanceof HTTPError)) throw err
        const toast = useToast()
        if (err.response.status == 410) {
            toast.display({ message: "Invalid order id", variant: "error" })
        }
    }
    calculateSpacerHeight()
}
// calculateSpacerHeight()
getMenu()

const itemOptions = reactive<{ show: boolean, info: tOrder.Item }>({
    show: false,
    info: { id: 0, type: "", name: "" }
})
function openVariantSelector(item: tOrder.Item) {
    itemOptions.info = item
    itemOptions.show = true
}

// order state
const cart = ref<tOrder.CartItem[]>([])
const cartPrice = computed(() => cart.value.reduce((pre, cur) => {
    if (cur.price != undefined) {
        return pre + cur.price
    } else if (cur.variant != undefined) {
        return pre + cur.variant.price
    }
    return 0
}, 0))

function addToCart(item: tOrder.CartItem) {
    cart.value.push(item)
}
</script>

<template>
    <div class="relative min-h-screen">
        <!-- Header -->
        <div class="fixed top-0 left-0 w-full z-10" ref="fixedTop">
            <div class="bg-base-100 shadow-lg">
                <div class="navbar">
                    <div class="navbar-start">
                        <img class="w-10" :src="storeInfo.image" />
                        <h2 class="ml-4 font-bold text-xl">{{ util.capitalizePhrase(storeInfo.name) }}</h2>
                    </div>
                    <div class="navbar-end" v-if="!lg">
                        <button class="btn btn-outline btn-circle" @click="showMobileResume = !showMobileResume"
                            :disabled="loading">
                            <div class="indicator" v-if="!showMobileResume">
                                <ShoppingCartIcon class="w-6 h-6" />
                                <span class="badge badge-sm badge-primary indicator-item">9</span>
                            </div>
                            <ArrowUturnLeftIcon class="w-6 h-6" v-else />
                        </button>
                    </div>
                </div>
                <div class="flex flex-nowrap overflow-auto bg-neutral px-4 gap-8" v-if="!showMobileResume">
                    <button class="py-4 text-xl whitespace-nowrap" v-for="category in categories"
                        @click="scrollToCategory(category)">
                        {{ util.capitalizePhrase(category.name) }}
                    </button>
                </div>
            </div>
        </div>
        <div :style="{ height: spacerHeight.top }"></div>
        <Overlay :show="loading">
            <!-- Mobile -->
            <div class="w-11/12 mx-auto py-8" v-if="!lg">
                <div class="grid gap-8" v-if="!showMobileResume">
                    <div ref="categorieDivs" v-for="category in categories">
                        <h4 class="text-xl font-bold">{{ util.capitalizePhrase(category.name) }}</h4>
                        <hr class="border-primary my-2">
                        <div class="grid grid-cols-2 sm:grid-cols-3 gap-4">
                            <ItemCard :item="item" v-for="item in category.items"
                                @showVariantSelector="openVariantSelector" />
                        </div>
                    </div>
                </div>
                <div v-else>
                    <h2 class="mb-8 text-xl font-bold text-center">Order Resume</h2>
                    <div class="flex flex-col gap-4 ">
                        <CartItem :item v-for="item in cart" />
                    </div>
                </div>
            </div>
            <!-- Desktop -->
            <div class="grid grid-cols-4" v-else>
                <div class="col-span-3 p-4 py-8">
                    <div v-for="category in categories">
                        <h4 class="text-xl font-bold ">{{ category.name }}</h4>
                        <hr class="border-primary my-2">
                        <div class="grid grid-cols-2 sm:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 gap-4">
                            <ItemCard :item v-for="item in category.items" @selected="addToCart"
                                @showVariantSelector="openVariantSelector" />
                        </div>
                    </div>
                </div>
                <div>
                    <div class="fixed right-0 flex flex-col w-1/4 bg-neutral"
                        :style="{ top: spacerHeight.top, height: desktopResumeHeight }">
                        <h2 class="py-4 text-xl font-bold text-center">Order Resume</h2>
                        <div class="flex-grow flex flex-col gap-4 p-4 overflow-auto">
                            <CartItem :item v-for="item in cart" />
                        </div>
                        <div class="w-full p-4 bg-base-100 border border-neutral">
                            <button class="btn btn-block btn-info" :disabled="cart.length == 0">
                                <span>
                                    Finalizar pedido
                                </span>
                                <span v-if="cartPrice > 0">
                                    S/. {{ cartPrice.toFixed(2) }}
                                </span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </Overlay>
        <template v-if="!lg">
            <div :style="{ height: spacerHeight.bottom }"></div>
            <div class="fixed bottom-0 left-0 w-full bg-neutral p-4" ref="fixedBottom">
                <button class="btn btn-block btn-info" :disabled="loading">
                    <span>
                        Finalizar pedido
                    </span>
                    <span>
                        S/. 99.99
                    </span>
                </button>
            </div>
        </template>
        <ItemVariantModal :show="itemOptions.show" :item="itemOptions.info" @closing="itemOptions.show = false"
            @selected="addToCart" />
        <!-- <ItemCompoundModal :show="itemOptions.show" :item="itemOptions.info" @closing="itemOptions.show = false" -->
        <!--     @selected="addToCart" /> -->
    </div>
</template>
