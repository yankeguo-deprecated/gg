const _ = require("lodash");

const DATA = [];

for (let i = 0; i < 5; i++) {
    for (let j = 0; j < 5; j++) {
        const OBJ = {
            i,
            j,
            G: "",
            GT: "",
            I: "",
            IV: "",
            O: "",
        };

        const ranges_i = _.range(1, i + 1);
        const ranges_j = _.range(1, j + 1);

        const G = [];
        const GT = [];

        for (const si of ranges_i) {
            G.push("I" + si + " any");
            GT.push("I" + si)
        }
        for (const sj of ranges_j) {
            G.push("O" + sj + " any");
            GT.push("O" + sj)
        }

        OBJ.G = G.join(", ");
        OBJ.GT = GT.join(", ");
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