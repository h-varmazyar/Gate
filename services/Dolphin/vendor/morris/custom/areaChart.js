// Morris Area Chart
Morris.Area({
	element: 'areaChart',
	data: [
		{ y: '2011', a: 10, b: 5, c: 2},
		{ y: '2012', a: 40, b: 15, c: 8},
		{ y: '2013', a: 15, b: 50, c: 25},
		{ y: '2014', a: 40, b: 15, c: 7},
		{ y: '2015', a: 20, b: 30, c: 20},
		{ y: '2016', a: 35, b: 15, c: 20},
		{ y: '2017', a: 20, b: 15, c: 51}
	],
	xkey: 'y',
	ykeys: ['a', 'b', 'c'],
	behaveLikeLine: !0,
	pointSize: 0,
	labels: ['فروش', 'هزینه ها', 'پروژه ها'],
	pointStrokeColors: ['#1a8e5f'],
	gridLineColor: "#e1e5f1",
	lineColors: ['#1a8e5f'],
	gridtextSize: 10,
	fillOpacity: 0.7,
	lineWidth: 0,
	hideHover: "auto",
	resize: true,
	redraw: true,
});