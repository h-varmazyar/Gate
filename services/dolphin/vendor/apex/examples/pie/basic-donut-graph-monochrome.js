var options = {
	chart: {
		width: 400,
		type: 'donut',
	},
	series: [25, 15, 44, 55, 41, 17],
	labels: ["شنبه", "یک شنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنجشنبه"],
	theme: {
		monochrome: {
			enabled: true,
			color: '#1a8e5f',
		}
	},
	title: {
		text: "فروش هفتگی",
	},
	responsive: [{
		breakpoint: 480,
		options: {
			chart: {
				width: 200
			},
			legend: {
				position: 'bottom'
			}
		}
	}],
	stroke: {
		width: 0,
	},
}
var chart = new ApexCharts(
	document.querySelector("#basic-donut-graph-monochrome"),
	options
);
chart.render();


