export type Product = {
    type: string
    name: string
    price: number
    category: string
}

export type Category = {
    id: number
    name: string
}

export type ProductVariant = {
    id: number
    name: string,
    price: number
}

export type ProductVariantUnsaved = Omit<ProductVariant, "id">
