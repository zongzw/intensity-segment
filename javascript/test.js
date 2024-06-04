let m = new Map()

m[0] = "a"
m[10] = "b"

m.set(20, "c")

console.log(m.keys())
console.log(Array.from(m.keys()))

let ks = Array.from(m.keys())

ks.forEach(k => {
    console.log(k)
})
for (let k in ks.forEach) {
    console.log(k)
}