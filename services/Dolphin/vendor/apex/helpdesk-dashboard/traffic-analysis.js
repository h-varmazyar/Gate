// By Device
var options = {
	chart: {
		width: '100%',
		height: 200,
		type: 'pie',
	},
	series: [2000, 3000, 4000, 5000, 6000],
	labels: ["وب", "فرم ها", "ایمیل ها", "چت", "تلفن"],
	stroke: {
		width: 0,
	},
	theme: {
		monochrome: {
			enabled: true,
			color: '#1a8e5f',
		}
	},
}
var chart = new ApexCharts(
	document.querySelector("#traffic-analysis"),
	options
);
chart.render();