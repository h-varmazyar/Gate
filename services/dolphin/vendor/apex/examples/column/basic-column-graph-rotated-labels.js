var options = {
  annotations: {
    points: [{
      x: 'Bananas',
      seriesIndex: 0,
      label: {
        borderColor: '#f38559',
        offsetY: 0,
        style: {
          color: '#ffffff',
          background: '#f38559',
          textSize: '10px',
        },
        text: 'فروش عالی',
      }
    }]
  },
  chart: {
    height: 380,
    type: "bar",
    toolbar: {
      show: true,
    },
  },
  plotOptions: {
    bar: {
      columnWidth: '50%',
      endingShape: 'rounded'
    }
  },
  dataLabels: {
    enabled: false
  },
  stroke: {
    width: 1
  },
  series: [{
    name: 'خدمات',
    data: [44, 55, 41, 67, 22, 43, 21, 33, 45, 31, 87, 65, 35, 25]
  }],
  grid: {
    row: {
      colors: ['#f5f9fe', '#ffffff']
    }
  },
  xaxis: {
    labels: {
      rotate: -45
    },
    categories: ['آپل', 'پرتقال ها', 'توت فرنگی', 'آناناس', 'انبه', 'موز',
      'تمشک ها', 'گلابی ها', 'هندوانه', 'گیلاس', 'انار', 'نارنگی', 'پاپایاس', 'هلو'
    ],
  },
  yaxis: {
    labels: {
      show: false,
    },
    axisBorder: {
      show: false,
    },
  },
  theme: {
    monochrome: {
      enabled: true,
      color: '#1a8e5f',
      shadeIntensity: 0.1
    },
  },
  fill: {
    type: 'gradient',
    gradient: {
      shade: 'light',
      type: "horizontal",
      gradientToColors: ['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
      shadeIntensity: 0.25,
      inverseColors: true,
      opacityFrom: 0.75,
      opacityTo: 0.85,
      stops: [50, 100]
    },
  },
  title: {
    text: 'اسفند 1398 فروش',
    floating: true,
    align: 'center',
    style: {
      color: '#444'
    }
  },
}

var chart = new ApexCharts(
  document.querySelector("#basic-column-graph-rotated-labels"),
  options
);

chart.render();