let input = document.getElementById('file');
// let epochCounter = document.getElementById('epochCounter');
let current_idx = 0;
let cur_log = document.getElementById('current');
let canvas = document.getElementById('canvas');
let context = canvas.getContext('2d');
let world = [];
let world_static = [];
let epoches = [];
let epoch_id = document.getElementById('epoch_id');

context.canvas.width = 700;
context.canvas.height = 500;
context.scale(10, 10);

function startLive(index, render_vector = -1) {
    let current_world = epoches[epoch_id.value][1];
    current_world.forEach(function (item) {
        if (item[index] === void 0) {
            return;
        }
        let last_current;
        let x;
        let y;
        let xy;
        if (item[index + render_vector] !== void 0) {
            last_current = item[index + render_vector].split(' ');
            if (last_current[3] !== void 0) {
                xy = last_current[3].split(',');
                x = Number.parseInt(xy[0]);
                y = Number.parseInt(xy[1]);
                context.clearRect(x, y, 1, 1);
            }
        }


        let current = item[index].split(' ');
        if (current[3] === void 0) {
            return;
        }
        xy = current[3].split(',');
        x = Number.parseInt(xy[0]);
        y = Number.parseInt(xy[1]);
        context.fillRect(x, y, 1, 1);
        if (current[0] === 'L') {
            context.fillStyle = 'yellow';
        }
        context.fill();

    });
    return index;
}

function staticRender() {
    let stat = epoches[epoch_id.value][0];
    stat.forEach(function (item) {
        let current = item.split(' ');
        let xy = current[2].split(',');
        let x = Number.parseInt(xy[0]);
        let y = Number.parseInt(xy[1]);
        context.fillRect(x, y, 1, 1);
        if (current[1] === 'W') {
            context.fillStyle = 'grey';
        }
        if (current[1] === 'P') {
            context.fillStyle = 'red';
        }
        if (current[1] === 'E') {
            context.fillStyle = 'green';
        }
        context.fill();

    });
}

function processLive(live) {
    let file = live.split("\n");
    file.forEach(function (str) {
        let current = str.split(' ');
        let e_id;
        if (current[0] === 'Epoh') {
            e_id = parseInt(current[1]);
            epoches[e_id] = [];
            world = [];
            world_static = [];
            return;
        }
        if (current[0] === 'STATIC') {
            world_static.push(str);
            return;
        }
        if (world[current[1]] !== void 0) {
            world[current[1]].push(str);
        } else {
            world[current[1]] = [str];
        }
        epoches[e_id][0] = world_static;
        epoches[e_id][1] = world;
    });
    staticRender();
    startLive(0);
}

let lastBtn = document.getElementById('last');
let nextBtn = document.getElementById('next');
let autoBtn = document.getElementById('auto');
let stopBtn = document.getElementById('stop');
let lastEpochBtn = document.getElementById('lastEpoch');
let nextEpochBtn = document.getElementById('nextEpoch');

lastBtn.addEventListener('click', function () {
    if (current_idx - 1 < 0) {
        return;
    }
    current_idx = current_idx - 1;
    cur_log.innerHTML = 'Текущий кадр: ' + current_idx;
    startLive(current_idx, 1);
});

nextBtn.addEventListener('click', function () {
    if (current_idx + 1 > world.length - 1) {
        return;
    }
    current_idx = current_idx + 1;
    cur_log.innerHTML = 'Текущий кадр: ' + current_idx;
    startLive(current_idx, -1);
});

input.addEventListener('change', function (evt) {
    let {files: file} = evt.target;
    let reader = new FileReader();
    reader.onload = function (f) {
        let {result: text} = f.target;
        processLive(text);
    };
    reader.readAsText(file[0]);
});


let auto_render_id = -1;
autoBtn.addEventListener('click', function () {
    context.clearRect(0, 0, context.width, context.height);
    staticRender();
    let current_index = 0;
    auto_render_id = setInterval(function () {
        cur_log.innerHTML = 'Текущий кадр: ' + current_index;
        startLive(current_index);
        current_index += 1;

    }, 500);
});

stopBtn.addEventListener('click', function () {
    clearInterval(auto_render_id);
    context.clearRect(0, 0, canvas.width, canvas.height);
    staticRender();
    // let current_index = 0;
    startLive(0);
});

lastEpochBtn.addEventListener('click', function () {
    let e_id = parseInt(epoch_id.value);
    e_id--;
    if (e_id < 0) {
        return;
    }
    clearInterval(auto_render_id);
    context.clearRect(0, 0, canvas.width, canvas.height);
    // let current_index = 0;
    epoch_id.value = e_id;
    staticRender();
    startLive(0);
});

nextEpochBtn.addEventListener('click', function () {
    let e_id = parseInt(epoch_id.value);
    e_id++;
    if (epoches[e_id] === void 0) {
        return;
    }
    clearInterval(auto_render_id);
    context.clearRect(0, 0, canvas.width, canvas.height);
    // current_index = 0;
    epoch_id.value = e_id;
    staticRender();
    startLive(0);
});
