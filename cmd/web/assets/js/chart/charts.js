import {mainChart} from "./chart-main";

const chartDefaultSettings = {
    type: 'line',
    data: {
        labels: ['Gennaio', 'Febbraio', 'Marzo', 'Aprile', 'Maggio', 'Giugno', 'Luglio', 'Agosto', 'Settembre', 'Ottobre', 'Novembre', 'Dicembre'],
        datasets: [
            {
                label: 'My First dataset',
                backgroundColor: 'transparent',
                borderColor: 'rgba(255,255,255,.55)',
                data: [65, 59, 84, 84, 51, 55, 40, 84, 84, 51, 55, 40]
            }
        ]
    },
    options: {
        plugins: {
            legend: {
                display: false
            }
        },
        maintainAspectRatio: false,
        scales: {
            x: {
                border: {
                    display: false
                },
                grid: {
                    display: false,
                    drawBorder: false
                },
                ticks: {
                    display: false
                }
            },
            y: {
                display: false,
                grid: {
                    display: false
                },
                ticks: {
                    display: false
                }
            }
        },
        elements: {
            line: {
                borderWidth: 1,
                tension: 0.4
            },
            point: {
                radius: 4,
                hitRadius: 10,
                hoverRadius: 4
            }
        }
    }
};


const enableSalesCharts = () => {
    const elements = document.querySelectorAll('.sales-chart')
    fetch('/sales/graph').then(res => {
        if (res.status === 200) {
            return res.json()
        }
    }).then(value => {
        const chartSettings = JSON.parse(JSON.stringify(chartDefaultSettings));
        chartSettings.data.labels = value.Labels
        chartSettings.data.datasets[0].data = value.Values
        chartSettings.data.datasets[0].label = "Fatturato"
        elements.forEach(element => new Chart(element, chartSettings));
    });
}


const enableCollectedCharts = () => {
    const elements = document.querySelectorAll('.collected-chart')
    fetch('/collected/graph').then(res => {
        if (res.status === 200) {
            return res.json()
        }
    }).then(value => {
        const chartSettings = JSON.parse(JSON.stringify(chartDefaultSettings));
        chartSettings.data.labels = value.Labels
        chartSettings.data.datasets[0].data = value.Values
        chartSettings.data.datasets[0].label = "Riscosso"
        elements.forEach(element => new Chart(element, chartSettings));
    });
}

const enableToBeCollectedCharts = () => {
    const elements = document.querySelectorAll('.to-be-collected-chart')
    fetch('/to-be-collected/graph').then(res => {
        if (res.status === 200) {
            return res.json()
        }
    }).then(value => {
        const chartSettings = JSON.parse(JSON.stringify(chartDefaultSettings));
        chartSettings.data.labels = value.Labels
        chartSettings.data.datasets[0].data = value.Values
        chartSettings.data.datasets[0].label = "Da riscuotere"
        elements.forEach(element => new Chart(element, chartSettings));
    });
}


export function enableCharts() {
    enableSalesCharts();
    enableCollectedCharts();
    enableToBeCollectedCharts();
    mainChart();
}
