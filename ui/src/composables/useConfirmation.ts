import { ref, reactive, watch } from "vue"

const active = ref(false)
const acceptTrigger = ref(Symbol())
const rejectTrigger = ref(Symbol())
const modalContent = reactive({
	title: "",
	message: "",
})

function display(opts: { title?: string; message?: string }): Promise<boolean> {
	modalContent.title = opts.title || "Are you shure to realize this action?"
	modalContent.message =
		opts.message || "You would not be able to revert this!"
	active.value = true

	return new Promise((resolve) => {
		watch(acceptTrigger, () => {
			resolve(true)
			active.value = false
		})
		watch(rejectTrigger, () => {
			resolve(false)
			active.value = false
		})
	})
}

export default function () {
	return {
		active,
		acceptTrigger,
		rejectTrigger,
		modalContent,
		display,
	}
}
