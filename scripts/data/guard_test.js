const _ = require("lodash");

const DATA = [];

for (let i = 2; i < 8; i++) {
    const ranges = _.range(1, i + 1);

    DATA.push({
        i,
        j: ranges,
        J: ranges.join(", "),
        T: _.map(ranges, function (j) {
            return "T" + j + " any";
        }).join(", "),
        A: _.map(ranges, function (j) {
            return "v" + j + " T" + j;
        }).join(", "),
        RT: _.map(ranges, function (j) {
            return "T" + j;
        }).join(", "),
        R: _.map(ranges, function (j) {
            return "v" + j;
        }).join(", "),
    });
}

module.exports = {
    data: DATA,
};