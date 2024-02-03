<script setup lang="ts">
import { computed, watch } from "vue"

const props = defineProps<{ total: number; count: number }>()
const model = defineModel<number>({ required: true })

function toPage(val: number) {
	model.value = val
}
const pageCount = computed(() => {
	return Math.ceil(props.total / props.count)
})
const isComplex = computed(() => pageCount.value > 8)
const inversed = computed(() => {
	let arr: number[] = []
	for (let i = pageCount.value; i > pageCount.value - 5; i--) {
		arr.push(i)
	}
	return arr.reverse()
})

watch(
	() => props.total,
	() => {
		if (model.value * props.count + 1 > props.total) toPage(1)
	}
)
</script>

<template>
	<div class="text-center">
		<div class="join">
			<template v-if="!isComplex">
				<button
					class="btn join-item"
					:class="{ 'btn-active': i == model }"
					v-for="i in pageCount"
					@click="toPage(i)"
				>
					{{ i }}
				</button>
			</template>
			<template v-else-if="isComplex && model < 5">
				<button
					class="btn join-item"
					:class="{ 'btn-active': i == model }"
					v-for="i in 5"
					@click="toPage(i)"
				>
					{{ i }}
				</button>
				<button class="btn btn-disabled join-item">...</button>
				<button class="btn join-item" @click="toPage(pageCount)">
					{{ pageCount }}
				</button>
			</template>
			<template v-else-if="isComplex && model > pageCount - 4">
				<button class="btn join-item" @click="toPage(1)">1</button>
				<button class="btn btn-disabled join-item">...</button>
				<button
					class="btn join-item"
					:class="{ 'btn-active': i == model }"
					v-for="i in inversed"
					@click="toPage(i)"
				>
					{{ i }}
				</button>
			</template>
			<template v-else>
				<button class="btn join-item" @click="toPage(1)">1</button>
				<button class="btn btn-disabled join-item">...</button>
				<button class="btn join-item" @click="toPage(model - 1)">
					{{ model - 1 }}
				</button>
				<button class="btn btn-active join-item">
					{{ model }}
				</button>
				<button class="btn join-item" @click="toPage(model + 1)">
					{{ model + 1 }}
				</button>
				<button class="btn btn-disabled join-item">...</button>
				<button class="btn join-item" @click="toPage(pageCount)">
					{{ pageCount }}
				</button>
			</template>
		</div>
	</div>
</template>
