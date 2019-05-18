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
            let i = 0;
            for (const gen of genom) {
                const gStr = resolveGenom(gen);
                if (gStr != "") {
                    let genDOM = document.createElement("span");
                    genDOM.className = "gen";
                    genDOM.innerText = `${i}_${gStr}`;
                    genomDOM.appendChild(genDOM);
                }
                i++;
            }
            this.genomPlace.appendChild(genomDOM);
        }
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
        case 17:
            return 'GEatUp';
        case 18:
            return 'GEatUpLeft';
        case 19:
            return 'GEatUpRight';
        case 20:
            return 'GEatLeft';
        case 21:
            return 'GEatRight';
        case 22:
            return 'GEatDown';
        case 23:
            return 'GEatDownLeft';
        case 24:
            return 'GEatDownRight';
        // case 25:
        //     return 'GAttackUp';
        // case 26:
        //     return 'GAttackUpLeft';
        // case 27:
        //     return 'GAttackUpRight';
        // case 28:
        //     return 'GAttackLeft';
        // case 29:
        //     return 'GAttackRight';
        // case 30:
        //     return 'GAttackDown';
        // case 31:
        //     return 'GAttackDownLeft';
        // case 32:
        //     return 'GAttackDownRight';
        case 25:
            return 'GRecycleUp';
        case 26:
            return 'GRecycleUpLeft';
        case 27:
            return 'GRecycleUpRight';
        case 28:
            return 'GRecycleLeft';
        case 29:
            return 'GRecycleRight';
        case 30:
            return 'GRecycleDown';
        case 31:
            return 'GRecycleDownLeft';
        case 32:
            return 'GRecycleDownRight';
        default:
            if (g >= 34 && g <= 66) {
                return `GJumpTo_${g - 34}`;
            }
            return "";
    }
}
class Player {
    constructor(played = false) {
        this.played = played;
        this.epoches = [];
        this.render = new Render();
        this.epochCounter = document.getElementById('epochCounter');
        this.epoch_id = document.getElementById('epoch_id');
        this.cur_log = document.getElementById('current');
        this.speed = document.getElementById('speed');
        this.last = document.getElementById('last');
        this.next = document.getElementById('next');
        this.eph_target = 0;
        this.shot_target = 0;
        this.epoch_id.onchange = (e) => {
            if (e.target != null) {
                // @ts-ignore
                this.eph_target = parseInt(e.target.value);
                if (this.eph_target < 0) {
                    this.eph_target = 0;
                }
                if (this.eph_target > this.epoches.length - 1) {
                    this.eph_target = 0;
                }
            }
        };
    }
    disable() {
        this.epoch_id.disabled = true;
        this.next.disabled = true;
        this.last.disabled = true;
    }
    enable() {
        this.epoch_id.disabled = false;
        this.next.disabled = false;
        this.last.disabled = false;
    }
    stop() {
        this.played = false;
        this.enable();
    }
    shot(ss) {
        for (const s of ss) {
            s.render(this.render.context);
        }
        this.cur_log.innerText =
            `Кадр: ${this.shot_target + 1}:${this.epoches[this.eph_target].snapshot.length}`;
    }
    prev_shot() {
        let eh = this.epoches[this.eph_target];
        this.shot_target--;
        if (this.shot_target < 0) {
            this.shot_target = 0;
            // prev_eph
            return;
        }
        // let ps = eh.snapshot[this.shot_target - 1];
        // switch (ps.type) {
        //     case ActionType.A_EAT:
        //
        // }
        this.shot(eh.snapshot[this.shot_target]);
    }
    next_shot() {
        let eh = this.epoches[this.eph_target];
        this.shot_target++;
        if (this.shot_target > eh.snapshot.length - 1) {
            // next_eph
            return;
        }
        this.shot(eh.snapshot[this.shot_target]);
    }
    play() {
        return __awaiter(this, void 0, void 0, function* () {
            this.played = true;
            this.disable();
            for (; this.eph_target < this.epoches.length; this.eph_target++) {
                if (!this.played) {
                    this.enable();
                    return;
                }
                let eh = this.epoches[this.eph_target];
                this.epochCounter.innerText = `Эпоха: ${this.eph_target + 1}:${this.epoches.length}`;
                this.epoch_id.value = this.eph_target.toString();
                eh.pre_render();
                for (; this.shot_target < eh.snapshot.length; this.shot_target++) {
                    if (!this.played) {
                        this.enable();
                        return;
                    }
                    this.shot(eh.snapshot[this.shot_target]);
                    yield sleep(parseInt(this.speed.value));
                }
                this.shot_target = 0;
            }
            this.played = false;
            this.enable();
        });
    }
    load(log) {
        let e_id = 0;
        let shot = [];
        let lines = log.split("\n");
        lines.forEach((line) => {
            let current = line.split(' ');
            switch (current[0]) {
                case "EPOCH": {
                    // Инициализация эпохи
                    e_id = parseInt(current[1]);
                    this.epoches[e_id] = new Epoch(this.render);
                    shot = [];
                    return;
                }
                case "I": {
                    const e = this.epoches[e_id];
                    let position = current[2].split(',');
                    e.init(resolveType(current[1]), new Position(parseInt(position[0]), parseInt(position[1])));
                    return;
                }
                case "SHOT_END": {
                    const e = this.epoches[e_id];
                    e.last_shot_append(shot);
                    shot = [];
                    return;
                }
                case "GENOM": {
                    const e = this.epoches[e_id];
                    e.genom.push(JSON.parse(current.slice(2).join(",")));
                    return;
                }
                case "MUTATION": {
                    return;
                }
                case "S": {
                    let action = resolveAction(current[1]);
                    let cmd = [...current.slice(2)];
                    if (action == ActionType.A_UNKNOWN) {
                        return;
                    }
                    switch (action) {
                        case ActionType.A_DEI: {
                            let position = cmd[0].split(',');
                            shot.push(new SnapshotElement(action, new Position(parseInt(position[0]), parseInt(position[1]))));
                            break;
                        }
                        case ActionType.A_EAT:
                        case ActionType.A_MOVE: {
                            let on = cmd[0].split(',');
                            let to = cmd[1].split(',');
                            shot.push(new SnapshotElement(action, new Position(parseInt(on[0]), parseInt(on[1])), new Position(parseInt(to[0]), parseInt(to[1]))));
                            break;
                        }
                    }
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
    constructor(type, p1, p2) {
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
class Epoch {
    constructor(render) {
        this.render = render;
        this.zero = [];
        this.snapshot = [];
        this.genom = [];
    }
    last_shot_append(s) {
        if (s.length > 0) {
            this.snapshot.push(s);
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
const stopBtn = document.getElementById('stop');
if (stopBtn != null) {
    stopBtn.addEventListener('click', () => {
        player.stop();
    });
}
const nextBtn = document.getElementById('next');
if (nextBtn != null) {
    nextBtn.addEventListener('click', () => {
        player.next_shot();
    });
}
const lastBtn = document.getElementById('last');
if (lastBtn != null) {
    lastBtn.addEventListener('click', () => {
        player.prev_shot();
    });
}
//# sourceMappingURL=engine.js.map