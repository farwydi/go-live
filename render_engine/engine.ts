class Render {
    public canvas: HTMLCanvasElement = <HTMLCanvasElement>document.getElementById('canvas');
    public context: CanvasRenderingContext2D = <CanvasRenderingContext2D>this.canvas.getContext("2d");

    constructor() {
        this.canvas.width = 700;
        this.canvas.height = 500;
        this.context.scale(10, 10);
    }

    clear() {
        // context.clearRect(x, y, 1, 1);
    }

    print(x: number, y: number) {
        this.context.fillRect(x, y, 1, 1);
        this.context.fillStyle = 'yellow';
        this.context.fill();
    }
}


class Player {
    epoches: Epoch[] = [];

    render: Render = new Render();

    epochCounter: HTMLInputElement = <HTMLInputElement>document.getElementById('epochCounter');

    async play() {
        for (const eh of this.epoches) {
            this.epochCounter.innerText = `Эпоха: ${eh.id}`;

            eh.pre_render();
            await eh.play();
        }
    }

    load(log: string) {
        let e_id = 0;
        let lines = log.split("\n");
        lines.forEach((line: string) => {
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

                    case "S": {
                        const e = this.epoches[e_id];
                        let position = current[2].split(',');
                        e.append(resolveAction(current[1]), [...current.slice(2)]);
                        return;
                    }

                }
            }
        );

        console.log(this.epoches);
    }
}

class Position {
    constructor(public x: number, public y: number) {

    }
}

enum ActionType {
    A_UNKNOWN = 0,
    A_DEI = 1,
    A_EAT = 2,
    A_MOVE = 4,
}

function resolveAction(t: string): ActionType {
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

enum ElementType {
    E_UNKNOWN = 0,
    E_WELL = 1,
    E_EAT = 2,
    E_POISON = 3,
    E_LIVE = 4,
    E_EMPTY = 5,
}

function resolveType(t: string): ElementType {
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
    constructor(public type: ElementType, public position: Position) {

    }

    render(ctx: CanvasRenderingContext2D) {
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
    constructor(public id: string, public type: ActionType, public p1: Position, public p2?: Position) {

    }

    render(ctx: CanvasRenderingContext2D) {
        switch (this.type) {
            case ActionType.A_MOVE:
            case ActionType.A_EAT: {
                ctx.clearRect(this.p1.x, this.p1.y, 1, 1);

                if (this.p2 == null) {
                    return
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

function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function resolveXY(x: number, y: number): number {
    return (32 * x) + y
}

class Epoch {
    public zero: BaseElement[] = [];
    public snapshot: SnapshotElement[] = [];

    cur_log: HTMLInputElement = <HTMLInputElement>document.getElementById('current');

    constructor(private render: Render, public id: string) {

    }

    get count(): number {
        return this.snapshot.length
    }

    speed: HTMLInputElement = <HTMLInputElement>document.getElementById('speed');

    async play() {
        for (const s of this.snapshot) {
            s.render(this.render.context);

            this.cur_log.innerText = `Кадр: ${s.id}`;

            await sleep(parseInt(this.speed.value));
        }
    }

    append(action: ActionType, cmd: string[]) {
        if (action == ActionType.A_UNKNOWN) {
            return
        }

        switch (action) {
            case ActionType.A_DEI: {
                let position = cmd[0].split(',');
                this.snapshot.push(new SnapshotElement(
                    this.snapshot.length.toString(),
                    action,
                    new Position(parseInt(position[0]), parseInt(position[1])),
                ));
                break;
            }
            case ActionType.A_EAT:
            case ActionType.A_MOVE: {
                let on = cmd[0].split(',');
                let to = cmd[1].split(',');
                this.snapshot.push(new SnapshotElement(
                    this.snapshot.length.toString(),
                    action,
                    new Position(parseInt(on[0]), parseInt(on[1])),
                    new Position(parseInt(to[0]), parseInt(to[1]))
                ));
                break;
            }
        }
    }

    init(type: ElementType, position: Position) {
        this.zero.push(new BaseElement(
            type,
            position,
        ))
    }

    pre_render() {
        this.render.context.clearRect(0, 0, 700, 500);

        for (const b of this.zero) {
            b.render(this.render.context)
        }
    }
}

const player = new Player();

const inputFile = document.getElementById("file");
if (inputFile != null) {
    inputFile.addEventListener("change", (event: Event) => {
        // @ts-ignore
        let {files: file} = event.target;
        let reader = new FileReader();
        reader.onload = (pe: ProgressEvent) => {
            // @ts-ignore
            let {result: text}: string = pe.target;
            player.load(text);
        };
        reader.readAsText(file[0]);
    });
}

const autoBtn = document.getElementById('auto');
if (autoBtn != null) {
    autoBtn.addEventListener('click', async () => {
        await player.play()
    });
}