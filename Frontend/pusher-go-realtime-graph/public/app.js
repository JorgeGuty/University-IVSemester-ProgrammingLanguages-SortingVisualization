document.getElementById("visualizeArray").addEventListener("click", function(){
    axios.get('/visualize')
})

const pusher = new Pusher('9515c01265248c7e86e8', {
    cluster: 'us2',
    encrypted: true
});

const visualization = pusher.subscribe('arrayVisualization');

visualization.bind('addNumber', data => {

    newLineChart.data.labels.push(data.Label);
    newLineChart.data.datasets[0].data.push(data.Value);

});

visualization.bind('update', () => {
    newLineChart.update();
});

function renderChart(userVisitsData) {
    var ctx = document.getElementById("realtimeChart").getContext("2d");

    var options = { animation: { duration: 0 }};

    newLineChart = new Chart(ctx, {
        type: "bar",
        data: userVisitsData,
        options: options
    });
}

var chartConfig = {
    labels: [],
    datasets: [
        {
            label: "",
            fill: false,
            lineTension: 0.1,
            backgroundColor: "rgba(75,192,192,0.4)",
            borderColor: "rgba(75,192,192,1)",
            borderCapStyle: 'butt',
            borderDash: [],
            borderDashOffset: 0.0,
            borderJoinStyle: 'miter',
            pointBorderColor: "rgba(75,192,192,1)",
            pointBackgroundColor: "#fff",
            pointBorderWidth: 1,
            pointHoverRadius: 5,
            pointHoverBackgroundColor: "rgba(75,192,192,1)",
            pointHoverBorderColor: "rgba(220,220,220,1)",
            pointHoverBorderWidth: 2,
            pointRadius: 1,
            pointHitRadius: 10,
            data: [],
            spanGaps: false,
        }
    ]
};

renderChart(chartConfig)