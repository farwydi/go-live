var input 		   = document.getElementById('file');
var epochCounter   = document.getElementById('epochCounter');
var current_idx    =  0;
var cur_log        = document.getElementById('current');
var canvas 		   = document.getElementById('canvas');
var context        = canvas.getContext('2d');
var world 		   = [];
var world_static   = [];
var epoches 	   = [];
var epoch_id 	   = document.getElementById('epoch_id');

context.canvas.width  = 700;
context.canvas.height = 500;
context.scale(10,10);

function startLive(index, render_vector = -1) {
	current_world = epoches[epoch_id.value][1];
	current_world.forEach(function(item, idx) {
		if (item[index] === void 0 ) {
			return;
		}
		if (item[index + render_vector] !== void 0) {
			last_current = item[index + render_vector].split(' ');
			if (last_current[3] !== void 0) {
				xy = last_current[3].split(',');
				x = Number.parseInt(xy[0]);
				y = Number.parseInt(xy[1]);
				context.clearRect(x, y, 1, 1);
			}
		}
		

		current = item[index].split(' ');
		if (current[3] === void 0) {
			return;
		} 
		xy = current[3].split(',');
		x = Number.parseInt(xy[0]);
		y = Number.parseInt(xy[1]);
		context.fillRect(x,y,1,1);
		if (current[0] == 'L') {
			context.fillStyle = 'yellow';
		}
		context.fill();

	});
	return index;
}

function staticRender()
{
	static = epoches[epoch_id.value][0];
	static.forEach(function(item) {
		current = item.split(' ');
		xy = current[2].split(',');
		x = Number.parseInt(xy[0]);
		y = Number.parseInt(xy[1]);
		context.fillRect(x,y,1,1);
		if (current[1] == 'W') {
			context.fillStyle = 'grey';
		}
		if (current[1] == 'P') {
			context.fillStyle = 'red';
		}
		if (current[1] == 'E') {
			context.fillStyle = 'green';
		}
		context.fill();
			
	});
}

function processLive(live) {
	var file = live.split("\n");
	file.forEach(function(str,index) {
		current = str.split(' ');
		if (current[0] == 'Epoh') {
			 e_id = parseInt(current[1]);
			 epoches[e_id] = [];
			 world 		   = [];
		     world_static  = [];
			 return;
		}
		if (current[0] == 'STATIC') {
			world_static.push(str)
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

var lastBtn        = document.getElementById('last');
var nextBtn        = document.getElementById('next');
var autoBtn        = document.getElementById('auto');
var stopBtn        = document.getElementById('stop');
var lastEpochBtn   = document.getElementById('lastEpoch');
var nextEpochBtn   = document.getElementById('nextEpoch');

lastBtn.addEventListener('click', function() {
	if (current_idx -1 < 0) {
		return;
	}
	current_idx = current_idx - 1;
	cur_log.innerHTML = 'Текущий кадр: ' + current_idx;
	startLive(current_idx, 1);
});

nextBtn.addEventListener('click', function() {
	if (current_idx + 1 > world.legth - 1) {
		return;
	}
	current_idx = current_idx + 1;
	cur_log.innerHTML = 'Текущий кадр: ' + current_idx;
	startLive(current_idx, -1);
});

input.addEventListener('change',function(evt) {
	var file = evt.target.files;
	var reader = new FileReader();
	reader.onload = function(f) {
		text = f.target.result;
		processLive(text);
	}
	reader.readAsText(file[0]);
});


var auto_render_id = -1;
autoBtn.addEventListener('click', function() {
	context.clearRect(0, 0, context.width, context.height);
	staticRender();
	current_index = 0;
	auto_render_id = setInterval(function() {
		cur_log.innerHTML = 'Текущий кадр: ' + current_index;
		startLive(current_index);
		current_index += 1;
		
	},500);
});

stopBtn.addEventListener('click', function() {
	clearInterval(auto_render_id);
	context.clearRect(0, 0, canvas.width, canvas.height);
	staticRender();
	current_index = 0;
    startLive(0);
});

lastEpochBtn.addEventListener('click', function() {
	e_id = parseInt(epoch_id.value);
	e_id--;
	if (e_id < 0) {
		return;
	}
	clearInterval(auto_render_id);
	context.clearRect(0, 0, canvas.width, canvas.height);
	current_index = 0;
	epoch_id.value = e_id;
	staticRender();
	startLive(0);
});

nextEpochBtn.addEventListener('click', function() {
	e_id = parseInt(epoch_id.value);
	e_id++;
	if (epoches[e_id] === void 0) {
		return;
	}
	clearInterval(auto_render_id);
	context.clearRect(0, 0, canvas.width, canvas.height);
	current_index = 0;
	epoch_id.value = e_id;
	staticRender();
	startLive(0);
});
