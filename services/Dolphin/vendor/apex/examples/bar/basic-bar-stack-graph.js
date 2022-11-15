var options = {
  chart: {
    height: 300,
    type: 'bar',
    stacked: true,
    zoom: {
			enabled: false
		},
  },
  dataLabels: {
		enabled: false
	},
  plotOptions: {
    bar: {
      horizontal: true,
    },
  },
  stroke: {
    width: 1,
    colors: ['#fff']
  },
  series: [{
    name: 'اسپری دریایی',
    data: [44, 55, 41, 37, 22, 43, 21]
  },{
    name: 'گوساله اعتصاب آور',
    data: [53, 32, 33, 52, 13, 43, 32]
  },{
    name: 'تصویر مخزن',
    data: [12, 17, 11, 9, 15, 11, 20]
  },{
    name: 'شیب سطل',
    data: [9, 7, 5, 8, 6, 9, 4]
  },{
    name: 'بچه متولد شده',
    data: [25, 12, 19, 32, 25, 24, 10]
  }],
  title: {
    text: 'فروش کتابهای داستانی',
    align: 'center'
  },
  xaxis: {
    categories: [2008, 2009, 2010, 2011, 2012, 2013, 2014],
    labels: {
      formatter: function(val) {
        return val + "هزار"
      }
    }
  },
  yaxis: {
    title: {
      text: undefined
    },
  },
  tooltip: {
    y: {
      formatter: function(val) {
      	return val + "هزار"
    	}
	  }
	},
	fill: {
		opacity: 1
	},
	legend: {
	  position: 'top',
	  horizontalAlign: 'left',
	  offsetX: 40
	},
	colors: ['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
}
var chart = new ApexCharts(
  document.querySelector("#basic-bar-stack-graph"),
  options
);
chart.render();