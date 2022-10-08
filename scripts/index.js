const fs = require('fs-extra');
const handlebars = require('handlebars');
const path = require('path');
const {exec} = require('node:child_process');

const beginOfGeneration = "__BEGIN_OF_GENERATION__";

const names = [
    "guard",
    "guard_test",
    "types",
    "types_test",
];

async function render(name) {
    const tmpl = await fs.readFile(path.join(__dirname, "data", name + ".hbs"), 'utf-8');
    const data = require("./data/" + name);
    return handlebars.compile(tmpl)(data);
}

try {

    (async function () {
        for (const name of names) {
            const file = path.join(__dirname, "..", name + ".go");
            const data = await render(name);

            const src = await fs.readFile(file, 'utf-8');
            let lines = src.split("\n");
            for (let i = 0; i < lines.length; i++) {
                if (lines[i].includes(beginOfGeneration)) {
                    lines = lines.slice(0, i);
                    break
                }
            }

            const content = (lines.join("\n") + "\n// " + beginOfGeneration + "\n" + data);
            await fs.writeFile(file, content);
        }


        await new Promise(function (resolve, reject) {
            const cmd = exec("go fmt ./...", {
                cwd: path.join(__dirname, "..")
            });

            cmd.on('error', function (err) {
                console.log(err)
            });

            cmd.on('close', function (code) {
                if (code === 0) {
                    resolve(null)
                } else {
                    reject("code: "+code)
                }
            });
        })

    })();
} catch (e) {
    console.log(e)
}