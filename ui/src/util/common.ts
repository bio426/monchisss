export default {
    capitalizePhrase(phrase: string) {
        let arr = phrase.split(" ")
        arr = arr.map(str => str.charAt(0).toUpperCase() + str.slice(1))
        return arr.join(" ")
    }
}
