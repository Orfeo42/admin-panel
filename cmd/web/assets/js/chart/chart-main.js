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
    }]
};

let mainChartElement;

function mainChart({labels, values}) {
    if (mainChartElement) {
        mainChartElement.destroy()
    }

    const inputData = defaultInputChart

    inputData.labels = labels
    if (values) {
        inputData.datasets[0].data = values
    }
    const element = document.getElementById('main-chart')
    const chartParam = {
        type: 'line',
        data: inputData,
        options: {
            maintainAspectRatio: true,
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
                        color: `rgba(255, 255, 255, 0.95)`
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
                        color: `rgba(255, 255, 255, 0.95)`,
                        max: 250,
                        maxTicksLimit: 5
                    }
                }
            },
            elements: {
                line: {
                    tension: 0.4
                },
                point: {
                    radius: 4,
                    hitRadius: 10,
                    hoverRadius: 4,
                    hoverBorderWidth: 3
                }
            }
        }
    }

    mainChartElement = new Chart(element, chartParam);
}


export function loadMainCharData() {
    fetch('/chart/sales/month').then(async res => {
        if (res.status === 200) {
            return await res.json()
        }
        throw "Errore chiamata /chart/sales";
    }).then(value => mainChart(value));
}
