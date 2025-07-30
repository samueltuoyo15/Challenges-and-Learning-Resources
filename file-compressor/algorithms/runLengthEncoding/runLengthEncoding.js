

import fs from "fs"

const original = fs.readFileSync("input.txt", "utf8")
const compressed = runLengthEncoding(original)
fs.writeFileSync("compressed.rle", compressed)
console.log("Compression done!")

console.log("Original File Size:", getFileSizeKb("input.txt"))
console.log("Compressed File Size:", getFileSizeKb("compressed.rle"))

function runLengthEncoding(string){
    let encodedString = []
    let count = 1;

    for (let i = 0; i < string.length; i++){
        if(string[i] === string[i + 1] && count < 9){
            count++
        } else {
            encodedString.push(count + string[i])
            count = 1
        }
    }
    return encodedString.join("")
}

function decompressRunLengthEncoding(string){
    let decodedString = []

    for(let i = 0; i < string.length; i +=2) {
        const count = parseInt(string[i], 10)
        const char = string[i + 1]
        decodedString.push(char.repeat(count))
    }
    return decodedString.join("")
}

const compressedFile = fs.readFileSync("compressed.rle", "utf8")
const decompressed = decompressRunLengthEncoding(compressedFile)
fs.writeFileSync("decompressed.txt", decompressed)
console.log("Decompression done!")
console.log("Compressed File Size:", getFileSizeKb("compressed.rle"))
console.log("Decompressed File Size:", getFileSizeKb("input.txt"))



function getFileSizeKb(filePath) {
    const stats = fs.statSync(filePath)
    return (stats.size / 1024).toFixed(2) + " KB"
}