const color = "#2596e7";

let time = new Date();

let activity = [
    {
        x: new Date(time),
        y: Math.random() * 100,
    },
];

let activity2 = [
    {
        x: new Date(time),
        y: Math.random() * 100,
    },
];
time.setMinutes(time.getMinutes() + 5);
activity.push({
    x: new Date(time),
    y: Math.random() * 100,
});
activity2.push({
    x: new Date(time),
    y: Math.random() * 100,
});
time.setMinutes(time.getMinutes() + 5);
activity.push({
    x: new Date(time),
    y: Math.random() * 100,
});
activity2.push({
    x: new Date(time),
    y: Math.random() * 100,
});

const options = {
    chart: {
        type: "line",
        width: "70%",
        foreColor: color,
        animations: {
            enabled: true,
            easing: "linear",
            dynamicAnimation: {
                speed: 200,
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
    tooltip: {
        enabled: true,
        followCursor: true,
        intersect: false,
        fillSeriesColor: false,
        theme: "dark",
        style: {
            fontSize: "1.3em",
        },
        onDatasetHover: {
            highlightDataSeries: false,
        },
    },

    stroke: {
        width: 4,
        curve: "smooth",
    },

    series: [
        { name: "Receive", data: activity },
        { name: "Send", data: activity2 },
    ],

    yaxis: {
        max: 100,
        min: 0,
        labels: {
            formatter: function (val) {
                return val / 1;
            },
        },
    },

    xaxis: {
        type: "datetime",
        //range: 3600000,
    },
};

const chart = new ApexCharts(document.getElementById("activity-chart"), options);
chart.render();

window.setInterval(function () {
    time.setMinutes(time.getMinutes() + 5);
    activity.push({
        x: new Date(time),
        y: Math.random() * 100,
    });
    activity2.push({
        x: new Date(time),
        y: Math.random() * 100,
    });

    chart.updateSeries([
        { name: "Receive", data: activity },
        { name: "Send", data: activity2 },
    ]);
}, 2000);
