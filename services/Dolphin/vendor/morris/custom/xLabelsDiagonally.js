// Displaying X Labels Diagonally (Bar Chart)
var day_data = [
	{"period": "1398/12/4", "licensed": 4, "Nework": 2},
	{"period": "1398/12/5", "licensed": 4, "Nework": 2},
	{"period": "1398/12/6", "licensed": 4, "Nework": 2},
	{"period": "1398/12/7", "licensed": 4, "Nework": 2},
	{"period": "1398/12/8", "licensed": 4, "Nework": 2},
	{"period": "1398/12/9", "licensed": 4, "Nework": 2},
	{"period": "1398/12/10", "licensed": 4, "Nework": 2},
	{"period": "1398/12/11", "licensed": 4, "Nework": 2},
	{"period": "1398/12/12", "licensed": 5, "Nework": 1},
	{"period": "1398/12/13", "licensed": 8, "Nework": 4},
	{"period": "1398/12/14", "licensed": 2, "Nework": 2},
	{"period": "1398/12/15", "licensed": 7, "Nework": 6},
	{"period": "1398/12/16", "licensed": 4, "Nework": 3},
	{"period": "1398/12/17", "licensed": 7, "Nework": 7},
	{"period": "1398/12/18", "licensed": 8, "Nework": 2},
	{"period": "1398/12/19", "licensed": 9, "Nework": 3},
	{"period": "1398/12/20", "licensed": 2, "Nework": 9}
];
Morris.Bar({
	element: 'xLabelsDiagonally',
	data: day_data,
	xkey: 'period',
	ykeys: ['licensed', 'Nework'],
	labels: ['مجوز', 'شبکه'],
	xLabelAngle: 45,
	gridLineColor: "#e1e5f1",
	resize: true,
	hideHover: "auto",
	barColors:['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
});