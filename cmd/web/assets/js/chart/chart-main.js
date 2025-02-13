const random = (min, max) => Math.floor(Math.random() * (max - min + 1) + min);

const defaultInputChart = {
    labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
    datasets: [{
        label: 'My First dataset',
        backgroundColor: `rgba(51, 153, 255, .1)`,
        borderColor: "#39f",
        pointHoverBackgroundColor: '#fff',
        borderWidth: 2,
        data: [random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200)],
        fill: true
    }, {
        label: 'My Second dataset',
        borderColor: "#1b9e3e",
        pointHoverBackgroundColor: '#fff',
        borderWidth: 2,
        data: [random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200), random(50, 200)]
    }]
};

let mainChartElement;

export function mainChart(inputData = defaultInputChart) {
    if (mainChartElement) {
        mainChartElement.destroy()
    }
    new Chart(document.getElementById('main-chart'), {
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
