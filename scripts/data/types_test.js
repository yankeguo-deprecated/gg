const _ = require("lodash");

const TYPES = [
    "int",
    "int64",
    "int32",
    "bool",
    "float64",
    "string",
    "byte",
    "rune",
];

const DATA = [];

for (let i = 0; i < 5; i++) {
    for (let j = 0; j < 5; j++) {
        const OBJ = {
            i,
            j,
            G: "",
            GT: "",
            TEST_G: "",
            I: "",
            IV: "",
            O: "",
        };

        const ranges_i = _.range(1, i + 1);
        const ranges_j = _.range(1, j + 1);

        const G = [];
        const GT = [];
        const TEST_G = [];

        for (const si of ranges_i) {
            G.push("I" + si + " any");
            GT.push("I" + si)
            TEST_G.push(TYPES[TEST_G.length])
        }
        for (const sj of ranges_j) {
            G.push("O" + sj + " any");
            GT.push("O" + sj)
            TEST_G.push(TYPES[TEST_G.length])
        }

        OBJ.G = G.join(", ");
        OBJ.GT = GT.join(", ");
        OBJ.TEST_G = TEST_G.join(", ")
        OBJ.I = _.map(ranges_i, function (i) {
            return "i" + i + " I" + i
        }).join(", ")
        OBJ.IV = _.map(ranges_i, function (i) {
            return "i" + i
        }).join(", ")
        OBJ.O = _.map(ranges_j, function (j) {
            return "o" + j + " O" + j
        }).join(", ")


        DATA.push(OBJ);
    }
}

module.exports = {
    data: DATA,
};