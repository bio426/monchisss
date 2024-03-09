import { HTTPError } from "ky";
import base from "./base"

import * as tAuth from "@/type/auth"

const path = "auth"

export default {
    async login(body: { username: string; password: string }) {
        const res = await base.post(path + "/login", { json: body })

        return res.json<tAuth.ExpirableUser>()
    },
    async logout() {
        const res = await base.post(path + "/logout")

        return res.status
    },
}
