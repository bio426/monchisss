import { HTTPError } from "ky";
import base from "./base"

import * as tAuth from "@/type/auth"

const path = "auth"

export default {
    async login(body: { username: string; password: string }) {
        const res = await base.post(path + "/login", { json: body })
        try {

        } catch (err) {
            if (!(err instanceof HTTPError)) throw err
            switch (err.response.status) {
                case 422:
                    break;
                default:

            }
        }

        return res.json<tAuth.ExpirableUser>()
    },
    async logout() {
        const res = await base.post(path + "/logout")

        return res.status
    },
}
