"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
class Render {
    constructor() {
        this.canvas = document.getElementById('canvas');
        this.context = this.canvas.getContext("2d");
        this.genomPlace = document.getElementById('genom');
        this.canvas.width = 640;
        this.canvas.height = 320;
        this.context.scale(10, 10);
    }
    render_genom(genoms) {
        this.genomPlace.innerHTML = "";
        for (const genom of genoms) {
            let genomDOM = document.createElement("div");
            genomDOM.className = "genom";
            for (const gen of genom) {
                const gStr = resolveGenom(gen);
                if (gStr != "x") {
                    let genDOM = document.createElement("span");
                    genDOM.className = "gen";
                    genDOM.innerText = gStr;
                    genomDOM.appendChild(genDOM);
                }
            }
            this.genomPlace.appendChild(genomDOM);
        }
    }
    clear() {
        // context.clearRect(x, y, 1, 1);
    }
    print(x, y) {
        this.context.fillRect(x, y, 1, 1);
        this.context.fillStyle = 'yellow';
        this.context.fill();
    }
}
function resolveGenom(g) {
    switch (g) {
        case 0:
            return 'GWait';
        case 1:
            return 'GMoveUp';
        case 2:
            return 'GMoveUpLeft';
        case 3:
            return 'GMoveUpRight';
        case 4:
            return 'GMoveLeft';
        case 5:
            return 'GMoveRight';
        case 6:
            return 'GMoveDown';
        case 7:
            return 'GMoveDownLeft';
        case 8:
            return 'GMoveDownRight';
        // Посмотреть
        case 9:
            return 'GSeeUp';
        case 10:
            return 'GSeeUpLeft';
        case 11:
            return 'GSeeUpRight';
        case 12:
            return 'GSeeLeft';
        case 13:
            return 'GSeeRight';
        case 14:
            return 'GSeeDown';
        case 15:
            return 'GSeeDownLeft';
        case 16:
            return 'GSeeDownRight';
        // Конец команд
        case 17:
            return 'GEnd';
        case 18:
            return 'GJumpStart';
        case 34:
            return 'GJumpEnd';
        default:
            return "x";
    }
}
class Player {
    constructor() {
        this.epoches = [];
        this.render = new Render();
        this.epochCounter = document.getElementById('epochCounter');
        this.epoch_id = document.getElementById('epoch_id');
    }
    play() {
        return __awaiter(this, void 0, void 0, function* () {
            for (const eh of this.epoches) {
                this.epochCounter.innerText = `Эпоха: ${eh.id}`;
                this.epoch_id.value = eh.id;
                eh.pre_render();
                yield eh.play();
            }
        });
    }
    load(log) {
        let e_id = 0;
        let lines = log.split("\n");
        lines.forEach((line) => {
            let current = line.split(' ');
            switch (current[0]) {
                case "EPOCH": {
                    // Инициализация эпохи
                    e_id = parseInt(current[1]);
                    this.epoches[e_id] = new Epoch(this.render, current[1]);
                    return;
                }
                case "I": {
                    const e = this.epoches[e_id];
                    let position = current[2].split(',');
                    e.init(resolveType(current[1]), new Position(parseInt(position[0]), parseInt(position[1])));
                    return;
                }
                case "GENOM": {
                    const e = this.epoches[e_id];
                    e.genom.push(JSON.parse(current.slice(2).join(",")));
                    return;
                }
                case "S": {
                    const e = this.epoches[e_id];
                    let position = current[2].split(',');
                    e.append(resolveAction(current[1]), [...current.slice(2)]);
                    return;
                }
            }
        });
        console.log(this.epoches);
    }
}
class Position {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }
}
var ActionType;
(function (ActionType) {
    ActionType[ActionType["A_UNKNOWN"] = 0] = "A_UNKNOWN";
    ActionType[ActionType["A_DEI"] = 1] = "A_DEI";
    ActionType[ActionType["A_EAT"] = 2] = "A_EAT";
    ActionType[ActionType["A_MOVE"] = 4] = "A_MOVE";
})(ActionType || (ActionType = {}));
function resolveAction(t) {
    switch (t) {
        case "M":
            return ActionType.A_MOVE;
        case "E":
            return ActionType.A_EAT;
        case "D":
            return ActionType.A_DEI;
        default:
            return ActionType.A_UNKNOWN;
    }
}
var ElementType;
(function (ElementType) {
    ElementType[ElementType["E_UNKNOWN"] = 0] = "E_UNKNOWN";
    ElementType[ElementType["E_WELL"] = 1] = "E_WELL";
    ElementType[ElementType["E_EAT"] = 2] = "E_EAT";
    ElementType[ElementType["E_POISON"] = 3] = "E_POISON";
    ElementType[ElementType["E_LIVE"] = 4] = "E_LIVE";
    ElementType[ElementType["E_EMPTY"] = 5] = "E_EMPTY";
})(ElementType || (ElementType = {}));
function resolveType(t) {
    switch (t) {
        case "W":
            return ElementType.E_WELL;
        case "E":
            return ElementType.E_EAT;
        case "P":
            return ElementType.E_POISON;
        case "L":
            return ElementType.E_LIVE;
        case "0":
            return ElementType.E_EMPTY;
        default:
            return ElementType.E_UNKNOWN;
    }
}
class BaseElement {
    constructor(type, position) {
        this.type = type;
        this.position = position;
    }
    render(ctx) {
        ctx.fillRect(this.position.x, this.position.y, 1, 1);
        switch (this.type) {
            case ElementType.E_EAT:
                ctx.fillStyle = 'red';
                break;
            case ElementType.E_LIVE:
                ctx.fillStyle = 'yellow';
                break;
            case ElementType.E_POISON:
                ctx.fillStyle = 'green';
                break;
            case ElementType.E_WELL:
                ctx.fillStyle = 'black';
                break;
        }
        ctx.fill();
    }
}
class SnapshotElement {
    constructor(id, type, p1, p2) {
        this.id = id;
        this.type = type;
        this.p1 = p1;
        this.p2 = p2;
    }
    render(ctx) {
        switch (this.type) {
            case ActionType.A_MOVE:
            case ActionType.A_EAT: {
                ctx.clearRect(this.p1.x, this.p1.y, 1, 1);
                if (this.p2 == null) {
                    return;
                }
                ctx.fillRect(this.p2.x, this.p2.y, 1, 1);
                ctx.fillStyle = 'yellow';
                ctx.fill();
                break;
            }
            case ActionType.A_DEI: {
                ctx.clearRect(this.p1.x, this.p1.y, 1, 1);
                break;
            }
        }
    }
}
function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
function resolveXY(x, y) {
    return (32 * x) + y;
}
class Epoch {
    constructor(render, id) {
        this.render = render;
        this.id = id;
        this.zero = [];
        this.snapshot = [];
        this.genom = [];
        this.cur_log = document.getElementById('current');
        this.speed = document.getElementById('speed');
    }
    get count() {
        return this.snapshot.length;
    }
    play() {
        return __awaiter(this, void 0, void 0, function* () {
            for (const s of this.snapshot) {
                s.render(this.render.context);
                this.cur_log.innerText = `Кадр: ${s.id}`;
                yield sleep(parseInt(this.speed.value));
            }
        });
    }
    append(action, cmd) {
        if (action == ActionType.A_UNKNOWN) {
            return;
        }
        switch (action) {
            case ActionType.A_DEI: {
                let position = cmd[0].split(',');
                this.snapshot.push(new SnapshotElement(this.snapshot.length.toString(), action, new Position(parseInt(position[0]), parseInt(position[1]))));
                break;
            }
            case ActionType.A_EAT:
            case ActionType.A_MOVE: {
                let on = cmd[0].split(',');
                let to = cmd[1].split(',');
                this.snapshot.push(new SnapshotElement(this.snapshot.length.toString(), action, new Position(parseInt(on[0]), parseInt(on[1])), new Position(parseInt(to[0]), parseInt(to[1]))));
                break;
            }
        }
    }
    init(type, position) {
        this.zero.push(new BaseElement(type, position));
    }
    pre_render() {
        this.render.context.clearRect(0, 0, 700, 500);
        this.render.render_genom(this.genom);
        for (const b of this.zero) {
            b.render(this.render.context);
        }
    }
}
const player = new Player();
const inputFile = document.getElementById("file");
if (inputFile != null) {
    inputFile.addEventListener("change", (event) => {
        // @ts-ignore
        let { files: file } = event.target;
        let reader = new FileReader();
        reader.onload = (pe) => {
            // @ts-ignore
            let { result: text } = pe.target;
            player.load(text);
        };
        reader.readAsText(file[0]);
    });
}
const nextEpochBtn = document.getElementById('nextEpoch');
if (nextEpochBtn != null) {
    nextEpochBtn.addEventListener('click', () => {
    });
}
const autoBtn = document.getElementById('auto');
if (autoBtn != null) {
    autoBtn.addEventListener('click', () => __awaiter(this, void 0, void 0, function* () {
        yield player.play();
    }));
}
//# sourceMappingURL=engine.js.map