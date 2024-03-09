import base from "./base"

import * as tOrder from "@/type/order"

const path = "order"

export default {
    async getMenu(orderId: string) {
        const res = await base.get(`${path}/${orderId}`)
        return res.json<{
            storeName: string
            storeImage: string
            categories: tOrder.Category[]
        }>()
    },

    async sendOrder(body: SendOrderPayload) {
        const res = await base.post(`${path}`)
        return res.status
    }

}

type OrderItem = {
    id: number,
    type: string,
    // if is a variant item
    variant?: number,
    // if is a compound item
    components?: number[]
}
type SendOrderPayload = {
    id: number,
    items: OrderItem[]
}
