var options = {
	chart: {
		height: 350,
		type: 'line',
		zoom: {
			enabled: false
		},
	},
	toolbar: {
		show: false
	},
	series: [{
		name: 'وب سایت',
		type: 'column',
		data: [440, 505, 414, 671, 227, 413, 201, 352, 752, 320, 257, 160]
	}, {
		name: 'شبکه های اجتماعی',
		type: 'line',
		data: [23, 42, 35, 27, 43, 22, 17, 31, 22, 22, 12, 16]
	}],
	stroke: {
		width: [0, 4]
	},
	title: {
		text: 'منابع ترافیکی',
		align: 'center'
	},
	// labels: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
	labels: ['01 Jan 2001', '02 Jan 2001', '03 Jan 2001', '04 Jan 2001', '05 Jan 2001', '06 Jan 2001', '07 Jan 2001', '08 Jan 2001', '09 Jan 2001', '10 Jan 2001', '11 Jan 2001', '12 Jan 2001'],
	xaxis: {
		type: 'datetime'
	},
	colors: ['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
	yaxis: [{
		title: {
			text: 'وب سایت',
		},
	},{
		opposite: true,
		title: {
			text: 'شبکه های اجتماعی'
		}
	}]
}
var chart = new ApexCharts(
	document.querySelector("#line-column-graph"),
	options
);
chart.render();