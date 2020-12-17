document.getElementById("visualizeArray").addEventListener("click", function(){
    axios.get('/visualize')
})

const pusher = new Pusher('9515c01265248c7e86e8', {
    cluster: 'us2',
    encrypted: true
});

const visualization = pusher.subscribe('arrayVisualization');

visualization.bind('addNumber', data => {

    bubbleSortChart.data.labels.push(data.Label);
    bubbleSortChart.data.datasets[0].data.push(data.Value);

});

visualization.bind('update', () => {
    bubbleSortChart.update();
});

visualization.bind('swap', data => {

    bubbleSortChart.data.labels.push(data.Label);
    bubbleSortChart.data.datasets[0].data.push(data.Value);

});

function renderBubbleSortChart(pData) {
    var ctx = document.getElementById("bubbleSort").getContext("2d");

    var options = { animation: { duration: 0 }};

    bubbleSortChart = new Chart(ctx, {
        type: "bar",
        data: pData,
        options: options
    });

}

function renderHeapSortChart(pData) {
    var ctx = document.getElementById("heapSort").getContext("2d");

    var options = { animation: { duration: 0 }};

    heapSortChart = new Chart(ctx, {
        type: "bar",
        data: pData,
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

renderBubbleSortChart(chartConfig)
renderHeapSortChart(chartConfig)