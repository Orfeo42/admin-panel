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


const enableCharts = () => {
    const elements = document.querySelectorAll('.chart')
    elements.forEach(element => {
        new Chart(element, {
            type: 'line',
            data: {
                labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
                datasets: [
                    {
                        label: 'My First dataset',
                        backgroundColor: 'transparent',
                        borderColor: 'rgba(255,255,255,.55)',
                        data: [65, 59, 84, 84, 51, 55, 40]
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
                        min: 30,
                        max: 89,
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
        })
    });
}


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

