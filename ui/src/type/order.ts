export type Category = {
    name: string
    items: Item[]
}

export type ItemVariant = {
    id: number
    name: string
    price: number
}

// when type simple only has price
// type variant has variants
export type Item = {
    id: number
    type: string
    name: string
    price?: number
    variants?: ItemVariant[]
}

export type CartItem = {
    id: number
    name: string
    price?: number
    variant?: ItemVariant
}
