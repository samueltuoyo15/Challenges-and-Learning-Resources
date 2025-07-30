import fs from "fs";

// ========== Huffman Classes & Functions ==========

class Node {
    constructor(char, freq, left = null, right = null) {
        this.char = char;
        this.freq = freq;
        this.left = left;
        this.right = right;
    }
}

// Frequency counter
function getFrequencies(text) {
    const freq = {};
    for (const char of text) {
        freq[char] = (freq[char] || 0) + 1;
    }
    return freq;
}

// Tree builder
function buildTree(freq) {
    const nodes = Object.entries(freq).map(([char, freq]) => new Node(char, freq));
    while (nodes.length > 1) {
        nodes.sort((a, b) => a.freq - b.freq);
        const left = nodes.shift();
        const right = nodes.shift();
        const newNode = new Node(null, left.freq + right.freq, left, right);
        nodes.push(newNode);
    }
    return nodes[0]; // Root node
}

// Generate code table
function buildCodeTable(node, prefix = "", table = {}) {
    if (!node) return;
    if (node.char !== null) {
        table[node.char] = prefix;
    }
    buildCodeTable(node.left, prefix + "0", table);
    buildCodeTable(node.right, prefix + "1", table);
    return table;
}

// Encode
function huffmanEncode(text) {
    const freq = getFrequencies(text);
    const tree = buildTree(freq);
    const codeTable = buildCodeTable(tree);
    const encoded = text.split("").map(char => codeTable[char]).join("");
    return { encoded, tree };
}

// Decode
function huffmanDecode(encoded, tree) {
    let result = "";
    let node = tree;
    for (const bit of encoded) {
        node = bit === "0" ? node.left : node.right;
        if (node.char !== null) {
            result += node.char;
            node = tree;
        }
    }
    return result;
}

// JSON-safe serialization
function treeToJson(node) {
    if (!node) return null;
    return {
        char: node.char,
        freq: node.freq,
        left: treeToJson(node.left),
        right: treeToJson(node.right)
    };
}

// JSON-safe deserialization
function jsonToTree(obj) {
    if (!obj) return null;
    return new Node(
        obj.char,
        obj.freq,
        jsonToTree(obj.left),
        jsonToTree(obj.right)
    );
}

// Get file size
function getFileSizeKb(filePath) {
    const stats = fs.statSync(filePath);
    return (stats.size / 1024).toFixed(2) + " KB";
}

// ========== Main Workflow ==========

// Step 1: Read original file
const original = fs.readFileSync("input.txt", "utf8");

// Step 2: Compress
const { encoded, tree } = huffmanEncode(original);
fs.writeFileSync("compressed.huff", encoded);
fs.writeFileSync("tree.json", JSON.stringify(treeToJson(tree)));
console.log("Compression done!");

// Step 3: Show file sizes
console.log("Original File Size:", getFileSizeKb("input.txt"));
console.log("Compressed File Size:", getFileSizeKb("compressed.huff"));

// Step 4: Decompress
const encodedData = fs.readFileSync("compressed.huff", "utf8");
const huffmanTree = jsonToTree(JSON.parse(fs.readFileSync("tree.json", "utf8")));
const decompressed = huffmanDecode(encodedData, huffmanTree);
fs.writeFileSync("decompressed.txt", decompressed);
console.log("Decompression done!");

// Step 5: Show decompressed file size
console.log("Decompressed File Size:", getFileSizeKb("decompressed.txt"));
