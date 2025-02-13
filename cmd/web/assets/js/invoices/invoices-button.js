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

const buttonAbortInsertClick = (element) => {
    Swal.fire({
        title: `Vuoi annullare l'inserimento?`,
        showCancelButton: true,
        confirmButtonText: 'Annulla',
        cancelButtonText: 'Continua',
        text: `Vuoi procedere all'annullamento del inserimento?`,
    }).then((result) => {
        if (result.isConfirmed) {
            element.closest('tr').remove();
        }
    });
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


export function enableInvoiceAbortInsertButtons() {
    const elements = document.querySelectorAll('.invoice-abort-add')
    elements.forEach(element => element.addEventListener('click', () => buttonAbortInsertClick(element)));
}

export function enableInvoiceEditButtons() {
    const elements = document.querySelectorAll('.invoice-edit-button')
    elements.forEach(element => element.addEventListener('click', () => buttonEditClick(element)));
}

export function enableInvoiceDeleteButtons() {
    const elements = document.querySelectorAll('.invoice-delete-button')
    elements.forEach(element => element.addEventListener('click', () => buttonDeleteClick(element)));
}

export function enableInvoicePayButtons() {
    const elements = document.querySelectorAll('.invoice-pay-button')
    elements.forEach(element => element.addEventListener('click', () => buttonPayClick(element)));
}



