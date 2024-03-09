import base from "./base"

import * as tStore from "@/type/store"

const path = "store"

export default {
    async list() {
        const res = await base.get(path)

        return res.json<{ rows: any[] }>()
    },
    async create(body: { name: string, token: string, ownerUsername: string, ownerPassword: string }) {
        const res = await base.post(path, { json: body })

        return res.status
    },
    async listUsers(id: number) {
        const res = await base.get(`${path}/${id}/user`)

        return res.json<{ rows: tStore.User[] }>()
    },
    async createUser(id: number, body: { username: string, password: string, role: string }) {
        const res = await base.post(`${path}/${id}/user`, { json: body })

        return res.status
    },
}
