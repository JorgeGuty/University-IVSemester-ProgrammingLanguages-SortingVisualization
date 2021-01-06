document.getElementById("visualizeArray").addEventListener("click", function(){
    let n = document.getElementById("elementsQuantity").value
    axios.get(`/visualize/${n}`)
    document.getElementById("solveAlgorithms").disabled = false
})
document.getElementById("solveAlgorithms").addEventListener("click", function(){
    axios.get('/solve')
})


const pusher = new Pusher('9515c01265248c7e86e8', {
    cluster: 'us2',
    encrypted: true
});

const visualization = pusher.subscribe('arrayVisualization');

visualization.bind('addNumber', data => {

    bubbleSortChart.data.labels.push(data.Label);
    bubbleSortChart.data.datasets[0].data.push(data.Value);

    heapSortChart.data.labels.push(data.Label);
    heapSortChart.data.datasets[0].data.push(data.Value);

    selectionSortChart.data.labels.push(data.Label);
    selectionSortChart.data.datasets[0].data.push(data.Value);

    quickSortChart.data.labels.push(data.Label);
    quickSortChart.data.datasets[0].data.push(data.Value);

    insertionSortChart.data.labels.push(data.Label);
    insertionSortChart.data.datasets[0].data.push(data.Value);

    treeSortChart.data.labels.push(data.Label);
    treeSortChart.data.datasets[0].data.push(data.Value);

});

visualization.bind('update', () => {

    bubbleSortChart.update();
    heapSortChart.update();
    selectionSortChart.update();
    quickSortChart.update();
    insertionSortChart.update();
    treeSortChart.update();
});

visualization.bind('swap', data => {

    let sortID = data.SortID
    let index1 = data.Index1
    let index2 = data.Index2

    switch (sortID)
    {
        case "bubble":
            swapArrayElements(bubbleSortChart.data.datasets[0].data, index1, index2)
            bubbleSortChart.update();
            break;
        case "heap":
            swapArrayElements(heapSortChart.data.datasets[0].data, index1, index2)
            heapSortChart.update();
            break;
        case "selection":
            swapArrayElements(selectionSortChart.data.datasets[0].data, index1, index2)
            selectionSortChart.update();
            break;
        case "insertion":
            swapArrayElements(insertionSortChart.data.datasets[0].data, index1, index2)
            insertionSortChart.update();
            break;
        case "quick":
            swapArrayElements(quickSortChart.data.datasets[0].data, index1, index2)
            quickSortChart.update();
            break;
        case "tree":
            swapArrayElements(treeSortChart.data.datasets[0].data, index1, index2)
            treeSortChart.update();
            break;
    }
});

function swapArrayElements(pArray, pIndex1, pIndex2){
    let auxiliarVariable = pArray[pIndex1]
    pArray[pIndex1] = pArray[pIndex2]
    pArray[pIndex2] = auxiliarVariable
}

function renderCharts() {

    var insertionCtx = document.getElementById("insertionSort").getContext("2d");
    var bubbleCtx    = document.getElementById("bubbleSort").getContext("2d");
    var heapCtx      = document.getElementById("heapSort").getContext("2d");
    var selectionCtx = document.getElementById("selectionSort").getContext("2d");
    var treeCtx      = document.getElementById("treeSort").getContext("2d");
    var quickCtx     = document.getElementById("quickSort").getContext("2d");

    var options = {
        animation: {
            duration: 0
        },
        scales: {
            xAxes: [{
                gridLines: {
                    color: "rgba(0, 0, 0, 0)",
                }
            }],
            yAxes: [{
                gridLines: {
                    color: "rgba(0, 0, 0, 0)",
                }
            }]
        }
    };

    insertionSortChart = new Chart(insertionCtx, {
        type: "bar",
        data: getConfig([]),
        options: options
    });
    quickSortChart = new Chart(quickCtx, {
        type: "bar",
        data: getConfig([]),
        options: options
    });
    treeSortChart = new Chart(treeCtx, {
        type: "bar",
        data: getConfig([]),
        options: options
    });
    selectionSortChart = new Chart(selectionCtx, {
        type: "bar",
        data: getConfig([]),
        options: options
    });
    heapSortChart = new Chart(heapCtx, {
        type: "bar",
        data: getConfig([]),
        options: options
    });
    bubbleSortChart = new Chart(bubbleCtx, {
        type: "bar",
        data: getConfig([]),
        options: options
    });

}

function getConfig(pDataArray){
    return {
        labels: [],
        datasets: [
            {
                label: "",
                fill: false,
                lineTension: 0.1,
                backgroundColor: "rgba(255,0,0,0.4)",
                borderColor: "rgb(0,0,0)",
                borderCapStyle: 'butt',
                borderDash: [],
                borderDashOffset: 0.0,
                borderJoinStyle: 'miter',
                pointBorderColor: "rgba(75,192,192,1)",
                pointBackgroundColor: "#000000",
                pointBorderWidth: 1,
                pointHoverRadius: 5,
                pointHoverBackgroundColor: "rgba(75,192,192,1)",
                pointHoverBorderColor: "rgba(220,220,220,1)",
                pointHoverBorderWidth: 2,
                pointRadius: 1,
                pointHitRadius: 10,
                data: pDataArray,
                spanGaps: false,
            }
        ]
    };
}

renderCharts()
