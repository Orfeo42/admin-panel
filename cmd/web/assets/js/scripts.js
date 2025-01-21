const buttonDeleteClick = (element) => {
    Swal.fire({
        title: 'Vuoi eliminare?',
        showCancelButton: true,
        confirmButtonText: 'Elimina',
        cancelButtonText: 'Annulla',
        text: 'Vuoi procedere all\'eliminazione di questo elemento?',
    }).then((result) => {
        if (result.isConfirmed) {
            htmx.trigger(element, 'delete')
        }
    });
}

const enableDeleteButtons = () => {
    const elements = document.querySelectorAll('.invoice-delete-button')
    elements.forEach(element => element.addEventListener('click', () => buttonDeleteClick(element)));
}

const buttonEditClick = (element) => {
    let editing = document.querySelector('.editing')
    if (editing) {
        Swal.fire({
            title: 'Stai già modificando!',
            showCancelButton: true,
            confirmButtonText: 'Sì, Modifica questa riga!',
            cancelButtonText: 'Annulla',
            text: 'Stai già modificando un altra riga. Vuoi annullare per modificare questa?'
        }).then((result) => {
            if (result.isConfirmed) {
                htmx.trigger(editing, 'cancel')
                htmx.trigger(element, 'edit')
            }
        })
        return;
    }
    htmx.trigger(element, 'edit')
}

const enableEditButtons = () => {
    const elements = document.querySelectorAll('.invoice-edit-button')
    elements.forEach(element => element.addEventListener('click', () => buttonEditClick(element)));
}

document.addEventListener("htmx:afterSwap", _ => {
    enableDeleteButtons();
    enableEditButtons();
});

const buttonPayClick = (element) => {
    Swal.fire({
        title: 'Vuoi impostare come pagata questa fattura?',
        showCancelButton: true,
        confirmButtonText: 'Paga',
        cancelButtonText: 'Annulla',
        text: 'Vuoi procedere all\'impostare questa fattura pagata?',
    }).then((result) => {
        if (result.isConfirmed) {
            htmx.trigger(element, 'pay')
        }
    });
}

const enablePayButtons = () => {
    const elements = document.querySelectorAll('.invoice-pay-button')
    elements.forEach(element => element.addEventListener('click', () => buttonPayClick(element)));
}


const enableHeaderToggler = () => {
    const elements = document.querySelectorAll('.header-toggler, .sidebar-toggler')

    elements.forEach(element => element.addEventListener('click', () => {
        coreui.Sidebar.getInstance(document.querySelector('#sidebar')).toggle()
    }));
}

const enableDropDown = () => {
    const dropdownElementList = document.querySelectorAll('.dropdown');
    dropdownElementList.forEach(element => new coreui.Dropdown(element))
}

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


const enableCharts = () => {
    enableSalesCharts();
    enableCollectedCharts();
    enableToBeCollectedCharts();
}

const updateValori = () => {
    const elementList = document.querySelectorAll('.total-button');
    elementList.forEach(element => element.addEventListener('click', () => {
        const selector = element.getAttribute("desc-target")
        const targetList = document.querySelectorAll(selector);
        targetList.forEach(target => target.innerHTML = element.innerHTML);
    }));
}

updateValori();
enableHeaderToggler();
enableDropDown();
enableCharts();

enableDeleteButtons();
enableEditButtons();
enablePayButtons();


document.addEventListener("htmx:afterSwap", _ => {
    enableDeleteButtons();
    enableEditButtons();
    enablePayButtons();
});

