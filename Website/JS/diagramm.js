const options = {
	chart: {
		type: "line",
		width: "100%",
		foreColor: "#2596e7",
		redrawOnParentResize: true,
		animations: {
			enabled: true,
			easing: "linear",
			dynamicAnimation: {
				speed: 300,
			},
		},
		zoom: {
			type: "x",
			enabled: true,
			autoScaleYaxis: true,
		},
		toolbar: {
			show: true,
			offsetX: -15,
			offsetY: 0,
			tools: {
				download: false,
				selection: false,
				pan: true,

				zoom: true,
				zoomin: true,
				zoomout: true,
				reset: true,
			},
			autoSelected: "zoom",
		},
	},

	dataLabels: {
		enabled: true,
		style: {
			fontSize: '1em',
			fontWeight: 'bold',
		},
		offsetY: -15,
	},

	tooltip: {
		enabled: true,
		followCursor: true,
		intersect: false,
		fillSeriesColor: false,
		theme: "dark",
		style: {
			fontSize: "1.6em",
		},
		onDatasetHover: {
			highlightDataSeries: true,
		},
	},

	stroke: {
		width: 4,
		curve: "smooth",
		lineCap: 'butt',
	},

	forecastDataPoints: {
		count: 1,
		fillOpacity: 0.8,
		strokeWidth: 4,
		dashArray: 6,
	},

	noData: {
		text: "No Data Available",
		align: 'center',
		verticalAlign: 'middle',
		style: {
			fontSize: '2.5em',
		}
	},

	yaxis: {
		max: 100,
		min: 0,
		labels: {
			style: {
				fontSize: '1.2em',
			},
			"formatter": (val) => {
				return val / 1;
			},
		},
	},

	xaxis: {
		type: "datetime",
		tickAmount: 7,
		labels: {
			show: false,
			style: {
				fontSize: '1em',
			},
			"formatter": (val) => {
				return new Date(val).getDay() + "/" + new Date(val).getMonth() + "/" + new Date(val).getFullYear() + ":" + new Date(val).getHours()
			},
		},
		axisTicks: {
			show: true,
			borderType: 'solid',
			height: 12,
			offsetX: 0,
			offsetY: -6
		},
	},
};

class Chart {

	chart

	series = []
	options = options

	constructor(id, ...names) {
		names.forEach((val) => {
			this.series.push({name: val, data: []})
		})
		this.options.series = this.series
		this.chart = new ApexCharts(document.getElementById(id), this.options);
		this.chart.render();
	}

	clear() {
		this.series.forEach((val) => {
			val.data.clear()
		})
	}

	add(name, data) {
		this.series.filter((val) => {
			return val.name === name
		})[0].data.push(data)
	}

	updateChart() {
		this.chart.updateSeries(this.series);
	}
}
