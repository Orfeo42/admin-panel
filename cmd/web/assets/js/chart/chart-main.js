const random = (min, max) => Math.floor(Math.random() * (max - min + 1) + min);

const defaultInputChart = {
    labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
    datasets: [{
        label: 'Fatturato',
        backgroundColor: `rgba(36, 149, 66, .1)`,
        borderColor: "#1b9e3e",
        pointHoverBackgroundColor: '#fff',
        borderWidth: 2,
        data: [random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200)],
        fill: true
    }, {
        label: 'Riscosso',
        borderColor: "#39f",
        pointHoverBackgroundColor: '#fff',
        borderWidth: 2,
        data: [random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200)]
    }, {
        label: 'Da Riscuotere',
        borderColor: `rgba(219, 93, 93, 1)`,
        pointHoverBackgroundColor: '#fff',
        borderWidth: 2,
        data: [random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200)]
    }]
};

let mainChartElement;

function mainChart({labels, sales, collected, toBeCollected}) {
    if (mainChartElement) {
        mainChartElement.destroy()
    }

    const inputData = defaultInputChart

    inputData.labels = labels
    if (sales) {
        inputData.datasets[0] = sales
    }
    if (collected) {
        inputData.datasets[1] = collected
    }
    if (toBeCollected) {
        inputData.datasets[2] = toBeCollected
    }

    mainChartElement = new Chart(document.getElementById('main-chart'), {
        type: 'line',
        data: inputData,
        options: {
            maintainAspectRatio: false,
            plugins: {
                annotation: {
                    annotations: {
                        line1: {
                            type: 'line',
                            yMin: 95,
                            yMax: 95,
                            borderColor: "#e55353",
                            borderWidth: 1,
                            borderDash: [8, 5]
                        }
                    }
                },
                legend: {
                    display: false
                }
            },
            scales: {
                x: {
                    grid: {
                        color: `rgba(8, 10, 12, 0.175)`,
                        drawOnChartArea: false
                    },
                    ticks: {
                        color: `rgba(37, 43, 54, 0.95)`
                    }
                },
                y: {
                    border: {
                        color: `rgba(8, 10, 12, 0.175)`
                    },
                    grid: {
                        color: `rgba(8, 10, 12, 0.175)`
                    },
                    ticks: {
                        beginAtZero: true,
                        color: `rgba(37, 43, 54, 0.95)`,
                        max: 250,
                        maxTicksLimit: 5,
                        stepSize: Math.ceil(250 / 5)
                    }
                }
            },
            elements: {
                line: {
                    tension: 0.4
                },
                point: {
                    radius: 0,
                    hitRadius: 10,
                    hoverRadius: 4,
                    hoverBorderWidth: 3
                }
            }
        }
    });
}


export function loadMainCharData() {
    fetch('/chart/main').then(res => {
        if (res.status === 200) {
            return res.json()
        }
        throw "Errore chiamata /chart/to-be-collected";
    }).then(async value => mainChart(await value));
}
